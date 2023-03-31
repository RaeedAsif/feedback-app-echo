package controllers

import (
	"github.com/RaeedAsif/feedback-app-echo/json"
	"github.com/RaeedAsif/feedback-app-echo/middleware"
	"github.com/RaeedAsif/feedback-app-echo/models"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// UserController model
type UserController struct {
	DB *gorm.DB
}

// NewUserController init
func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

// GetUser handler function to get in current user
func (u *UserController) GetCurrentUser(c echo.Context) error {
	user, err := middleware.DeserializeUser(c)
	if err != nil {
		return json.Error(c, err)
	}

	userData := models.UserData{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return json.Success(c, userData)
}
