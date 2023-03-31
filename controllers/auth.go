package controllers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/RaeedAsif/feedback-app-echo/config"
	"github.com/RaeedAsif/feedback-app-echo/json"
	"github.com/RaeedAsif/feedback-app-echo/models"
	"github.com/RaeedAsif/feedback-app-echo/service"
	"github.com/RaeedAsif/feedback-app-echo/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// AuthController model
type AuthController struct {
	DB *gorm.DB
}

// NewAuthController init
func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

// SignUpUser handler function to sign up user
func (ac *AuthController) SignUpUser(c echo.Context) error {
	var payload *models.SignUpInput

	if err := c.Bind(&payload); err != nil {
		return json.Error(c, err)
	}

	if payload.Password == "" {
		return json.Error(c, errors.New("password cannot be empty"))
	}

	err := utils.IsValidPasswordCheck(payload.Password)
	if err != nil {
		return json.Error(c, err)
	}

	if payload.Password != payload.ConfirmPassword {
		return json.Error(c, errors.New("passwords do not match"))
	}

	if !utils.IsEmailValid(payload.Email) {
		return json.Error(c, errors.New("invalid email"))
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		return json.Error(c, err)
	}

	now := time.Now()
	newUser := models.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     strings.ToLower(payload.Email),
		Password:  hashedPassword,
		CreatedAt: now,
		UpdatedAt: now,
	}

	id, err := service.CreateUser(ac.DB, newUser)
	if err != nil {
		return json.Error(c, err)
	}

	return json.SuccessCreate(c, id)
}

// SignInUser handler function to log in user
func (ac *AuthController) SignInUser(c echo.Context) error {
	var payload *models.SignInInput
	if err := c.Bind(&payload); err != nil {
		return json.Error(c, err)
	}

	if !utils.IsEmailValid(payload.Email) {
		return json.Error(c, errors.New("invalid email"))
	}

	user, err := service.FindUserByEmail(ac.DB, payload.Email)
	if err != nil {
		return json.Error(c, err)
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		return json.Error(c, errors.New("invalid email or password"))
	}

	config, _ := config.LoadConfig(".")

	// Generate Tokens
	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		return json.Error(c, err)
	}

	refresh_token, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		return json.Error(c, err)
	}

	tokenResponse := &models.Token{AccessToken: access_token, RefreshToken: refresh_token}

	return json.Success(c, tokenResponse)
}

// RefreshAccessToken handler function to refresh access token
func (ac *AuthController) RefreshAccessToken(c echo.Context) error {
	message := "could not refresh access token"
	var refresh_token string

	authorizationHeader := c.Request().Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)
	if len(fields) != 0 && fields[0] == "Bearer" {
		refresh_token = fields[1]
	} else {
		return json.Error(c, errors.New(message))
	}

	config, _ := config.LoadConfig(".")

	sub, err := utils.ValidateToken(refresh_token, config.RefreshTokenPublicKey)
	if err != nil {
		return json.Error(c, errors.New(message))
	}

	id, err := strconv.Atoi(fmt.Sprint(sub))
	if err != nil {
		return json.Error(c, err)
	}

	user, err := service.FindUser(ac.DB, id)
	if err != nil {
		return json.Error(c, err)
	}

	access_token, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		return json.Error(c, err)
	}

	tokenResponse := &models.Token{AccessToken: access_token, RefreshToken: refresh_token}

	return json.Success(c, tokenResponse)
}
