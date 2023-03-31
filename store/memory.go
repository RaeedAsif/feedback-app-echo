package store

import (
	"errors"
	"log"
	"sort"

	"github.com/RaeedAsif/feedback-app-echo/models"
)

var isMemory = false

var lastUserID int
var lastFeedbackID int

var users map[int]models.User
var feedbacks map[int][]*models.Feedback

// IsMemory to get isMemory
func IsMemory() bool {
	return isMemory
}

// SetIsMemory to set isMemory
func setIsMemory(value bool) {
	isMemory = value
}

// SetUser to set users
func SetUser(user models.User) (int, error) {
	u, _ := GetUserByEmail(user.Email)
	if u != nil {
		return -1, errors.New("user exists with same email")
	}

	lastUserID = lastUserID + 1
	user.ID = lastUserID

	users[lastUserID] = user

	return user.ID, nil
}

// GetUser to get user
func GetUser(id int) (*models.User, error) {
	if val, ok := users[id]; ok {
		return &val, nil
	}

	return nil, errors.New("user not found")
}

// GetUserByEmail to get user by email
func GetUserByEmail(email string) (*models.User, error) {
	for _, u := range users {
		if u.Email == email {
			return &u, nil
		}
	}

	return nil, errors.New("user not found")

}

// SetFeedBack to set feedbacks
func SetFeedBack(fb models.Feedback, userId int) int {
	lastFeedbackID = lastFeedbackID + 1
	fb.ID = lastFeedbackID

	feedbacks[userId] = append(feedbacks[userId], &fb)

	return fb.ID
}

// GetFeedbacks to get feedbacks
func GetFeedbacksByUser(userID, page int, type_ string) (int64, []*models.Feedback) {
	if val, ok := feedbacks[userID]; ok {
		data := make([]*models.Feedback, 0)

		if type_ == "" {
			data = val
		} else {
			for _, v := range val {
				if v.Type == type_ {
					data = append(data, v)
				}
			}
		}

		sort.Slice(data, func(i, j int) bool {
			return data[i].Date.After(data[j].Date)
		})

		count := int64(len(data))
		offset := 10
		limit := offset * page
		if offset > len(data[limit-10:]) {
			return int64(len(data)), data
		}
		return count, data[limit-10 : limit]
	}

	return -1, nil
}

// InitMemory to initliase memory db
func InitMemory() error {
	setIsMemory(true)
	users = make(map[int]models.User)
	feedbacks = make(map[int][]*models.Feedback)
	lastUserID = 0
	lastFeedbackID = 0

	log.Println("? Initialised Memory Database Successfully")
	return nil
}
