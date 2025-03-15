package command

import (
	"errors"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

const deletePublishedChanelWelcomeMessageHelp = "Delete the published welcome message for the **current channel** (if any)"

type DeletePublishedChanelWelcomeMessage struct{}

func (c *DeletePublishedChanelWelcomeMessage) Trigger() string {
	return "delete_published_channel_welcome_message"
}

func (c *DeletePublishedChanelWelcomeMessage) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.IsSysadmin(args.UserId) || p.CanManageChannel(args.UserId, args.ChannelId)
}

func (c *DeletePublishedChanelWelcomeMessage) Help() string {
	return "`/welcomebot delete_published_channel_welcome_message` - " + deletePublishedChanelWelcomeMessageHelp
}

func (c *DeletePublishedChanelWelcomeMessage) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	cmd := command.DeletePublishedChanelWelcomeMessage{
		CommandMessenger:   p.Container().NewCommandMessenger(args),
		ChannelWelcomeRepo: p.Container().ChannelWelcomeRepo(),
	}

	cmd.Call(args.ChannelId)
}

func (c *DeletePublishedChanelWelcomeMessage) Validate(parameters []string) error {
	if len(parameters) > 0 {
		return errors.New("`delete_published_channel_welcome_message` command does not accept any extra parameters")
	}

	return nil
}

func (c *DeletePublishedChanelWelcomeMessage) AutocompleteData() *model.AutocompleteData {
	return model.NewAutocompleteData("delete_published_channel_welcome_message", "", deletePublishedChanelWelcomeMessageHelp)
}
