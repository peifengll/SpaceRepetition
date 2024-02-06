package main

import (
	"github.com/peifengll/SpaceRepetition/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:peifeng@tcp(127.0.0.1:3306)/spacere?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("包打错")
	}
	db.AutoMigrate(&model.Deck{}, &model.Floder{}, &model.Knowledge{}, &model.Record{}, &model.Section{}, &model.User{})
}
