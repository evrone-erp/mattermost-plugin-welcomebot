package main

import (
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/core"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler/command"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler/hook"
	"github.com/mattermost/mattermost/server/public/plugin"
)

func main() {
	corePlugin := core.NewPlugin(manifest)
	container := Container{plugin: &corePlugin}
	corePlugin.RegisterDependencyContainer(&container)

	corePlugin.RegisterCommand(&command.GetPersonalChanelWelcome{})
	corePlugin.RegisterCommand(&command.SetPersonalChanelWelcome{})
	corePlugin.RegisterCommand(&command.DeletePersonalChanelWelcome{})
	corePlugin.RegisterCommand(&command.GetPublishedChanelWelcome{})
	corePlugin.RegisterCommand(&command.SetPublishedChanelWelcome{})
	corePlugin.RegisterCommand(&command.DeletePublishedChanelWelcome{})

	corePlugin.RegisterCommand(&command.GetTeamWelcome{})
	corePlugin.RegisterCommand(&command.SetTeamWelcomeMessage{})
	corePlugin.RegisterCommand(&command.RemoveTeamWelcomeMessage{})

	corePlugin.RegisterCommand(&command.ListChannelWelcomes{})
	corePlugin.RegisterCommand(&command.AddTeamDefaultChannels{})
	corePlugin.RegisterCommand(&command.RemoveTeamDefaultChannels{})

	corePlugin.RegisterUserHasJoinedChannelHook(&hook.PersonalWelcomeNotifier{})
	corePlugin.RegisterUserHasJoinedChannelHook(&hook.PublishedWelcomeNotifier{})
	corePlugin.RegisterUserHasJoinedTeamHook(&hook.TeamWelcomeNotifier{})

	plugin.ClientMain(&corePlugin)
}
