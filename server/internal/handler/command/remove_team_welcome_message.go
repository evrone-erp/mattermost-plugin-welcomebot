package command

import (
	"errors"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type RemoveTeamWelcomeMessage struct{}

func (c *RemoveTeamWelcomeMessage) Trigger() string {
	return "remove_team_welcome"
}

func (c *RemoveTeamWelcomeMessage) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.IsSysadmin(args.UserId)
}

func (c *RemoveTeamWelcomeMessage) Help() string {
	return "`/welcomebot remove_team_welcome` - Remove welcome message after joining the team"
}

func (c *RemoveTeamWelcomeMessage) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	cmd := command.RemoveTeamWelcomeMessage{
		CommandMessenger: p.Container().NewCommandMessenger(args),
		TeamWelcomeRepo:  p.Container().TeamWelcomeRepo(),
	}

	cmd.Call(args.TeamId)
}

func (c *RemoveTeamWelcomeMessage) Validate(parameters []string) error {
	if len(parameters) != 0 {
		return errors.New("`remove_team_welcome` command does not accept any args")
	}

	return nil
}

func (c *RemoveTeamWelcomeMessage) AutocompleteData() *model.AutocompleteData {
	data := model.NewAutocompleteData("remove_team_welcome", "", "removes team welcome message")

	return data
}
