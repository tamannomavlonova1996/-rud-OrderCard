package db

import (
	"awesomeProject2/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"log"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	DSN := "host=localhost user=postgres password=postgres dbname=visa_application port=5467 sslmode=disable"
	DB, err = gorm.Open("postgres", DSN)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	DB.AutoMigrate(&models.OrderCard{}, &models.Card{}, &models.User{}, &models.Account{}, &models.Operation{})

	return DB, nil
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
