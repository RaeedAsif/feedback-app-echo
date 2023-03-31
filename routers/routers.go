package routers

import (
	"log"

	"github.com/labstack/echo/v4"
)

func InitRouters(e *echo.Echo) {
	initServerHealthRouter(e)
	initAuthRouters(e)
	initFeedbackRouters(e)
	initUserRouters(e)

	log.Println("? Routers Initialised")
}
