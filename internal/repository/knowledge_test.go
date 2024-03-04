package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestKnowledgeRepository_GetAllReviewCard(t *testing.T) {
	dsn := "root:peifeng@tcp(peifeng.site:3306)/spacere?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	r := NewKnowledgeRepository(&Repository{
		db:     db.Debug(),
		rdb:    nil,
		logger: nil,
	})
	card, err := r.GetAllReviewCard("B8tqbl8XoH")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v\n", card)

}
