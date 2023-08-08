package main

import (
	_ "SpaceRepetition/src/config"
	"SpaceRepetition/src/db"
	"SpaceRepetition/src/models"
)

func main() {

	k := models.Knowledge{
		Font: "this is font",
		Back: "and here is back",
	}
	db.GetConnect().Create(&k)
}
