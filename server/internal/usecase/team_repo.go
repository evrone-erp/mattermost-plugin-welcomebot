package usecase

import (
	"github.com/mattermost/mattermost/server/public/model"
)

type TeamRepo interface {
	GetByTeamID(string) (*model.Team, *model.AppError)
	GetByTeamName(string) (*model.Team, *model.AppError)
}
