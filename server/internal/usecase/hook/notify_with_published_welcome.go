package hook

import (
	"time"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/utils"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/shared/mlog"
)

type NotifyWithPublishedWelcome struct {
	Messenger               usecase.Messenger
	ChannelWelcomeRepo      usecase.ChannelWelcomeRepo
	ChannelRepo             usecase.ChannelRepo
	WelcomeMessagePresenter usecase.WelcomeMessagePresenter
}

func (uc *NotifyWithPublishedWelcome) Call(channelMember *model.ChannelMember) {
	if channel, appErr := uc.ChannelRepo.Get(channelMember.ChannelId); appErr != nil {
		mlog.Error(
			"error occurred while checking the type of the chanel",
			mlog.String("channelId", channelMember.ChannelId),
			mlog.Err(appErr),
		)
		return
	} else if !utils.IsChannelWithWelcomeSupport(channel) {
		return
	}

	welcome, appErr := uc.ChannelWelcomeRepo.GetPublishedChanelWelcome(channelMember.ChannelId)
	if appErr != nil {
		mlog.Error(
			"error occurred while retrieving the welcome message",
			mlog.String("channelId", channelMember.ChannelId),
			mlog.Err(appErr),
		)
		return
	}

	if welcome == nil {
		return
	}

	message, err := uc.WelcomeMessagePresenter.Render(welcome.Message, channelMember.UserId)

	if err != nil {
		mlog.Error(
			"Error while rendering message %s: %s",
			mlog.String("UserId", channelMember.UserId),
			mlog.Err(err),
		)
		return
	}

	time.Sleep(1 * time.Second)
	appErr = uc.Messenger.Post(channelMember.ChannelId, message)

	if appErr != nil {
		mlog.Error(
			"error ocuring while posting message to the channel",
			mlog.String("channelId", channelMember.ChannelId),
			mlog.Err(appErr),
		)
		return
	}
}
