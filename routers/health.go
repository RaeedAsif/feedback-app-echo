package routers

import (
	"github.com/RaeedAsif/feedback-app-echo/controllers"
	"github.com/labstack/echo/v4"
)

func initServerHealthRouter(e *echo.Echo) {
	e.GET("/health", getHealth)
}

// getHealth godoc
// @Summary      Show server health
// @Description  GET API which returs server health
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.ResponseHealth
// @Failure      500  {object}  json.ResponseError
// @Router       /health [get]
func getHealth(c echo.Context) error {
	return controllers.ServerHealth(c)
}
