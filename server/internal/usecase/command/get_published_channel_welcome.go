package command

import (
	"fmt"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
)

type GetPublishedChanelWelcome struct {
	CommandMessenger        usecase.CommandMessenger
	ChannelWelcomeRepo      usecase.ChannelWelcomeRepo
	WelcomeMessagePresenter usecase.WelcomeMessagePresenter
}

func (uc *GetPublishedChanelWelcome) Call(currentUserID string, channelID string) {
	welcome, appErr := uc.ChannelWelcomeRepo.GetPublishedChanelWelcome(channelID)

	if appErr != nil {
		message := fmt.Sprintf("error occurred while retrieving the welcome message for the chanel: `%s`", appErr)
		uc.CommandMessenger.PostCommandResponse(message)
		return
	}

	if welcome == nil {
		uc.CommandMessenger.PostCommandResponse("welcome message has not been set yet")
		return
	}

	message, appErr := uc.WelcomeMessagePresenter.Render(welcome.Message, currentUserID)

	if appErr != nil {
		response := fmt.Sprintf("Error while rendering message %s: %s", currentUserID, appErr)
		uc.CommandMessenger.PostCommandResponse(response)
		return
	}

	message = fmt.Sprintf("Welcome message is:\n%s", message)
	uc.CommandMessenger.PostCommandResponse(message)
}
