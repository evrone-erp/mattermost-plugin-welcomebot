package command

import (
	"errors"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type DeletePublishedChanelWelcome struct{}

func (c *DeletePublishedChanelWelcome) Trigger() string {
	return "delete_published_channel_welcome"
}

func (c *DeletePublishedChanelWelcome) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.IsSysadmin(args.UserId) || p.CanManageChannel(args.UserId, args.ChannelId)
}

func (c *DeletePublishedChanelWelcome) Help() string {
	return "`/welcomebot delete_published_channel_welcome` - delete the published welcome message for the given channel (if any)"
}

func (c *DeletePublishedChanelWelcome) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	cmd := command.DeletePublishedChanelWelcome{
		CommandMessenger:   p.Container().NewCommandMessenger(args),
		ChannelWelcomeRepo: p.Container().ChannelWelcomeRepo(),
	}

	cmd.Call(args.ChannelId)
}

func (c *DeletePublishedChanelWelcome) Validate(parameters []string) error {
	if len(parameters) > 0 {
		return errors.New("`delete_published_channel_welcome` command does not accept any extra parameters")
	}

	return nil
}

func (c *DeletePublishedChanelWelcome) AutocompleteData() *model.AutocompleteData {
	return model.NewAutocompleteData("delete_published_channel_welcome", "", "Delete the welcome message for the channel")
}
