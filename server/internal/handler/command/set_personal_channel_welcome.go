package command

import (
	"errors"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type SetPersonalChanelWelcome struct{}

func (c *SetPersonalChanelWelcome) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.IsSysadmin(args.UserId) || p.CanManageChannel(args.UserId, args.ChannelId)
}

func (c *SetPersonalChanelWelcome) Trigger() string {
	return "set_personal_channel_welcome"
}

func (c *SetPersonalChanelWelcome) Help() string {
	return "`/welcomebot set_personal_channel_welcome [welcome-message]` - set the personal welcome message for the given channel. Direct channels are not supported."
}

func (c *SetPersonalChanelWelcome) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	cmd := command.SetPersonalChanelWelcome{
		CommandMessenger:   p.Container().NewCommandMessenger(args),
		ChannelWelcomeRepo: p.Container().ChannelWelcomeRepo(),
		ChannelRepo:        p.Container().ChannelRepo(),
	}

	cmd.Call(args.Command, args.ChannelId)
}

func (c *SetPersonalChanelWelcome) Validate(parameters []string) error {
	if len(parameters) == 0 {
		return errors.New("`set_personal_channel_welcome` command requires the message to be provided")
	}

	return nil
}

func (c *SetPersonalChanelWelcome) AutocompleteData() *model.AutocompleteData {
	data := model.NewAutocompleteData("set_personal_channel_welcome", "[welcome-message]", "Set the welcome message for the channel")
	data.AddTextArgument("Welcome message for the channel", "[welcome-message]", "")

	return data
}
