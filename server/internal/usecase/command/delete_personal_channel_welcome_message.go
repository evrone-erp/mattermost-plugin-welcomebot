package command

import (
	"fmt"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
)

type DeletePersonalChanelWelcomeMessage struct {
	CommandMessenger   usecase.CommandMessenger
	ChannelWelcomeRepo usecase.ChannelWelcomeRepo
}

func (uc *DeletePersonalChanelWelcomeMessage) Call(channelID string) {
	welcome, err := uc.ChannelWelcomeRepo.GetPersonalChanelWelcome(channelID)

	if err != nil {
		msg := fmt.Sprintf("error occurred while retrieving the welcome message for the chanel: `%s`", err)
		uc.CommandMessenger.PostCommandResponse(msg)
		return
	}

	if welcome == nil {
		uc.CommandMessenger.PostCommandResponse("welcome message has not been set yet")
		return
	}

	if err := uc.ChannelWelcomeRepo.DeletePersonalChanelWelcome(channelID); err != nil {
		msg := fmt.Sprintf("error occurred while deleting the welcome message for the chanel: `%s`", err)
		uc.CommandMessenger.PostCommandResponse(msg)

		return
	}

	uc.CommandMessenger.PostCommandResponse("welcome message has been deleted")
}
