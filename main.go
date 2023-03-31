package main

import (
	"fmt"
	"log"

	"github.com/RaeedAsif/feedback-app-echo/config"
	"github.com/RaeedAsif/feedback-app-echo/routers"
	"github.com/RaeedAsif/feedback-app-echo/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/RaeedAsif/feedback-app-echo/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title feedback-app-echo API
// @version 1.0
// @description This is a swagger docs for feedback-app-echo assignment.
// @termsOfService http://swagger.io/terms/

// @host limitless-tor-38427.herokuapp.com
// @BasePath /
// @schemes https

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func main() {
	isMemory := false

	// Define API wrapper
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())  // Logger
	e.Use(middleware.Recover()) // Recover

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	// DB CONNECT
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	err = store.ConnectDB(&config)
	if err != nil {
		log.Println("? Failed to connect to the Database, accessing memory store")
		isMemory = true
	}

	if isMemory {
		store.InitMemory()
	}

	// ROUTERS
	routers.InitRouters(e)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// SERVER
	e.Logger.Fatal(e.Start(":" + fmt.Sprintf("%d", config.PORT)))
}
