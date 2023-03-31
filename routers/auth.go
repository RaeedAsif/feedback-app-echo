package routers

import (
	"github.com/labstack/echo/v4"

	"github.com/RaeedAsif/feedback-app-echo/controllers"
	"github.com/RaeedAsif/feedback-app-echo/store"
)

func initAuthRouters(e *echo.Echo) {
	e.POST("/register", signUp)
	e.POST("/login", login)
	e.GET("/refreshtoken", refreshToken)
}

// signUp godoc
// @Summary      User signup
// @Description  POST API which creates user
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.ResponseSuccess{id=number}
// @Failure      500  {object}  json.ResponseError
// @Param 		 request body models.SignUpInput true "query params"
// @Router       /register [post]
func signUp(c echo.Context) error {
	authController := controllers.NewAuthController(store.DB)
	return authController.SignUpUser(c)
}

// login godoc
// @Summary      Sign In
// @Description  POST API which logs in user
// @Accept       json
// @Produce      json
// @Success      200  {object}  json.ResponseSuccess{data=models.Token}
// @Failure      403  {object}  json.ResponseError
// @Failure      500  {object}  json.ResponseError
// @Param 		 request body models.SignInInput true "query params"
// @Router       /login [post]
func login(c echo.Context) error {
	authController := controllers.NewAuthController(store.DB)
	return authController.SignInUser(c)
}

// refreshToken godoc
// @Summary      Refresh token
// @Description  GET API which refresh jwt token
// @Accept       json
// @Produce      json
// @param 		 Authorization header string true "Authorization"
// @Success      200  {object}  json.ResponseSuccess{data=models.Token}
// @Failure      403  {object}  json.ResponseError
// @Failure      500  {object}  json.ResponseError
// @Router       /refreshtoken [get]
func refreshToken(c echo.Context) error {
	authController := controllers.NewAuthController(store.DB)
	return authController.RefreshAccessToken(c)
}
