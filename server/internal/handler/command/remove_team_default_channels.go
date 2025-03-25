package command

import (
	"errors"
	"fmt"
	"strings"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/command"
	"github.com/mattermost/mattermost/server/public/model"
)

type RemoveTeamDefaultChannels struct{}

const removeTeamDefaultChannelsHelp = "Remove channels from the auto-join list when entering the **current team**"

type RemoveTeamDefaultChannelsInput struct {
	channelNames []string
}

func (c *RemoveTeamDefaultChannels) Trigger() string {
	return "remove_team_default_channels"
}

func (c *RemoveTeamDefaultChannels) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.IsSysadmin(args.UserId)
}

func (c *RemoveTeamDefaultChannels) Help() string {
	return "`/welcomebot remove_team_default_channels <[~channel]>` - " + removeTeamDefaultChannelsHelp
}

func (c *RemoveTeamDefaultChannels) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	msgr := p.Container().NewCommandMessenger(args)
	input, err := c.parse(args)

	if err != nil {
		msg := fmt.Sprintf("Error while parsing command %s", err)
		msgr.PostCommandResponse(msg)
		return
	}

	cmd := command.RemoveTeamDefaultChannels{
		CommandMessenger: msgr,
		ChannelRepo:      p.Container().ChannelRepo(),
		TeamWelcomeRepo:  p.Container().TeamWelcomeRepo(),
	}

	cmd.Call(args.TeamId, input.channelNames)
}

func (c *RemoveTeamDefaultChannels) Validate(parameters []string) error {
	if len(parameters) < 1 {
		return errors.New("`remove_team_default_channels` command requires the channels list")
	}

	return nil
}

func (c *RemoveTeamDefaultChannels) AutocompleteData() *model.AutocompleteData {
	data := model.NewAutocompleteData("remove_team_default_channels", "", removeTeamDefaultChannelsHelp)
	data.AddTextArgument("Message", "[~space_separated_channels]", "")

	return data
}

func (c *RemoveTeamDefaultChannels) parse(args *model.CommandArgs) (*RemoveTeamDefaultChannelsInput, error) {
	fields := strings.Fields(args.Command)

	if len(fields) < 3 {
		msg := fmt.Sprintf("Unable to parse command %s", args.Command)
		return nil, errors.New(msg)
	}

	result := new(RemoveTeamDefaultChannelsInput)

	result.channelNames = fields[2:]

	for i, name := range result.channelNames {
		result.channelNames[i] = strings.TrimLeft(name, "~")
	}

	return result, nil
}
