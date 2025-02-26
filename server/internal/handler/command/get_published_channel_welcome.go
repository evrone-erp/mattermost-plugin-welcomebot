package command

import (
	"errors"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/presenter"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type GetPublishedChanelWelcome struct{}

func (c *GetPublishedChanelWelcome) Trigger() string {
	return "get_published_channel_welcome"
}

func (c *GetPublishedChanelWelcome) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.IsSysadmin(args.UserId) || p.CanManageChannel(args.UserId, args.ChannelId)
}

func (c *GetPublishedChanelWelcome) Help() string {
	return "`/welcomebot get_published_channel_welcome` - print the published welcome message set for the given channel (if any)"
}

func (c *GetPublishedChanelWelcome) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	cmd := command.GetPublishedChanelWelcome{
		CommandMessenger:        p.Container().NewCommandMessenger(args),
		ChannelWelcomeRepo:      p.Container().ChannelWelcomeRepo(),
		WelcomeMessagePresenter: &presenter.WelcomeMessagePresenter{UserRepo: p.Container().UserRepo()},
	}

	cmd.Call(args.UserId, args.ChannelId)
}

func (c *GetPublishedChanelWelcome) Validate(parameters []string) error {
	if len(parameters) > 0 {
		return errors.New("`get_published_channel_welcome` command does not accept any extra parameters")
	}

	return nil
}

func (c *GetPublishedChanelWelcome) AutocompleteData() *model.AutocompleteData {
	return model.NewAutocompleteData("get_published_channel_welcome", "", "Print the welcome message set for the channel")
}
