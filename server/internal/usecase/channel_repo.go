package usecase

import (
	"github.com/mattermost/mattermost/server/public/model"
)

type ChannelRepo interface {
	Get(channelID string) (*model.Channel, *model.AppError)
	GetDirect(userID string) (*model.Channel, *model.AppError)
	GetByName(teamID string, channelName string) (*model.Channel, *model.AppError)
	AddMemberByUserID(channelID, userID string) (*model.ChannelMember, *model.AppError)
}
