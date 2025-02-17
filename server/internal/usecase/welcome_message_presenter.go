package usecase

import (
	"github.com/mattermost/mattermost/server/public/model"
)

type WelcomeMessagePresenter interface {
	Render(message string, userID string) (string, *model.AppError)
}
