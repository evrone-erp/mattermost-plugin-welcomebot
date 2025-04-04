package command

import (
	"errors"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/presenter"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type GetTeamWelcomeSettings struct{}

const getTeamWelcomeSettingsHelp = "Display the welcome settings set for the **current team**"

func (c *GetTeamWelcomeSettings) Trigger() string {
	return "get_team_welcome_settings"
}

func (c *GetTeamWelcomeSettings) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.CanManageTeam(args.UserId, args.TeamId)
}

func (c *GetTeamWelcomeSettings) Help() string {
	return "`/welcomebot get_team_welcome_settings` - " + getTeamWelcomeSettingsHelp
}

func (c *GetTeamWelcomeSettings) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	cmd := command.GetTeamWelcomeSettings{
		CommandMessenger:        p.Container().NewCommandMessenger(args),
		TeamWelcomeRepo:         p.Container().TeamWelcomeRepo(),
		ChannelRepo:             p.Container().ChannelRepo(),
		WelcomeMessagePresenter: &presenter.WelcomeMessagePresenter{UserRepo: p.Container().UserRepo()},
	}

	cmd.Call(args.UserId, args.TeamId)
}

func (c *GetTeamWelcomeSettings) Validate(parameters []string) error {
	if len(parameters) != 0 {
		return errors.New("`get_team_welcome_settings` command does not accept any args")
	}

	return nil
}

func (c *GetTeamWelcomeSettings) AutocompleteData() *model.AutocompleteData {
	data := model.NewAutocompleteData("get_team_welcome_settings", "", getTeamWelcomeSettingsHelp)

	return data
}
