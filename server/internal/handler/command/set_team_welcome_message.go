package command

import (
	"errors"
	"strings"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type SetTeamWelcomeMessage struct{}

func (c *SetTeamWelcomeMessage) Trigger() string {
	return "set_team_welcome_message"
}

func (c *SetTeamWelcomeMessage) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.IsSysadmin(args.UserId)
}

func (c *SetTeamWelcomeMessage) Help() string {
	return "`/welcomebot set_team_welcome_message [team_name] [message]` - Get welcome message after joining the team"
}

func (c *SetTeamWelcomeMessage) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	msgr := p.Container().NewCommandMessenger(args)

	split := strings.SplitN(args.Command, " ", 3)

	if len(split) < 3 {
		msgr.PostCommandResponse("Error while parsing arguments ")
		return
	}

	message := strings.TrimSpace(split[2])

	cmd := command.SetTeamWelcomeMessage{
		Messenger:       msgr,
		TeamWelcomeRepo: p.Container().TeamWelcomeRepo(),
	}

	cmd.Call(args.TeamId, message)
}

func (c *SetTeamWelcomeMessage) Validate(parameters []string) error {
	if len(parameters) < 1 {
		return errors.New("`set_team_welcome_message` command requires the team_name to be provided and a message")
	}

	return nil
}

func (c *SetTeamWelcomeMessage) AutocompleteData() *model.AutocompleteData {
	data := model.NewAutocompleteData("set_team_welcome_message", "[message]", "set team welcome message")
	data.AddTextArgument("Message", "[message]", "")

	return data
}
