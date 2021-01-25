package user

import (
	"github.com/innovember/forum/api/models"
)

type UserUsecase interface {
	Create(user *models.User) (status int, err error)
	GetAllUsers() (users []models.User, status int, err error)
	GetPassword(username string) (password string, status int, err error)
	FindUserByUsername(username string) (user *models.User, status int, err error)
	UpdateSession(userID int64, sessionValue string) (err error)
	ValidateSession(sessionValue string) (user *models.User, status int, err error)
	CheckSessionByUsername(username string) (status int, err error)
}
