package command

import (
	"fmt"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
)

type RemoveTeamDefaultChannels struct {
	CommandMessenger usecase.CommandMessenger
	TeamWelcomeRepo  usecase.TeamWelcomeRepo
	ChannelRepo      usecase.ChannelRepo
}

func (uc *RemoveTeamDefaultChannels) Call(teamID string, channelNames []string) {
	currentWelcome, appErr := uc.TeamWelcomeRepo.GetTeamWelcome(teamID)
	if appErr != nil {
		response := fmt.Sprintf("Error while fetching current welcome %s: %s", teamID, appErr)
		uc.CommandMessenger.PostCommandResponse(response)
		return
	}

	if currentWelcome == nil {
		uc.CommandMessenger.PostCommandResponse("There is nothing to update")
		return
	}

	activeChannels := make(map[string]bool)

	for _, id := range currentWelcome.ChannelIDs {
		_, appErr := uc.ChannelRepo.Get(id)

		if appErr == nil {
			activeChannels[id] = true
		} else {
			response := fmt.Sprintf("Dropping previously saved broken channel#id %s", id)
			uc.CommandMessenger.PostCommandResponse(response)
		}
	}

	for _, name := range channelNames {
		channel, appErr := uc.ChannelRepo.GetByName(teamID, name)

		if appErr == nil {
			if _, wasSaved := activeChannels[channel.Id]; wasSaved {
				activeChannels[channel.Id] = false
			} else {
				response := fmt.Sprintf("Ignoring not configured channel %s", name)
				uc.CommandMessenger.PostCommandResponse(response)
			}
		} else {
			response := fmt.Sprintf("Error while checking channel %s:%s. Ignoring", name, appErr)
			uc.CommandMessenger.PostCommandResponse(response)
		}
	}

	resultChannels := make([]string, 0, len(activeChannels))
	for id, active := range activeChannels {
		if active {
			resultChannels = append(resultChannels, id)
		}
	}

	currentWelcome.ChannelIDs = resultChannels

	appErr = uc.TeamWelcomeRepo.SetTeamWelcome(teamID, currentWelcome)

	if appErr != nil {
		response := fmt.Sprintf("Error while saving team %s: %s", teamID, appErr)
		uc.CommandMessenger.PostCommandResponse(response)
		return
	}

	uc.CommandMessenger.PostCommandResponse("Welcome channels were updated")
}
