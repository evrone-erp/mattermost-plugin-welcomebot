package command

import (
	"fmt"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
)

type DeletePublishedChanelWelcomeMessage struct {
	CommandMessenger   usecase.CommandMessenger
	ChannelWelcomeRepo usecase.ChannelWelcomeRepo
}

func (uc *DeletePublishedChanelWelcomeMessage) Call(channelID string) {
	welcome, err := uc.ChannelWelcomeRepo.GetPublishedChanelWelcome(channelID)

	if err != nil {
		msg := fmt.Sprintf("error occurred while retrieving the welcome message for the chanel: `%s`", err)
		uc.CommandMessenger.PostCommandResponse(msg)
		return
	}

	if welcome == nil {
		uc.CommandMessenger.PostCommandResponse("welcome message has not been set yet")
		return
	}

	if err := uc.ChannelWelcomeRepo.DeletePublishedChanelWelcome(channelID); err != nil {
		msg := fmt.Sprintf("error occurred while deleting the welcome message for the chanel: `%s`", err)
		uc.CommandMessenger.PostCommandResponse(msg)

		return
	}

	uc.CommandMessenger.PostCommandResponse("welcome message has been deleted")
}
