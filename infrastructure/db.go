package infrastructure

import (
	MODEL_ACCOUNT "go-user/model/account"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&MODEL_ACCOUNT.User{})

	return db
}
