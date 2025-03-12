package hook

import (
	"time"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/utils"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/shared/mlog"
)

type NotifyWithPersonalWelcome struct {
	Messenger               usecase.Messenger
	ChannelWelcomeRepo      usecase.ChannelWelcomeRepo
	ChannelRepo             usecase.ChannelRepo
	WelcomeMessagePresenter usecase.WelcomeMessagePresenter
}

func (uc *NotifyWithPersonalWelcome) Call(channelMember *model.ChannelMember) {
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

	welcome, appErr := uc.ChannelWelcomeRepo.GetPersonalChanelWelcome(channelMember.ChannelId)
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

	dmChannel, err := uc.ChannelRepo.GetDirect(channelMember.UserId)
	if err != nil {
		mlog.Error(
			"error occurred while creating direct channel to the user",
			mlog.String("UserId", channelMember.UserId),
			mlog.Err(err),
		)
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

	if appErr := uc.Messenger.PostDirect(dmChannel.Id, message); appErr != nil {
		mlog.Error("failed to post welcome message to the channel",
			mlog.String("channelId", dmChannel.Id),
			mlog.Err(appErr),
		)
	}

	time.Sleep(1 * time.Second)
	uc.Messenger.PostChannelEphemeral(channelMember.ChannelId, channelMember.UserId, message)
}
