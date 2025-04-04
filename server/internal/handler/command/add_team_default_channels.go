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

type AddTeamDefaultChannels struct{}

const addTeamDefaultChannelsHelp = "Add channels to automatically join when entering the **current team**"

type AddTeamDefaultChannelsInput struct {
	channelNames []string
}

func (c *AddTeamDefaultChannels) Trigger() string {
	return "add_team_default_channels"
}

func (c *AddTeamDefaultChannels) Help() string {
	return "`/welcomebot add_team_default_channels <[~channel]>` - " + addTeamDefaultChannelsHelp
}

func (c *AddTeamDefaultChannels) Execute(p handler.BotAPIProvider, args *model.CommandArgs) {
	msgr := p.Container().NewCommandMessenger(args)
	input, err := c.parse(args)

	if err != nil {
		msg := fmt.Sprintf("Error while parsing command %s", err)
		msgr.PostCommandResponse(msg)
		return
	}

	cmd := command.AddTeamDefaultChannels{
		CommandMessenger: msgr,
		TeamWelcomeRepo:  p.Container().TeamWelcomeRepo(),
		ChannelRepo:      p.Container().ChannelRepo(),
	}

	cmd.Call(args.TeamId, input.channelNames)
}

func (c *AddTeamDefaultChannels) IsPermitted(p usecase.Policy, args *model.CommandArgs) bool {
	return p.CanManageTeam(args.UserId, args.TeamId)
}

func (c *AddTeamDefaultChannels) Validate(parameters []string) error {
	if len(parameters) < 1 {
		return errors.New("`add_team_default_channels` command requires the channels list")
	}

	return nil
}

func (c *AddTeamDefaultChannels) AutocompleteData() *model.AutocompleteData {
	data := model.NewAutocompleteData("add_team_default_channels", "[~channel_name]", addTeamDefaultChannelsHelp)
	data.AddTextArgument("Message", "[~space_separated_channels]", "")

	return data
}

func (c *AddTeamDefaultChannels) parse(args *model.CommandArgs) (*AddTeamDefaultChannelsInput, error) {
	fields := strings.Fields(args.Command)

	if len(fields) < 3 {
		msg := fmt.Sprintf("Unable to parse command %s", args.Command)
		return nil, errors.New(msg)
	}

	result := new(AddTeamDefaultChannelsInput)

	result.channelNames = fields[2:]

	for i, name := range result.channelNames {
		result.channelNames[i] = strings.TrimLeft(name, "~")
	}

	return result, nil
}
