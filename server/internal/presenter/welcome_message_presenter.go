package presenter

import (
	"strings"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/mattermost/mattermost/server/public/model"
)

type WelcomeMessagePresenter struct {
	UserRepo usecase.UserRepo
}

func (p *WelcomeMessagePresenter) Render(message string, userID string) (string, *model.AppError) {
	user, err := p.UserRepo.GetByID(userID)

	if err != nil {
		return "", err
	}

	// just in case
	if user == nil {
		return "", &model.AppError{Message: "User not found"}
	}

	displayName := user.GetDisplayName(model.ShowNicknameFullName)
	handleName := user.Username

	message = strings.ReplaceAll(message, "{{.UserDisplayName}}", displayName)
	message = strings.ReplaceAll(message, "{{.UserHandleName}}", "@"+handleName)

	return message, nil
}
