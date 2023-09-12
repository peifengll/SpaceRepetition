package service

import (
	"SpaceRepetition/src/internal/model"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"strconv"
)

func HelpUserRegisterStroge(db *gorm.DB, id int) {
	db.Table(model.FloderNsp.PersonTableName(strconv.Itoa(id))).AutoMigrate(&model.Floder{})
	db.Table(model.DeckNsp.PersonTableName(strconv.Itoa(id))).AutoMigrate(&model.Deck{})
	db.Table(model.KnowledgeNsp.PersonTableName(strconv.Itoa(id))).AutoMigrate(&model.Knowledge{})
	db.Table(model.RecordNsp.PersonTableName(strconv.Itoa(id))).AutoMigrate(&model.Record{})
	db.Table(model.SectionNsp.PersonTableName(strconv.Itoa(id))).AutoMigrate(&model.Section{})
	// 注册这些表，
}

func GetJwtToken(secretKey string, iat, seconds int64, userId int) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
