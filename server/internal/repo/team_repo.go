package repo

import (
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
)

type TeamRepo struct {
	api       plugin.API
	botUserID string
}

func NewTeamRepo(p BotAPIProvider) *TeamRepo {
	return &TeamRepo{
		api:       p.APIHandle(),
		botUserID: p.BotUserIDHandle(),
	}
}

func (r *TeamRepo) GetByTeamID(teamID string) (*model.Team, *model.AppError) {
	return r.api.GetTeam(teamID)
}

func (r *TeamRepo) GetByTeamName(teamName string) (*model.Team, *model.AppError) {
	return r.api.GetTeamByName(teamName)
}
