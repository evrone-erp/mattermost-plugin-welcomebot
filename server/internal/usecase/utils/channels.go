package utils

import (
	"github.com/mattermost/mattermost/server/public/model"
)

func IsChannelWithWelcomeSupport(ch *model.Channel) bool {
	return ch.Type == model.ChannelTypePrivate || ch.Type == model.ChannelTypeOpen
}
