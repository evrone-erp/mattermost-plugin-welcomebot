package command

import (
	"errors"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/presenter"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type GetPersonalChanelWelcomeMessage struct{}

const getPersonalChanelWelcomeHelp = "Display the personal welcome message set for the **current channel** (if any)"

func (c *GetPersonalChanelWelcomeMessage) Trigger() string {
	return "get_personal_channel_welcome_message"
}

func (c *GetPersonalChanelWelcomeMessage) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.IsSysadmin(args.UserId) || p.CanManageChannel(args.UserId, args.ChannelId)
}

func (c *GetPersonalChanelWelcomeMessage) Help() string {
	return "`/welcomebot get_personal_channel_welcome_message` - " + getPersonalChanelWelcomeHelp
}

func (c *GetPersonalChanelWelcomeMessage) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	cmd := command.GetPersonalChanelWelcomeMessage{
		CommandMessenger:        p.Container().NewCommandMessenger(args),
		ChannelWelcomeRepo:      p.Container().ChannelWelcomeRepo(),
		WelcomeMessagePresenter: &presenter.WelcomeMessagePresenter{UserRepo: p.Container().UserRepo()},
	}

	cmd.Call(args.UserId, args.ChannelId)
}

func (c *GetPersonalChanelWelcomeMessage) Validate(parameters []string) error {
	if len(parameters) > 0 {
		return errors.New("`get_personal_channel_welcome_message` command does not accept any extra parameters")
	}

	return nil
}

func (c *GetPersonalChanelWelcomeMessage) AutocompleteData() *model.AutocompleteData {
	return model.NewAutocompleteData("get_personal_channel_welcome_message", "", getPersonalChanelWelcomeHelp)
}
