package command

import (
	"fmt"
	"strings"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/utils"
)

type SetPersonalChanelWelcome struct {
	CommandMessenger   usecase.CommandMessenger
	ChannelWelcomeRepo usecase.ChannelWelcomeRepo
	ChannelRepo        usecase.ChannelRepo
}

func (uc *SetPersonalChanelWelcome) Call(fullCommand string, channelID string) {
	channel, appErr := uc.ChannelRepo.Get(channelID)
	if appErr != nil {
		response := fmt.Sprintf("error occurred while checking the type of the chanelId `%s`: `%s`", channelID, appErr)
		uc.CommandMessenger.PostCommandResponse(response)
		return
	}

	if !utils.IsChannelWithWelcomeSupport(channel) {
		uc.CommandMessenger.PostCommandResponse("Channel type is not supported")
		return
	}

	parsedCommand := strings.SplitN(fullCommand, "set_personal_channel_welcome", 2)

	if len(parsedCommand) != 2 {
		response := fmt.Sprintf("error ocured while parsing command %s", fullCommand)
		uc.CommandMessenger.PostCommandResponse(response)

		return
	}

	message := parsedCommand[1]
	message = strings.TrimSpace(message)

	if message == "" {
		uc.CommandMessenger.PostCommandResponse("unable to store empty message")
		return
	}

	if appErr := uc.ChannelWelcomeRepo.SetPersonalChanelWelcome(channel.Id, message); appErr != nil {
		response := fmt.Sprintf("error occurred while storing the welcome message for the chanel: `%s`", appErr)
		uc.CommandMessenger.PostCommandResponse(response)
		return
	}

	response := fmt.Sprintf("stored the welcome message:\n%s", message)
	uc.CommandMessenger.PostCommandResponse(response)
}
