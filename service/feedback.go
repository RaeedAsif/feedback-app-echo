package service

import (
	"github.com/RaeedAsif/feedback-app-echo/models"
	"github.com/RaeedAsif/feedback-app-echo/store"
	"gorm.io/gorm"
)

var feedbackColumn = []string{"id", "date", "user_id", "type", "feedback", "created_at", "updated_at"}

// CreateFeedback service function to create feedback
func CreateFeedback(DB *gorm.DB, feedback models.Feedback) (int, error) {
	if store.IsMemory() {
		return store.SetFeedBack(feedback, feedback.UserID), nil
	}

	result := DB.Create(&feedback)
	if result.Error != nil {
		return -1, result.Error
	}

	return feedback.ID, nil
}

// FindFeedbacksByUser service function to get feedbacks by user id
func FindFeedbacksByUser(DB *gorm.DB, id, page int, type_ string) (models.FeedbackListResponse, error) {
	if store.IsMemory() {
		count, feedback := store.GetFeedbacksByUser(id, page, type_)
		return models.FeedbackListResponse{
			Count:   count,
			Content: feedback,
		}, nil
	}

	var count int64
	if type_ == "" {
		DB.Model(models.Feedback{}).Where("user_id = ?", id).Count(&count)
	} else {
		DB.Model(models.Feedback{}).Where("user_id = ? AND type = ?", id, type_).Count(&count)
	}
	var feedbacks []*models.Feedback
	perPage := 10
	offset := (page - 1) * perPage

	if type_ == "" {
		DB.Select(feedbackColumn).Find(&feedbacks, "user_id = ? ORDER BY date LIMIT ? OFFSET ?", id, perPage, offset)
	} else {
		DB.Select(feedbackColumn).Where("user_id = ? AND type = ?", id, type_).Order("date").Limit(perPage).Offset(offset).Find(&feedbacks)
	}

	return models.FeedbackListResponse{
		Count:   count,
		Content: feedbacks,
	}, nil
}
