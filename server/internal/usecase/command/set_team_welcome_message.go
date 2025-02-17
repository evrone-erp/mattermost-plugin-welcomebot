package command

import (
	"fmt"
	"strings"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
)

type SetTeamWelcomeMessage struct {
	Messenger       usecase.CommandMessenger
	TeamWelcomeRepo usecase.TeamWelcomeRepo
}

func (uc *SetTeamWelcomeMessage) Call(teamID string, message string) {
	currentWelcome, appErr := uc.TeamWelcomeRepo.GetTeamWelcome(teamID)

	if appErr != nil {
		// if something went wrong, notify about this and try to rewrite data further
		response := fmt.Sprintf("Error while fetching current welcome %s: %s, trying to rewrite...", teamID, appErr)
		uc.Messenger.PostCommandResponse(response)
	}

	if currentWelcome == nil {
		currentWelcome = &pmodel.TeamWelcome{}
	}

	currentWelcome.Message = strings.TrimSpace(message)
	appErr = uc.TeamWelcomeRepo.SetTeamWelcome(teamID, currentWelcome)

	if appErr != nil {
		response := fmt.Sprintf("Error while saving team %s: %s", teamID, appErr)
		uc.Messenger.PostCommandResponse(response)
		return
	}

	uc.Messenger.PostCommandResponse("Welcome message was updated")
}
