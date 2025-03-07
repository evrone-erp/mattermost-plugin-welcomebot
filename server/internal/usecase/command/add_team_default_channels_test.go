package command

import (
	"testing"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	mmodel "github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/mock"
)

func TestAddTeamDefaultChannels(t *testing.T) {
	type Setup struct {
		CommandMessenger *usecase.MockCommandMessenger
		ChannelRepo      *usecase.MockChannelRepo
		TeamWelcomeRepo  *usecase.MockTeamWelcomeRepo
		Subject          *AddTeamDefaultChannels
	}

	setup := func() *Setup {
		commandMessenger := new(usecase.MockCommandMessenger)
		channelRepo := new(usecase.MockChannelRepo)
		teamWelcomeRepo := new(usecase.MockTeamWelcomeRepo)

		commandMessenger.On("PostCommandResponse", mock.Anything).Return()

		subject := &AddTeamDefaultChannels{
			CommandMessenger: commandMessenger,
			TeamWelcomeRepo:  teamWelcomeRepo,
			ChannelRepo:      channelRepo,
		}

		return &Setup{
			CommandMessenger: commandMessenger,
			TeamWelcomeRepo:  teamWelcomeRepo,
			ChannelRepo:      channelRepo,
			Subject:          subject,
		}
	}

	t.Run("happy path without previous welcome", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(nil, nil)

		updatedWelcome := &pmodel.TeamWelcome{
			ChannelIDs: []string{"ci1", "ci2"},
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(nil)
		s.ChannelRepo.On("GetByName", "team-id", "cn1").Return(&mmodel.Channel{Id: "ci1"}, nil)
		s.ChannelRepo.On("GetByName", "team-id", "cn2").Return(&mmodel.Channel{Id: "ci2"}, nil)

		s.Subject.Call("team-id", []string{"cn1", "cn2"})

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome channels were updated")
	})

	t.Run("with duplicated input", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(nil, nil)

		updatedWelcome := &pmodel.TeamWelcome{
			ChannelIDs: []string{"ci1", "ci2"},
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(nil)
		s.ChannelRepo.On("GetByName", "team-id", "cn1").Return(&mmodel.Channel{Id: "ci1"}, nil)
		s.ChannelRepo.On("GetByName", "team-id", "cn2").Return(&mmodel.Channel{Id: "ci2"}, nil)

		s.Subject.Call("team-id", []string{"cn1", "cn2", "cn2", "cn1"})

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome channels were updated")
	})

	t.Run("happy path without previous welcome and partially broken args", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(nil, nil)

		updatedWelcome := &pmodel.TeamWelcome{
			ChannelIDs: []string{"ci1"},
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(nil)
		s.ChannelRepo.On("GetByName", "team-id", "cn1").Return(&mmodel.Channel{Id: "ci1"}, nil)
		s.ChannelRepo.On("GetByName", "team-id", "garbage").Return(nil, &mmodel.AppError{Message: "not found"})

		s.Subject.Call("team-id", []string{"cn1", "garbage"})

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 2)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Skipping unknown channel garbage")
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome channels were updated")
	})

	t.Run("when there is some previous garbage", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(
			&pmodel.TeamWelcome{ChannelIDs: []string{"pr1", "garbage"}},
			nil,
		)

		s.ChannelRepo.On("GetByName", "team-id", "cn1").Return(&mmodel.Channel{Id: "ci1"}, nil)
		s.ChannelRepo.On("Get", "pr1").Return(&mmodel.Channel{Id: "pr1"}, nil)
		s.ChannelRepo.On("Get", "garbage").Return(nil, &mmodel.AppError{Message: "not found"})
		s.ChannelRepo.On("GetByName", "team-id", "garbage").Return(nil, &mmodel.AppError{Message: "not found"})

		updatedWelcome := &pmodel.TeamWelcome{
			ChannelIDs: []string{"ci1", "pr1"},
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(nil)

		s.Subject.Call("team-id", []string{"cn1"})

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 2)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Dropping previously saved channel#id garbage")
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome channels were updated")
	})

	t.Run("when tehre was a valid previous channel", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(&pmodel.TeamWelcome{ChannelIDs: []string{"pr1"}}, nil)
		s.ChannelRepo.On("Get", "pr1").Return(&mmodel.Channel{Id: "pr1"}, nil)

		updatedWelcome := &pmodel.TeamWelcome{
			ChannelIDs: []string{"ci1", "pr1"},
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(nil)
		s.ChannelRepo.On("GetByName", "team-id", "cn1").Return(&mmodel.Channel{Id: "ci1"}, nil)

		s.Subject.Call("team-id", []string{"cn1"})

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome channels were updated")
	})

	t.Run("with a saving error", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(nil, nil)

		updatedWelcome := &pmodel.TeamWelcome{
			ChannelIDs: []string{"ci1"},
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(&mmodel.AppError{Message: "fail"})
		s.ChannelRepo.On("GetByName", "team-id", "cn1").Return(&mmodel.Channel{Id: "ci1"}, nil)

		s.Subject.Call("team-id", []string{"cn1"})

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Error while saving team team-id: fail")
	})
}
