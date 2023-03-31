package models

import (
	"time"
)

// User model
type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(255);not null" json:"last_name"`
	Email     string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"pasword,omitempty"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

// User model
type UserData struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName  string    `gorm:"type:varchar(255);not null" json:"last_name"`
	Email     string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

// SignUpInput model
type SignUpInput struct {
	FirstName       string `json:"first_name" binding:"required"`
	LastName        string `json:"last_name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" binding:"required,min=8"`
}

// SignInInput model
type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

// UserResponse model
type UserResponse struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
