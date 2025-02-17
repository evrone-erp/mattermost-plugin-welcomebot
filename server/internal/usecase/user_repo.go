package usecase

import (
	"github.com/mattermost/mattermost/server/public/model"
)

type UserRepo interface {
	GetByID(userID string) (*model.User, *model.AppError)
}
