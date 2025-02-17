package command

import (
	"fmt"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
)

type RemoveTeamWelcomeMessage struct {
	CommandMessenger usecase.CommandMessenger
	TeamWelcomeRepo  usecase.TeamWelcomeRepo
}

func (uc *RemoveTeamWelcomeMessage) Call(teamID string) {
	currentWelcome, appErr := uc.TeamWelcomeRepo.GetTeamWelcome(teamID)

	if appErr != nil {
		response := fmt.Sprintf("Error while fetching current welcome %s: %s, trying to rewrite...", teamID, appErr)
		uc.CommandMessenger.PostCommandResponse(response)
	}

	if currentWelcome == nil {
		currentWelcome = &pmodel.TeamWelcome{}
	}

	currentWelcome.Message = ""
	appErr = uc.TeamWelcomeRepo.SetTeamWelcome(teamID, currentWelcome)

	if appErr != nil {
		response := fmt.Sprintf("Error while saving team %s: %s", teamID, appErr)
		uc.CommandMessenger.PostCommandResponse(response)
		return
	}

	uc.CommandMessenger.PostCommandResponse("Welcome message was removed")
}
