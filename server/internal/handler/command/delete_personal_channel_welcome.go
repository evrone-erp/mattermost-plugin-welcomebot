package command

import (
	"errors"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type DeletePersonalChanelWelcome struct{}

func (c *DeletePersonalChanelWelcome) Trigger() string {
	return "delete_personal_channel_welcome"
}

func (c *DeletePersonalChanelWelcome) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.IsSysadmin(args.UserId) || p.CanManageChannel(args.UserId, args.ChannelId)
}

func (c *DeletePersonalChanelWelcome) Help() string {
	return "`/welcomebot delete_personal_channel_welcome` - delete the personal welcome message for the given channel (if any)"
}

func (c *DeletePersonalChanelWelcome) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	cmd := command.DeletePersonalChanelWelcome{
		CommandMessenger:   p.Container().NewCommandMessenger(args),
		ChannelWelcomeRepo: p.Container().ChannelWelcomeRepo(),
	}

	cmd.Call(args.ChannelId)
}

func (c *DeletePersonalChanelWelcome) Validate(parameters []string) error {
	if len(parameters) > 0 {
		return errors.New("`delete_personal_channel_welcome` command does not accept any extra parameters")
	}

	return nil
}

func (c *DeletePersonalChanelWelcome) AutocompleteData() *model.AutocompleteData {
	return model.NewAutocompleteData("delete_personal_channel_welcome", "", "Delete the welcome message for the channel")
}
