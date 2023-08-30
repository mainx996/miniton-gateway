package demo

import (
	"context"
	"www.miniton-gateway.com/pkg/mysql"
)

const TableName = "test"

type (
	Entity struct {
		ID   int64  `gorm:"column:id" json:"id"`
		Name string `gorm:"column:name"  json:"name"`
		Age  int64  `gorm:"column:age" json:"age"`
	}
)

func (e *Entity) TableName() string {
	return TableName
}

func DetailByID(ctx context.Context, id int64) (entity *Entity, err error) {
	entity = new(Entity)
	err = mysql.DB.Model(entity).Where("id = ?", id).Find(&entity).Error
	return
}

func List(ctx context.Context) (entities []*Entity, err error) {
	entities = make([]*Entity, 0)
	err = mysql.DB.Model(&Entity{}).Find(&entities).Error
	return
}

func Count(ctx context.Context) (total int64, err error) {
	err = mysql.DB.Model(&Entity{}).Count(&total).Error
	return
}

func Create(ctx context.Context, entity *Entity) (err error) {
	return mysql.DB.Model(&Entity{}).Create(entity).Error
}
