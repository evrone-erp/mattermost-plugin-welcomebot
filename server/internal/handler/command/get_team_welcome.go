package command

import (
	"errors"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/presenter"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type GetTeamWelcome struct{}

func (c *GetTeamWelcome) Trigger() string {
	return "get_team_welcome"
}

func (c *GetTeamWelcome) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.IsSysadmin(args.UserId)
}

func (c *GetTeamWelcome) Help() string {
	return "`/welcomebot get_team_welcome` - Get welcome message after joining the team"
}

func (c *GetTeamWelcome) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	cmd := command.GetTeamWelcome{
		CommandMessenger:        p.Container().NewCommandMessenger(args),
		TeamWelcomeRepo:         p.Container().TeamWelcomeRepo(),
		ChannelRepo:             p.Container().ChannelRepo(),
		WelcomeMessagePresenter: &presenter.WelcomeMessagePresenter{UserRepo: p.Container().UserRepo()},
	}

	cmd.Call(args.UserId, args.TeamId)
}

func (c *GetTeamWelcome) Validate(parameters []string) error {
	if len(parameters) != 0 {
		return errors.New("`get_team_welcome` command does not accept any args")
	}

	return nil
}

func (c *GetTeamWelcome) AutocompleteData() *model.AutocompleteData {
	data := model.NewAutocompleteData("get_team_welcome", "", "displays information about team welcome")

	return data
}
