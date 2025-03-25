package command

import (
	"errors"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type DeletePersonalChanelWelcomeMessage struct{}

const deletePersonalChanelWelcomeMessageHelp = "Delete the personal welcome message for the **current channel** (if any)"

func (c *DeletePersonalChanelWelcomeMessage) Trigger() string {
	return "delete_personal_channel_welcome_message"
}

func (c *DeletePersonalChanelWelcomeMessage) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.IsSysadmin(args.UserId) || p.CanManageChannel(args.UserId, args.ChannelId)
}

func (c *DeletePersonalChanelWelcomeMessage) Help() string {
	return "`/welcomebot delete_personal_channel_welcome_message` - " + deletePersonalChanelWelcomeMessageHelp
}

func (c *DeletePersonalChanelWelcomeMessage) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	cmd := command.DeletePersonalChanelWelcomeMessage{
		CommandMessenger:   p.Container().NewCommandMessenger(args),
		ChannelWelcomeRepo: p.Container().ChannelWelcomeRepo(),
	}

	cmd.Call(args.ChannelId)
}

func (c *DeletePersonalChanelWelcomeMessage) Validate(parameters []string) error {
	if len(parameters) > 0 {
		return errors.New("`delete_personal_channel_welcome_message` command does not accept any extra parameters")
	}

	return nil
}

func (c *DeletePersonalChanelWelcomeMessage) AutocompleteData() *model.AutocompleteData {
	return model.NewAutocompleteData("delete_personal_channel_welcome_message", "", deletePersonalChanelWelcomeMessageHelp)
}
