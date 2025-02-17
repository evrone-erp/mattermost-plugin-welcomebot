package usecase

import (
	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	"github.com/mattermost/mattermost/server/public/model"
)

type TeamWelcomeRepo interface {
	GetTeamWelcome(teamID string) (*pmodel.TeamWelcome, *model.AppError)
	SetTeamWelcome(teamID string, tw *pmodel.TeamWelcome) *model.AppError
}
