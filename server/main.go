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

	corePlugin.RegisterCommand(&command.GetPersonalChanelWelcomeMessage{})
	corePlugin.RegisterCommand(&command.SetPersonalChanelWelcomeMessage{})
	corePlugin.RegisterCommand(&command.DeletePersonalChanelWelcomeMessage{})
	corePlugin.RegisterCommand(&command.GetPublishedChanelWelcomeMessage{})
	corePlugin.RegisterCommand(&command.SetPublishedChanelWelcomeMessage{})
	corePlugin.RegisterCommand(&command.DeletePublishedChanelWelcomeMessage{})

	corePlugin.RegisterCommand(&command.GetTeamWelcomeSettings{})
	corePlugin.RegisterCommand(&command.SetTeamWelcomeMessage{})
	corePlugin.RegisterCommand(&command.DeleteTeamWelcomeMessage{})

	corePlugin.RegisterCommand(&command.ListChannelWelcomes{})
	corePlugin.RegisterCommand(&command.AddTeamDefaultChannels{})
	corePlugin.RegisterCommand(&command.RemoveTeamDefaultChannels{})

	corePlugin.RegisterUserHasJoinedChannelHook(&hook.PersonalWelcomeNotifier{})
	corePlugin.RegisterUserHasJoinedChannelHook(&hook.PublishedWelcomeNotifier{})
	corePlugin.RegisterUserHasJoinedTeamHook(&hook.TeamWelcomeNotifier{})

	plugin.ClientMain(&corePlugin)
}
