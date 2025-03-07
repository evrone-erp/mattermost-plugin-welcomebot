package command

import (
	"errors"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type SetPublishedChanelWelcome struct{}

func (c *SetPublishedChanelWelcome) Trigger() string {
	return "set_published_channel_welcome"
}

func (c *SetPublishedChanelWelcome) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.IsSysadmin(args.UserId) || p.CanManageChannel(args.UserId, args.ChannelId)
}

func (c *SetPublishedChanelWelcome) Help() string {
	return "`/welcomebot set_published_channel_welcome [welcome-message]` - set the published welcome message for the given channel. Direct channels are not supported."
}

func (c *SetPublishedChanelWelcome) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	cmd := command.SetPublishedChanelWelcome{
		CommandMessenger:   p.Container().NewCommandMessenger(args),
		ChannelWelcomeRepo: p.Container().ChannelWelcomeRepo(),
		ChannelRepo:        p.Container().ChannelRepo(),
	}

	cmd.Call(args.Command, args.ChannelId)
}

func (c *SetPublishedChanelWelcome) Validate(parameters []string) error {
	if len(parameters) == 0 {
		return errors.New("`set_published_channel_welcome` command requires the message to be provided")
	}

	return nil
}

func (c *SetPublishedChanelWelcome) AutocompleteData() *model.AutocompleteData {
	data := model.NewAutocompleteData("set_published_channel_welcome", "[welcome-message]", "Set the welcome message for the channel")
	data.AddTextArgument("Welcome message for the channel", "[welcome-message]", "")

	return data
}
