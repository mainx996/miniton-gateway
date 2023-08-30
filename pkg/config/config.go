package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var Config config

var Mode string

const (
	DevMode  = "dev"
	TestMode = "test"
	ProdMode = "prod"
)

type config struct {
	HTTPConfig  httpConfig  `yaml:"http"`
	EnvConfig   envConfig   `yaml:"env"`
	LogConfig   logConfig   `yaml:"log"`
	MysqlConfig mysqlConfig `yaml:"mysql"`
}

type httpConfig struct {
	Port            int `yaml:"port"`
	ShutdownTimeout int `yaml:"shutdownTimeout"`
}

type mysqlConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Addr     string `json:"addr"`
	DataBase string `json:"dataBase"`
}

type envConfig struct {
	Name string `yaml:"name"`
	Mode string `yaml:"mode"`
}

type logConfig struct {
	FileName  string `yaml:"filename"`
	MaxSize   int    `yaml:"maxSize"`
	MaxBackup int    `yaml:"maxBackups"`
	MaxAge    int    `yaml:"maxAge"`
	Level     int    `yaml:"level"`
}

func Init(mode string) {

	f := fmt.Sprintf("config/%v/config.yaml", mode)
	var err error
	f, err = filepath.Abs(f)
	if err != nil {
		panic("config cfgFile path error")
	}

	c := new(config)
	fileContent, err := os.ReadFile(f)
	if err != nil {
		panic(fmt.Sprintf("config ReadFile err,err is %v", err.Error()))
	}

	err = yaml.Unmarshal(fileContent, &c)
	if err != nil {
		panic(fmt.Sprintf("config Unmarshal err,err is %v", err.Error()))
	}
	Config = *c

	Mode = c.EnvConfig.Mode
	return
}
