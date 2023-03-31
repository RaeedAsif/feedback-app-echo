package routers

import (
	"github.com/RaeedAsif/feedback-app-echo/controllers"
	"github.com/RaeedAsif/feedback-app-echo/middleware"
	"github.com/RaeedAsif/feedback-app-echo/store"
	"github.com/labstack/echo/v4"
)

func initUserRouters(e *echo.Echo) {
	e.GET("/user", middleware.AuthJWTMiddleware(getCurrentUser))
}

// getCurrentUser godoc
// @Summary      Show current user
// @Description  GET API which returs current user
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.ResponseSuccess{data=models.UserResponse}
// @Failure      403  {object}  json.ResponseError
// @Failure      500  {object}  json.ResponseError
// @Security 	ApiKeyAuth
// @Router       /user [get]
func getCurrentUser(c echo.Context) error {
	uController := controllers.NewUserController(store.DB)
	return uController.GetCurrentUser(c)
}
