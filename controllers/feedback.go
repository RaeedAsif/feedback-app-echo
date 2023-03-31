package controllers

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/RaeedAsif/feedback-app-echo/json"
	"github.com/RaeedAsif/feedback-app-echo/middleware"
	"github.com/RaeedAsif/feedback-app-echo/models"
	"github.com/RaeedAsif/feedback-app-echo/service"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// UserController model
type FeedbackController struct {
	DB *gorm.DB
}

// NewAuthController init
func NewFeedbackController(DB *gorm.DB) FeedbackController {
	return FeedbackController{DB}
}

// SignUpUser handler function to sign up user
func (f *FeedbackController) CreateFeedback(c echo.Context) error {
	user, err := middleware.DeserializeUser(c)
	if err != nil {
		return json.Error(c, err)
	}

	var payload *models.FeedbackInput
	if err := c.Bind(&payload); err != nil {
		return json.Error(c, err)
	}

	if payload.Type == "" {
		return json.Error(c, errors.New("type cannot be empty"))
	}

	if payload.Feedback == "" {
		return json.Error(c, errors.New("feedback cannot be empty"))
	}

	now := time.Now()
	newFeedback := models.Feedback{
		Date:      now,
		Type:      payload.Type,
		Feedback:  payload.Feedback,
		UserID:    user.ID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	id, err := service.CreateFeedback(f.DB, newFeedback)
	if err != nil {
		return json.Error(c, err)
	}

	return json.SuccessCreate(c, id)
}

// GetFeedbacks handler function to get signed user's feedbacks
func (f *FeedbackController) GetFeedbacks(c echo.Context) error {
	user, err := middleware.DeserializeUser(c)
	if err != nil {
		return json.Error(c, err)
	}

	queryParams := c.QueryParams()
	pageStr := queryParams.Get("page")

	page, _ := strconv.Atoi(pageStr)
	if err != nil {
		return json.Error(c, errors.New("invalid page parameter"))
	}

	if page < 1 {
		page = 1
	}

	typeStr := queryParams.Get("type")

	fmt.Println(page)
	feedbacks, err := service.FindFeedbacksByUser(f.DB, user.ID, page, typeStr)
	if err != nil {
		return json.Error(c, err)
	}

	return json.Success(c, feedbacks)
}
