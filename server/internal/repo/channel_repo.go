package repo

import (
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
)

type ChannelRepo struct {
	api       plugin.API
	botUserID string
}

func NewChannelRepo(p BotAPIProvider) *ChannelRepo {
	return &ChannelRepo{
		api:       p.APIHandle(),
		botUserID: p.BotUserIDHandle(),
	}
}

func (r *ChannelRepo) Get(channelID string) (*model.Channel, *model.AppError) {
	return r.api.GetChannel(channelID)
}

func (r *ChannelRepo) GetByName(teamID string, channelName string) (*model.Channel, *model.AppError) {
	return r.api.GetChannelByName(teamID, channelName, false)
}

func (r *ChannelRepo) AddMemberByUserID(channelID, userID string) (*model.ChannelMember, *model.AppError) {
	return r.api.AddChannelMember(channelID, userID)
}

func (r *ChannelRepo) GetDirect(userID string) (*model.Channel, *model.AppError) {
	if r.botUserID == "" {
		panic("no bot user id")
	}

	return r.api.GetDirectChannel(userID, r.botUserID)
}
