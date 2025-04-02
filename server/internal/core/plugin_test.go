package core

import (
	"regexp"
	"testing"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler/command"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"

	"github.com/stretchr/testify/mock"
)

func TestExecuteCommand(t *testing.T) {
	type Setup struct {
		Container          *handler.MockDependencyContainer
		CommandMessenger   *usecase.MockCommandMessenger
		ChannelWelcomeRepo *usecase.MockChannelWelcomeRepo
		Policy             *usecase.MockPolicy
		UserRepo           *usecase.MockUserRepo
		ChannelRepo        *usecase.MockChannelRepo
		Subject            *Plugin
		ExecuteCommand     func(string)
	}

	buildCommandArgs := func(cmd string) *model.CommandArgs {
		return &model.CommandArgs{
			UserId:    "user-id",
			ChannelId: "channel-id",
			TeamId:    "team-id",
			Command:   cmd,
		}
	}

	setup := func() *Setup {
		container := new(handler.MockDependencyContainer)
		commandMessenger := new(usecase.MockCommandMessenger)
		policy := new(usecase.MockPolicy)
		channelWelcomeRepo := new(usecase.MockChannelWelcomeRepo)
		userRepo := new(usecase.MockUserRepo)
		channelRepo := new(usecase.MockChannelRepo)
		teamWelcomeRepo := new(usecase.MockTeamWelcomeRepo)

		container.On("NewCommandMessenger", mock.Anything).Return(commandMessenger)
		container.On("Policy").Return(policy)
		container.On("ChannelWelcomeRepo").Return(channelWelcomeRepo)
		container.On("UserRepo").Return(userRepo)
		container.On("ChannelRepo").Return(channelRepo)
		container.On("TeamWelcomeRepo").Return(teamWelcomeRepo)

		commandMessenger.On("PostCommandResponse", mock.Anything).Return(commandMessenger)
		policy.On("IsSysadmin", mock.Anything).Return(true)
		policy.On("CanManageTeam", mock.Anything, mock.Anything).Return(true)

		channel1 := &model.Channel{
			Id:   "channel-id",
			Name: "channel-name",
			Type: model.ChannelTypeOpen,
		}

		channelRepo.On("Get", "channel-id").Return(channel1, nil)
		channelRepo.On("GetByName", "team-id", "channel-name").Return(channel1, nil)

		userRepo.On("GetByID", "user-id").Return(&model.User{
			Username:  "johny",
			FirstName: "John",
			LastName:  "Doe",
		}, nil)

		personalWelcome := &pmodel.ChannelWelcome{
			ID:        "channel-id",
			Message:   "Personal Hello for {{.UserHandleName}}!",
			Type:      pmodel.WelcomeMessageTypePersonal,
			ChannelID: "channel-id",
		}

		publishedWelcome := &pmodel.ChannelWelcome{
			ID:        "channel-id",
			Message:   "Published Hello for {{.UserHandleName}}!",
			Type:      pmodel.WelcomeMessageTypePublished,
			ChannelID: "channel-id",
		}

		teamWelcomeMsg := "Team welcome for {{.UserHandleName}}!"
		teamWelcome := &pmodel.TeamWelcome{
			Message: teamWelcomeMsg,
		}

		channelWelcomeRepo.On("GetPersonalChanelWelcome", "channel-id").Return(personalWelcome, nil)
		channelWelcomeRepo.On("GetPublishedChanelWelcome", "channel-id").Return(publishedWelcome, nil)
		channelWelcomeRepo.On("SetPublishedChanelWelcome", "channel-id", "new pub").Return(nil)
		channelWelcomeRepo.On("SetPersonalChanelWelcome", "channel-id", "new per").Return(nil)
		channelWelcomeRepo.On("DeletePublishedChanelWelcome", "channel-id").Return(nil)
		channelWelcomeRepo.On("DeletePersonalChanelWelcome", "channel-id").Return(nil)
		channelWelcomeRepo.On("ListChannelsWithWelcome").Return([]string{"channel-id"}, []string{}, nil)

		teamWelcomeRepo.On("SetTeamWelcome", "team-id", &pmodel.TeamWelcome{Message: "new tw"}).Return(nil)
		teamWelcomeRepo.On("SetTeamWelcome", "team-id", &pmodel.TeamWelcome{Message: ""}).Return(nil)
		teamWelcomeRepo.On("SetTeamWelcome", "team-id", &pmodel.TeamWelcome{Message: teamWelcomeMsg, ChannelIDs: []string{"channel-id"}}).Return(nil)
		teamWelcomeRepo.On("SetTeamWelcome", "team-id", &pmodel.TeamWelcome{Message: teamWelcomeMsg, ChannelIDs: []string{}}).Return(nil)
		teamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(teamWelcome, nil)

		p := NewPlugin(&model.Manifest{})
		p.RegisterDependencyContainer(container)
		p.RegisterCommand(&command.GetPersonalChanelWelcomeMessage{})
		p.RegisterCommand(&command.GetPublishedChanelWelcomeMessage{})
		p.RegisterCommand(&command.SetPublishedChanelWelcomeMessage{})
		p.RegisterCommand(&command.SetPersonalChanelWelcomeMessage{})
		p.RegisterCommand(&command.DeletePublishedChanelWelcomeMessage{})
		p.RegisterCommand(&command.DeletePersonalChanelWelcomeMessage{})
		p.RegisterCommand(&command.ListChannelWelcomes{})
		p.RegisterCommand(&command.SetTeamWelcomeMessage{})
		p.RegisterCommand(&command.GetTeamWelcomeSettings{})
		p.RegisterCommand(&command.DeleteTeamWelcomeMessage{})
		p.RegisterCommand(&command.AddTeamDefaultChannels{})
		p.RegisterCommand(&command.RemoveTeamDefaultChannels{})

		executeCommand := func(cmd string) {
			_, _ = p.ExecuteCommand(&plugin.Context{}, buildCommandArgs(cmd))
		}

		return &Setup{
			Container:          container,
			CommandMessenger:   commandMessenger,
			Policy:             policy,
			ChannelWelcomeRepo: channelWelcomeRepo,
			UserRepo:           userRepo,
			Subject:            &p,
			ExecuteCommand:     executeCommand,
		}
	}

	t.Run("help message", func(t *testing.T) {
		s := setup()
		s.ExecuteCommand("/welcomebot")

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", mock.MatchedBy(func(arg string) bool {
			return regexp.MustCompile(`welcomebot add_team_default_channels`).MatchString(arg) &&
				regexp.MustCompile(`Welcome messages support a simple templating`).MatchString(arg)
		}))
	})

	t.Run("invalid command", func(t *testing.T) {
		s := setup()
		s.ExecuteCommand("/welcomebot dunno")

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Unknown action dunno")
	})

	t.Run("happy get_personal_channel_welcome_message", func(t *testing.T) {
		s := setup()
		s.ExecuteCommand("/welcomebot get_personal_channel_welcome_message")

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome message is:\nPersonal Hello for @johny!")
	})

	t.Run("happy get_publilshed_channel_welcome", func(t *testing.T) {
		s := setup()
		s.ExecuteCommand("/welcomebot get_published_channel_welcome_message")

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome message is:\nPublished Hello for @johny!")
	})

	t.Run("happy set_publilshed_channel_welcome", func(t *testing.T) {
		s := setup()
		s.ExecuteCommand("/welcomebot set_published_channel_welcome_message new pub")

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "stored the welcome message:\nnew pub")
	})

	t.Run("happy set_personal_channel_welcome_message", func(t *testing.T) {
		s := setup()
		s.ExecuteCommand("/welcomebot set_personal_channel_welcome_message new per")

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "stored the welcome message:\nnew per")
	})

	t.Run("happy delete_publilshed_channel_welcome_message", func(t *testing.T) {
		s := setup()
		s.ExecuteCommand("/welcomebot delete_published_channel_welcome_message")

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "welcome message has been deleted")
	})

	t.Run("happy delete_personal_channel_welcome_message", func(t *testing.T) {
		s := setup()
		s.ExecuteCommand("/welcomebot delete_personal_channel_welcome_message")

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "welcome message has been deleted")
	})

	t.Run("happy list_channel_welcomes", func(t *testing.T) {
		s := setup()
		s.ExecuteCommand("/welcomebot list_channel_welcomes")

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Channels with personal welcome message:\n~channel-name\n")
	})

	t.Run("happy set_team_welcome_message", func(t *testing.T) {
		s := setup()
		s.ExecuteCommand("/welcomebot set_team_welcome_message new tw")

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome message was updated")
	})

	t.Run("happy delete_team_welcome_message", func(t *testing.T) {
		s := setup()
		s.ExecuteCommand("/welcomebot delete_team_welcome_message")

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome message was deleted")
	})

	t.Run("happy get_team_welcome_settings", func(t *testing.T) {
		s := setup()
		s.ExecuteCommand("/welcomebot get_team_welcome_settings")

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Current welcome is:\nTeam welcome for @johny!\n\n\nNo default channels\n")
	})

	t.Run("happy add_team_default_channels", func(t *testing.T) {
		s := setup()
		s.ExecuteCommand("/welcomebot add_team_default_channels ~channel-name")

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome channels were updated")
	})

	t.Run("happy remove_team_default_channels", func(t *testing.T) {
		s := setup()
		s.ExecuteCommand("/welcomebot remove_team_default_channels ~channel-name")

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome channels were updated")
	})
}
