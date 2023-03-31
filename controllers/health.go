package controllers

import (
	"github.com/RaeedAsif/feedback-app-echo/json"
	"github.com/labstack/echo/v4"
)

// LogoutUser handler function to log out signed user
func ServerHealth(ctx echo.Context) error {
	return json.Health(ctx)
}
