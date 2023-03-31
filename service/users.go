package service

import (
	"errors"
	"strings"

	"github.com/RaeedAsif/feedback-app-echo/models"
	"github.com/RaeedAsif/feedback-app-echo/store"
	"gorm.io/gorm"
)

var userColumn = []string{"id", "first_name", "last_name", "email", "created_at", "updated_at"}

// CreateUser service function to create user
func CreateUser(DB *gorm.DB, user models.User) (int, error) {
	if store.IsMemory() {
		return store.SetUser(user)
	}

	result := DB.Create(&user)
	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return -1, errors.New("user exists with same email")
	} else if result.Error != nil {
		return -1, result.Error
	}

	return user.ID, nil
}

// FindUser service function to get user by id
func FindUser(DB *gorm.DB, id int) (*models.User, error) {
	if store.IsMemory() {
		return store.GetUser(id)
	}

	var user *models.User
	result := DB.Select(userColumn).Find(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	user.Password = ""

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func FindUserByEmail(DB *gorm.DB, email string) (*models.User, error) {
	if store.IsMemory() {
		return store.GetUserByEmail(email)
	}

	var user models.User
	result := DB.First(&user, "email = ?", strings.ToLower(email))
	if result.Error != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}
