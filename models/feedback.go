package models

import "time"

// Feedback model
type Feedback struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Date      time.Time `gorm:"not null" json:"date"`
	Type      string    `gorm:"not null" json:"type"`
	Feedback  string    `gorm:"not null" json:"feedback"`
	UserID    int       `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}

type FeedbackInput struct {
	Type     string `json:"type" binding:"required"`
	Feedback string `json:"feedback" binding:"required"`
}

// FeedbackListResponse model
type FeedbackListResponse struct {
	Count   int64       `json:"count"`
	Content []*Feedback `json:"content"`
}

// FeedbackResponse model
type FeedbackResponse struct {
	Date     time.Time `json:"date"`
	Type     string    `json:"type"`
	Feedback string    `json:"feedback"`
}
