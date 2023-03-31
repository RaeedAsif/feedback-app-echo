package middleware

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/RaeedAsif/feedback-app-echo/config"
	"github.com/RaeedAsif/feedback-app-echo/json"
	"github.com/RaeedAsif/feedback-app-echo/models"
	"github.com/RaeedAsif/feedback-app-echo/service"
	"github.com/RaeedAsif/feedback-app-echo/store"
	"github.com/RaeedAsif/feedback-app-echo/utils"
	"github.com/labstack/echo/v4"
)

func DeserializeUser(c echo.Context) (*models.User, error) {
	var access_token string
	authorizationHeader := c.Request().Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		access_token = fields[1]
	}

	if access_token == "" {
		return nil, errors.New("user_not_logged_in")
	}

	config, _ := config.LoadConfig(".")
	sub, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
	if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(fmt.Sprint(sub))
	if err != nil {
		return nil, err
	}

	return service.FindUser(store.DB, id)
}

// DeserializeUser middleware fucntion to deserialize user
func AuthJWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user, err := DeserializeUser(c)
		if err != nil {
			return json.Error(c, err)
		}

		c.Set("currentUser", user)

		return next(c)
	}
}
