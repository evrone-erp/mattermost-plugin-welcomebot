package command

import (
	"errors"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/presenter"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type GetPublishedChanelWelcomeMessage struct{}

const getPublishedChanelWelcomeHelp = "Display the published welcome message set for the **current channel** (if any)"

func (c *GetPublishedChanelWelcomeMessage) Trigger() string {
	return "get_published_channel_welcome_message"
}

func (c *GetPublishedChanelWelcomeMessage) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.IsSysadmin(args.UserId) || p.CanManageChannel(args.UserId, args.ChannelId)
}

func (c *GetPublishedChanelWelcomeMessage) Help() string {
	return "`/welcomebot get_published_channel_welcome_message` - " + getPublishedChanelWelcomeHelp
}

func (c *GetPublishedChanelWelcomeMessage) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	cmd := command.GetPublishedChanelWelcomeMessage{
		CommandMessenger:        p.Container().NewCommandMessenger(args),
		ChannelWelcomeRepo:      p.Container().ChannelWelcomeRepo(),
		WelcomeMessagePresenter: &presenter.WelcomeMessagePresenter{UserRepo: p.Container().UserRepo()},
	}

	cmd.Call(args.UserId, args.ChannelId)
}

func (c *GetPublishedChanelWelcomeMessage) Validate(parameters []string) error {
	if len(parameters) > 0 {
		return errors.New("`get_published_channel_welcome_message` command does not accept any extra parameters")
	}

	return nil
}

func (c *GetPublishedChanelWelcomeMessage) AutocompleteData() *model.AutocompleteData {
	return model.NewAutocompleteData("get_published_channel_welcome_message", "", getPublishedChanelWelcomeHelp)
}
