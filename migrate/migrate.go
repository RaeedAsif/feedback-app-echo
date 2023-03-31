package main

import (
	"log"

	"github.com/RaeedAsif/feedback-app-echo/config"
	"github.com/RaeedAsif/feedback-app-echo/models"
	"github.com/RaeedAsif/feedback-app-echo/store"
)

func init() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalln("? Could not load environment variables", err)
	}

	store.ConnectDB(&config)
}

func main() {
	err := store.DB.AutoMigrate(
		&models.User{},
		&models.Feedback{},
	)
	if err != nil {
		log.Fatalln("? Could not complete migration", err)
	}

	log.Println("? Migration complete")
}
