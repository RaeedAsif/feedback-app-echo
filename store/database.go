package store

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/RaeedAsif/feedback-app-echo/config"
)

var DB *gorm.DB

// ConnectDB to connect gorm to db
func ConnectDB(config *config.Config) error {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.New("failed_to_connect_to_the_database")
	}

	log.Println("? Connected Successfully to the Database")
	return nil
}
