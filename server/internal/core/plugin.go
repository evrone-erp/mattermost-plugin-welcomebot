package core

import (
	"fmt"
	"sort"
	"strings"
	"sync/atomic"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
	"github.com/mattermost/mattermost/server/public/pluginapi"
	"github.com/pkg/errors"
)

const (
	commandTriggerHelp = "help"

	botUsername    = "welcomebot"
	botDisplayName = "Welcomebot"
	botDescription = "A bot account created by the Welcomebot plugin."
)

type CommandInterface interface {
	Trigger() string
	Help() string
	Execute(handler.BotAPIProvider, *model.CommandArgs)
	Validate([]string) error
	AutocompleteData() *model.AutocompleteData
	IsPermitted(p usecase.Policy, args *model.CommandArgs) bool
}

type ChannelJoinHook interface {
	Execute(handler.BotAPIProvider, *model.ChannelMember)
}

type TeamJoinHook interface {
	Execute(handler.BotAPIProvider, *model.TeamMember, *model.User)
}

// Plugin represents the welcome bot plugin
type Plugin struct {
	plugin.MattermostPlugin

	BotUserID string
	client    *pluginapi.Client
	// BotUserID of the created bot account.
	Manifest                  *model.Manifest
	welcomeMessages           atomic.Value
	commands                  map[string]CommandInterface
	container                 handler.DependencyContainer
	userHasJoinedChannelHooks []ChannelJoinHook
	userHasJoinedTeamHooks    []TeamJoinHook
}

func NewPlugin(manifest *model.Manifest) Plugin {
	return Plugin{
		Manifest:                  manifest,
		commands:                  make(map[string]CommandInterface),
		userHasJoinedChannelHooks: make([]ChannelJoinHook, 0),
		userHasJoinedTeamHooks:    make([]TeamJoinHook, 0),
	}
}

func (p *Plugin) RegisterDependencyContainer(c handler.DependencyContainer) {
	p.container = c
}

func (p *Plugin) Container() handler.DependencyContainer {
	return p.container
}

func (p *Plugin) RegisterCommand(c CommandInterface) {
	_, conflict := p.commands[c.Trigger()]

	if conflict {
		msg := fmt.Sprintf("Command trigger %s already registered", c.Trigger())
		panic(msg)
	}

	p.commands[c.Trigger()] = c
}

func (p *Plugin) RegisterUserHasJoinedChannelHook(h ChannelJoinHook) {
	p.userHasJoinedChannelHooks = append(p.userHasJoinedChannelHooks, h)
}

func (p *Plugin) RegisterUserHasJoinedTeamHook(h TeamJoinHook) {
	p.userHasJoinedTeamHooks = append(p.userHasJoinedTeamHooks, h)
}

func (p *Plugin) APIHandle() plugin.API {
	return p.API
}

func (p *Plugin) BotUserIDHandle() string {
	return p.BotUserID
}

// OnActivate ensure the bot account exists
func (p *Plugin) OnActivate() error {
	p.client = pluginapi.NewClient(p.API, p.Driver)

	bot := &model.Bot{
		Username:    botUsername,
		DisplayName: botDisplayName,
		Description: botDescription,
	}
	BotUserID, appErr := p.client.Bot.EnsureBot(bot)
	if appErr != nil {
		return errors.Wrap(appErr, "failed to ensure bot user")
	}
	p.BotUserID = BotUserID

	err := p.API.RegisterCommand(p.GetCommand())
	if err != nil {
		return errors.Wrap(err, "failed to register command")
	}

	return nil
}

// UserHasJoinedTeam is invoked after the membership has been committed to the database. If
// actor is not nil, the user was added to the team by the actor.
func (p *Plugin) UserHasJoinedTeam(_ *plugin.Context, teamMember *model.TeamMember, actor *model.User) {
	for _, hook := range p.userHasJoinedTeamHooks {
		hook.Execute(p, teamMember, actor)
	}
}

// UserHasJoinedChannel is invoked after the membership has been committed to
// the database. If actor is not nil, the user was invited to the channel by
// the actor.
func (p *Plugin) UserHasJoinedChannel(_ *plugin.Context, channelMember *model.ChannelMember, _ *model.User) {
	for _, hook := range p.userHasJoinedChannelHooks {
		hook.Execute(p, channelMember)
	}
}

func (p *Plugin) GetCommand() *model.Command {
	return &model.Command{
		Trigger:          "welcomebot",
		DisplayName:      "welcomebot",
		Description:      "Welcome Bot helps add new team members to channels.",
		AutoComplete:     true,
		AutoCompleteDesc: p.AutoCompleteDesc(),
		AutoCompleteHint: "[command]",
		AutocompleteData: p.GetAutocompleteData(),
	}
}

func (p *Plugin) AutoCompleteDesc() string {
	triggers := make([]string, len(p.commands)+1)

	for key := range p.commands {
		triggers = append(triggers, key)
	}

	triggers = append(triggers, commandTriggerHelp)
	result := fmt.Sprintf("Available commands: %s", strings.Join(triggers, ", "))

	return result
}

func (p *Plugin) ExecuteCommand(_ *plugin.Context, args *model.CommandArgs) (*model.CommandResponse, *model.AppError) {
	split := strings.Fields(args.Command)
	command := split[0]
	parameters := []string{}
	action := ""
	if len(split) > 1 {
		action = split[1]
	}
	if len(split) > 2 {
		parameters = split[2:]
	}

	if command != "/welcomebot" {
		return &model.CommandResponse{}, nil
	}

	if action == "" || action == commandTriggerHelp {
		p.printHelp(args)
		return &model.CommandResponse{}, nil
	}

	commandObj, ok := p.commands[action]

	if !ok {
		msg := fmt.Sprintf("Unknown action %v", action)
		p.Container().NewCommandMessenger(args).PostCommandResponse(msg)
		return &model.CommandResponse{}, nil
	}

	isPermitted := commandObj.IsPermitted(p.Container().Policy(), args)

	if !isPermitted {
		p.Container().NewCommandMessenger(args).PostCommandResponse("You have no permissions for this command")
		return &model.CommandResponse{}, nil
	}

	err := commandObj.Validate(parameters)
	if err != nil {
		p.Container().NewCommandMessenger(args).PostCommandResponse(err.Error())
		return &model.CommandResponse{}, nil
	}

	commandObj.Execute(p, args)
	return &model.CommandResponse{}, nil
}

func (p *Plugin) printHelp(args *model.CommandArgs) {
	commandsHelp := make([]string, 0, len(p.commands))
	for _, command := range p.commands {
		commandsHelp = append(commandsHelp, command.Help())
	}

	sort.Strings(commandsHelp)

	text := "###### Mattermost welcomebot Plugin - Slash Command Help\n" + strings.Join(commandsHelp, "\n")
	p.Container().NewCommandMessenger(args).PostCommandResponse(text)
}

func (p *Plugin) GetAutocompleteData() *model.AutocompleteData {
	triggers := make([]string, len(p.commands))

	for key := range p.commands {
		triggers = append(triggers, key)
	}

	description := fmt.Sprintf("Available commands: %s", strings.Join(triggers, ", "))
	welcomebot := model.NewAutocompleteData("welcomebot", "[command]", description)

	for _, command := range p.commands {
		welcomebot.AddCommand(command.AutocompleteData())
	}

	return welcomebot
}
