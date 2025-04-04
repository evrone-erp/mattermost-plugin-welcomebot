package command

import (
	"errors"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type DeleteTeamWelcomeMessage struct{}

const deleteTeamWelcomeMessageHelp = "Delete the welcome message for the **current team**"

func (c *DeleteTeamWelcomeMessage) Trigger() string {
	return "delete_team_welcome_message"
}

func (c *DeleteTeamWelcomeMessage) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.CanManageTeam(args.UserId, args.TeamId)
}

func (c *DeleteTeamWelcomeMessage) Help() string {
	return "`/welcomebot delete_team_welcome_message` - " + deleteTeamWelcomeMessageHelp
}

func (c *DeleteTeamWelcomeMessage) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	cmd := command.DeleteTeamWelcomeMessage{
		CommandMessenger: p.Container().NewCommandMessenger(args),
		TeamWelcomeRepo:  p.Container().TeamWelcomeRepo(),
	}

	cmd.Call(args.TeamId)
}

func (c *DeleteTeamWelcomeMessage) Validate(parameters []string) error {
	if len(parameters) != 0 {
		return errors.New("`delete_team_welcome_message` command does not accept any args")
	}

	return nil
}

func (c *DeleteTeamWelcomeMessage) AutocompleteData() *model.AutocompleteData {
	data := model.NewAutocompleteData("delete_team_welcome_message", "", deleteTeamWelcomeMessageHelp)

	return data
}
