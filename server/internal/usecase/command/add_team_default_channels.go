package command

import (
	"fmt"
	"slices"
	"sort"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
)

type AddTeamDefaultChannels struct {
	CommandMessenger usecase.CommandMessenger
	TeamWelcomeRepo  usecase.TeamWelcomeRepo
	ChannelRepo      usecase.ChannelRepo
}

func (uc *AddTeamDefaultChannels) Call(teamID string, channelNames []string) {
	currentWelcome, appErr := uc.TeamWelcomeRepo.GetTeamWelcome(teamID)
	if appErr != nil {
		// if something went wrong, notify about this and try to rewrite data further
		response := fmt.Sprintf("Error while fetching current welcome %s: %s, trying to rewrite...", teamID, appErr)
		uc.CommandMessenger.PostCommandResponse(response)
	}

	if currentWelcome == nil {
		currentWelcome = &pmodel.TeamWelcome{}
	}

	channelIDs := make([]string, 0, len(currentWelcome.ChannelIDs))

	for _, id := range currentWelcome.ChannelIDs {
		channel, appErr := uc.ChannelRepo.Get(id)

		if appErr == nil {
			channelIDs = append(channelIDs, channel.Id)
		} else {
			response := fmt.Sprintf("Dropping previously saved channel#id %s", id)
			uc.CommandMessenger.PostCommandResponse(response)
		}
	}

	for _, name := range channelNames {
		channel, appErr := uc.ChannelRepo.GetByName(teamID, name)

		if appErr == nil {
			channelIDs = append(channelIDs, channel.Id)
		} else {
			response := fmt.Sprintf("Skipping unknown channel %s", name)
			uc.CommandMessenger.PostCommandResponse(response)
		}
	}

	sort.Strings(channelIDs)
	channelIDs = slices.Compact(channelIDs)
	currentWelcome.ChannelIDs = channelIDs

	appErr = uc.TeamWelcomeRepo.SetTeamWelcome(teamID, currentWelcome)

	if appErr != nil {
		response := fmt.Sprintf("Error while saving team %s: %s", teamID, appErr)
		uc.CommandMessenger.PostCommandResponse(response)
		return
	}

	uc.CommandMessenger.PostCommandResponse("Welcome channels were updated")
}
