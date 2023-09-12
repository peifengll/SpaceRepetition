package svc

import (
	"SpaceRepetition/src/api/internal/config"
	"SpaceRepetition/src/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config   config.Config
	DbEngine *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	//启动Gorm支持
	db, err := gorm.Open(mysql.Open(c.DataSourceName), &gorm.Config{})
	//如果出错就GameOver了
	if err != nil {
		panic(err)
	}
	//自动同步更新表结构,不要建表了O(∩_∩)O哈哈~
	db.AutoMigrate(&model.User{})

	return &ServiceContext{
		Config:   c,
		DbEngine: db,
	}
}
