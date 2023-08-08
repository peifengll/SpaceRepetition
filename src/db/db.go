package db

import (
	"SpaceRepetition/src/config"
	"SpaceRepetition/src/models"
	"SpaceRepetition/src/utils/fsrs"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		config.Conf.MysqlConf.DbUser, config.Conf.MysqlConf.DbPassWord,
		config.Conf.MysqlConf.DbHost, config.Conf.MysqlConf.DbPort, config.Conf.MysqlConf.DbName, "10s")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	db.AutoMigrate(models.Deck{})
	db.AutoMigrate(models.Knowledge{})
	db.AutoMigrate(models.User{})
	db.AutoMigrate(models.Note{})
	db.AutoMigrate(models.Deck{})
	db.AutoMigrate(fsrs.Card{})
	db.AutoMigrate(fsrs.ReviewLog{})

}
func GetConnect() *gorm.DB {
	return db
}
