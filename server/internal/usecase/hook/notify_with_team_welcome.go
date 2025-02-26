package hook

import (
	"time"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/shared/mlog"
)

type NotifyWithTeamWelcome struct {
	Messenger               usecase.Messenger
	TeamWelcomeRepo         usecase.TeamWelcomeRepo
	ChannelRepo             usecase.ChannelRepo
	WelcomeMessagePresenter usecase.WelcomeMessagePresenter
}

func (uc *NotifyWithTeamWelcome) Call(member *model.TeamMember) {
	welcome, appErr := uc.TeamWelcomeRepo.GetTeamWelcome(member.TeamId)

	if appErr != nil {
		mlog.Error(
			"error occurred while checking welcome",
			mlog.String("teamID", member.TeamId),
			mlog.Err(appErr),
		)
	}

	if welcome == nil {
		return
	}

	for _, id := range welcome.ChannelIDs {
		_, err := uc.ChannelRepo.AddMemberByUserID(id, member.UserId)

		if err != nil {
			mlog.Error(
				"error occurred while joining user to a channel",
				mlog.String("UserID", member.UserId),
				mlog.String("ChannelID", id),
				mlog.Err(err),
			)
		}
	}

	if welcome.Message == "" {
		return
	}

	time.Sleep(1 * time.Second)

	dmChannel, err := uc.ChannelRepo.GetDirect(member.UserId)
	if err != nil {
		mlog.Error(
			"error occurred while creating direct channel to the user",
			mlog.String("UserID", member.UserId),
			mlog.Err(err),
		)
		return
	}

	message, err := uc.WelcomeMessagePresenter.Render(welcome.Message, member.UserId)

	if err != nil {
		mlog.Error(
			"Error while rendering message %s: %s",
			mlog.String("UserId", member.UserId),
			mlog.Err(err),
		)
		return
	}

	if appErr := uc.Messenger.PostDirect(dmChannel.Id, message); appErr != nil {
		mlog.Error("failed to post team welcome message to the userId",
			mlog.String("userID", member.UserId),
			mlog.Err(appErr),
		)
	}
}
