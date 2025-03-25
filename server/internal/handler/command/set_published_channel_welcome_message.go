package command

import (
	"errors"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type SetPublishedChanelWelcomeMessage struct{}

const setPublishedChanelWelcomeMessageHelp = "Set a published welcome message for the **current channel** (Direct channels are not supported)"

func (c *SetPublishedChanelWelcomeMessage) Trigger() string {
	return "set_published_channel_welcome_message"
}

func (c *SetPublishedChanelWelcomeMessage) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.IsSysadmin(args.UserId) || p.CanManageChannel(args.UserId, args.ChannelId)
}

func (c *SetPublishedChanelWelcomeMessage) Help() string {
	return "`/welcomebot set_published_channel_welcome_message [welcome-message]` - " + setPublishedChanelWelcomeMessageHelp
}

func (c *SetPublishedChanelWelcomeMessage) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	cmd := command.SetPublishedChanelWelcomeMessage{
		CommandMessenger:   p.Container().NewCommandMessenger(args),
		ChannelWelcomeRepo: p.Container().ChannelWelcomeRepo(),
		ChannelRepo:        p.Container().ChannelRepo(),
	}

	cmd.Call(args.Command, args.ChannelId)
}

func (c *SetPublishedChanelWelcomeMessage) Validate(parameters []string) error {
	if len(parameters) == 0 {
		return errors.New("`set_published_channel_welcome_message` command requires the message to be provided")
	}

	return nil
}

func (c *SetPublishedChanelWelcomeMessage) AutocompleteData() *model.AutocompleteData {
	data := model.NewAutocompleteData("set_published_channel_welcome_message", "[welcome-message]", setPublishedChanelWelcomeMessageHelp)
	data.AddTextArgument("WelcomeMessage message for the channel", "[welcome-message]", "")

	return data
}
