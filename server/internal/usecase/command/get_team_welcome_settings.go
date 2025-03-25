package command

import (
	"fmt"
	"strings"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
)

type GetTeamWelcomeSettings struct {
	CommandMessenger        usecase.CommandMessenger
	TeamWelcomeRepo         usecase.TeamWelcomeRepo
	ChannelRepo             usecase.ChannelRepo
	WelcomeMessagePresenter usecase.WelcomeMessagePresenter
}

func (uc *GetTeamWelcomeSettings) Call(userID string, teamID string) {
	currentWelcome, appErr := uc.TeamWelcomeRepo.GetTeamWelcome(teamID)

	if appErr != nil {
		response := fmt.Sprintf("Error while fetching current welcome %s: %s, you may try to overwrite it", teamID, appErr)
		uc.CommandMessenger.PostCommandResponse(response)
		return
	}

	var response strings.Builder

	if currentWelcome == nil || currentWelcome.Message == "" {
		response.WriteString("No welcome message\n\n\n")
	} else {
		message, appErr := uc.WelcomeMessagePresenter.Render(currentWelcome.Message, userID)

		if appErr != nil {
			response := fmt.Sprintf("Error while rendering message %s: %s", teamID, appErr)
			uc.CommandMessenger.PostCommandResponse(response)
			return
		}

		response.WriteString(fmt.Sprintf("Current welcome is:\n%s\n\n\n", message))
	}

	if currentWelcome == nil {
		response.WriteString("No default channels\n")
		uc.CommandMessenger.PostCommandResponse(response.String())

		return
	}

	channelNames := make([]string, 0, len(currentWelcome.ChannelIDs))

	for _, id := range currentWelcome.ChannelIDs {
		channel, appErrr := uc.ChannelRepo.Get(id)

		if appErrr == nil {
			channelNames = append(channelNames, channel.Name)
		}
	}

	if len(channelNames) > 0 {
		response.WriteString("Default channels:\n")
		for _, name := range channelNames {
			response.WriteString(fmt.Sprintf("- ~%s\n", name))
		}
	} else {
		response.WriteString("No default channels\n")
	}

	uc.CommandMessenger.PostCommandResponse(response.String())
}
