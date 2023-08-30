package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"www.miniton-gateway.com/pkg/config"
)

var (
	pingFrequency = 10 * time.Second
	DB            *gorm.DB
)

func Init() {
	c := config.Config.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true&loc=Local", c.Username, c.Password, c.Addr, c.DataBase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("mysql connect err,err is %v", err.Error()))
	}
	if config.Mode != config.ProdMode {
		db = db.Debug()
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("mysql connect err,err is %v", err.Error()))
	}
	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetMaxIdleConns(15)
	sqlDB.SetConnMaxLifetime(time.Duration(900) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(900) * time.Second)
	DB = db
	go Ping()
}

func Ping() {
	t := time.NewTicker(pingFrequency)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			sqlDB, err := DB.DB()
			if err == nil {
				_ = sqlDB.Ping()
			}
		}
	}
}
