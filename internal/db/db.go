package db

import (
	"awesomeProject2/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error
	DSN := "host=localhost user=postgres password=postgres dbname=visa_application port=5467 sslmode=disable"
	DB, err = gorm.Open("postgres", DSN)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	DB.AutoMigrate(&models.Cards{})

	return DB, nil
}
