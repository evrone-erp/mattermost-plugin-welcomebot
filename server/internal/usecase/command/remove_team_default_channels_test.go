package command

import (
	"testing"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	mmodel "github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/mock"
)

func TestRemoveTeamDefaultChannels(t *testing.T) {
	type Setup struct {
		CommandMessenger *usecase.MockCommandMessenger
		TeamWelcomeRepo  *usecase.MockTeamWelcomeRepo
		ChannelRepo      *usecase.MockChannelRepo
		Subject          *RemoveTeamDefaultChannels
	}

	setup := func() *Setup {
		messenger := new(usecase.MockCommandMessenger)
		teamWelcomeRepo := new(usecase.MockTeamWelcomeRepo)
		channelRepo := new(usecase.MockChannelRepo)

		messenger.On("PostCommandResponse", mock.Anything).Return()

		subject := &RemoveTeamDefaultChannels{
			CommandMessenger: messenger,
			TeamWelcomeRepo:  teamWelcomeRepo,
			ChannelRepo:      channelRepo,
		}

		return &Setup{
			CommandMessenger: messenger,
			TeamWelcomeRepo:  teamWelcomeRepo,
			ChannelRepo:      channelRepo,
			Subject:          subject,
		}
	}

	t.Run("happy path", func(t *testing.T) {
		s := setup()

		originalWelcome := &pmodel.TeamWelcome{
			Message:    "foo",
			ChannelIDs: []string{"ci1", "ci2", "ci3"},
		}

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(originalWelcome, nil)

		updatedWelcome := &pmodel.TeamWelcome{
			Message:    "foo",
			ChannelIDs: []string{"ci2"},
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(nil)
		s.ChannelRepo.On("GetByName", "team-id", "cn1").Return(&mmodel.Channel{Id: "ci1"}, nil)
		s.ChannelRepo.On("GetByName", "team-id", "cn3").Return(&mmodel.Channel{Id: "ci3"}, nil)
		s.ChannelRepo.On("Get", "ci1").Return(&mmodel.Channel{Id: "ci1"}, nil)
		s.ChannelRepo.On("Get", "ci2").Return(&mmodel.Channel{Id: "ci2"}, nil)
		s.ChannelRepo.On("Get", "ci3").Return(&mmodel.Channel{Id: "ci3"}, nil)

		s.Subject.Call("team-id", []string{"cn1", "cn3"})

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome channels were updated")
	})

	t.Run("broken previous channels", func(t *testing.T) {
		s := setup()

		originalWelcome := &pmodel.TeamWelcome{
			Message:    "foo",
			ChannelIDs: []string{"ci1", "ci2", "ci3"},
		}

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(originalWelcome, nil)

		updatedWelcome := &pmodel.TeamWelcome{
			Message:    "foo",
			ChannelIDs: []string{"ci2"},
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(nil)
		s.ChannelRepo.On("GetByName", "team-id", "cn3").Return(&mmodel.Channel{Id: "ci3"}, nil)
		s.ChannelRepo.On("Get", "ci1").Return(nil, &mmodel.AppError{Message: "not found"})
		s.ChannelRepo.On("Get", "ci2").Return(&mmodel.Channel{Id: "ci2"}, nil)
		s.ChannelRepo.On("Get", "ci3").Return(&mmodel.Channel{Id: "ci3"}, nil)

		s.Subject.Call("team-id", []string{"cn3"})

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 2)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Dropping previously saved broken channel#id ci1")
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome channels were updated")
	})

	t.Run("invalid channels", func(t *testing.T) {
		s := setup()

		originalWelcome := &pmodel.TeamWelcome{
			Message:    "foo",
			ChannelIDs: []string{"ci1", "ci2"},
		}

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(originalWelcome, nil)

		updatedWelcome := &pmodel.TeamWelcome{
			Message:    "foo",
			ChannelIDs: []string{"ci1"},
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(nil)
		s.ChannelRepo.On("GetByName", "team-id", "cn2").Return(&mmodel.Channel{Id: "ci2"}, nil)
		s.ChannelRepo.On("GetByName", "team-id", "garbage").Return(nil, &mmodel.AppError{Message: "Not found"})
		s.ChannelRepo.On("Get", "ci1").Return(&mmodel.Channel{Id: "ci1"}, nil)
		s.ChannelRepo.On("Get", "ci2").Return(&mmodel.Channel{Id: "ci2"}, nil)

		s.Subject.Call("team-id", []string{"cn2", "garbage"})

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 2)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Error while checking channel garbage:Not found. Ignoring")
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome channels were updated")
	})

	t.Run("valid unconfigured channels", func(t *testing.T) {
		s := setup()

		originalWelcome := &pmodel.TeamWelcome{
			Message:    "foo",
			ChannelIDs: []string{"ci1", "ci2"},
		}

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(originalWelcome, nil)

		updatedWelcome := &pmodel.TeamWelcome{
			Message:    "foo",
			ChannelIDs: []string{"ci1"},
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(nil)
		s.ChannelRepo.On("GetByName", "team-id", "cn2").Return(&mmodel.Channel{Id: "ci2"}, nil)
		s.ChannelRepo.On("GetByName", "team-id", "cn3").Return(&mmodel.Channel{Id: "ci3"}, nil)
		s.ChannelRepo.On("Get", "ci1").Return(&mmodel.Channel{Id: "ci1"}, nil)
		s.ChannelRepo.On("Get", "ci2").Return(&mmodel.Channel{Id: "ci2"}, nil)
		s.ChannelRepo.On("Get", "ci3").Return(&mmodel.Channel{Id: "ci3"}, nil)

		s.Subject.Call("team-id", []string{"cn2", "cn3"})

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 2)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Ignoring not configured channel cn3")
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome channels were updated")
	})

	t.Run("error while saving", func(t *testing.T) {
		s := setup()
		originalWelcome := &pmodel.TeamWelcome{
			Message:    "foo",
			ChannelIDs: []string{"ci1"},
		}

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(originalWelcome, nil)

		updatedWelcome := &pmodel.TeamWelcome{
			Message:    "foo",
			ChannelIDs: []string{},
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(&mmodel.AppError{Message: "fail"})
		s.ChannelRepo.On("GetByName", "team-id", "cn1").Return(&mmodel.Channel{Id: "ci1"}, nil)
		s.ChannelRepo.On("Get", "ci1").Return(&mmodel.Channel{Id: "ci1"}, nil)

		s.Subject.Call("team-id", []string{"cn1"})

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Error while saving team team-id: fail")
	})

	t.Run("no welcome at all", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(nil, nil)

		s.Subject.Call("team-id", []string{"cn1"})

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "There is nothing to update")
	})
}
