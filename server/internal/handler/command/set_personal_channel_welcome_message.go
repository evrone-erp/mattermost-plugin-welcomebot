package command

import (
	"errors"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type SetPersonalChanelWelcomeMessage struct{}

const setPersonalChanelWelcomeMessageHelp = "Set a personal welcome message for the **current channel** (Direct channels are not supported)"

func (c *SetPersonalChanelWelcomeMessage) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.IsSysadmin(args.UserId) || p.CanManageChannel(args.UserId, args.ChannelId)
}

func (c *SetPersonalChanelWelcomeMessage) Trigger() string {
	return "set_personal_channel_welcome_message"
}

func (c *SetPersonalChanelWelcomeMessage) Help() string {
	return "`/welcomebot set_personal_channel_welcome_message [welcome-message]` - " + setPersonalChanelWelcomeMessageHelp
}

func (c *SetPersonalChanelWelcomeMessage) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	cmd := command.SetPersonalChanelWelcomeMessage{
		CommandMessenger:   p.Container().NewCommandMessenger(args),
		ChannelWelcomeRepo: p.Container().ChannelWelcomeRepo(),
		ChannelRepo:        p.Container().ChannelRepo(),
	}

	cmd.Call(args.Command, args.ChannelId)
}

func (c *SetPersonalChanelWelcomeMessage) Validate(parameters []string) error {
	if len(parameters) == 0 {
		return errors.New("`set_personal_channel_welcome_message` command requires the message to be provided")
	}

	return nil
}

func (c *SetPersonalChanelWelcomeMessage) AutocompleteData() *model.AutocompleteData {
	data := model.NewAutocompleteData("set_personal_channel_welcome_message", "[welcome-message]", setPersonalChanelWelcomeMessageHelp)
	data.AddTextArgument("WelcomeMessage message for the channel", "[welcome-message]", "")

	return data
}
