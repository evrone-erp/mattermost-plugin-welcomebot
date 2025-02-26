package repo

import (
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
)

type UserRepo struct {
	api       plugin.API
	botUserID string
}

func NewUserRepo(p BotAPIProvider) *UserRepo {
	return &UserRepo{
		api:       p.APIHandle(),
		botUserID: p.BotUserIDHandle(),
	}
}

func (r *UserRepo) GetByID(userID string) (*model.User, *model.AppError) {
	return r.api.GetUser(userID)
}
