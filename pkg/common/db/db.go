package db

import (
	"log"

	"github.com/kevin-neves/golang-restApi-with-docker/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(&models.Profile{})
	return db
}
