package config

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type MysqlConfig struct {
	Db         string `ini:"db"`
	DbHost     string `ini:"dbHost"`
	DbPort     int    `ini:"dbPort"`
	DbUser     string `ini:"dbUser"`
	DbPassWord string `ini:"dbPassWord"`
	DbName     string `ini:"dbName"`
}
type Config struct {
	MysqlConf MysqlConfig `ini:"mysql"`
}

var Conf Config

func init() {
	load, err := ini.Load("config/conf.ini")
	if err != nil {
		fmt.Printf("配置加载失败", err)
		return
	}
	load.MapTo(&Conf)
	if err != nil {
		fmt.Printf("配置映射失败", err)
		return
	}
}
