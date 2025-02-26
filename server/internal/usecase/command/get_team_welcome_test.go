package command

import (
	"testing"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	mmodel "github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/mock"
)

func TestGetTeamWelcome(t *testing.T) {
	type Setup struct {
		CommandMessenger        *usecase.MockCommandMessenger
		TeamWelcomeRepo         *usecase.MockTeamWelcomeRepo
		ChannelRepo             *usecase.MockChannelRepo
		WelcomeMessagePresenter *usecase.MockWelcomeMessagePresenter
		Subject                 *GetTeamWelcome
	}

	setup := func() *Setup {
		commandMessenger := new(usecase.MockCommandMessenger)
		teamWelcomeRepo := new(usecase.MockTeamWelcomeRepo)
		channelRepo := new(usecase.MockChannelRepo)
		welcomeMessagePresenter := new(usecase.MockWelcomeMessagePresenter)

		commandMessenger.On("PostCommandResponse", mock.Anything).Return()

		subject := &GetTeamWelcome{
			CommandMessenger:        commandMessenger,
			TeamWelcomeRepo:         teamWelcomeRepo,
			ChannelRepo:             channelRepo,
			WelcomeMessagePresenter: welcomeMessagePresenter,
		}

		return &Setup{
			CommandMessenger:        commandMessenger,
			TeamWelcomeRepo:         teamWelcomeRepo,
			ChannelRepo:             channelRepo,
			WelcomeMessagePresenter: welcomeMessagePresenter,
			Subject:                 subject,
		}
	}

	t.Run("with a valid message", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(&pmodel.TeamWelcome{Message: "hello"}, nil)
		s.WelcomeMessagePresenter.On("Render", "hello", "user-id").Return("hellop", nil)

		s.Subject.Call("user-id", "team-id")

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Current welcome is:\nhellop\n\n\nNo default channels\n")
	})

	t.Run("without message model", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(nil, nil)

		s.Subject.Call("user-id", "team-id")

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "No welcome message\n\n\nNo default channels\n")
	})

	t.Run("with empty message model", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(&pmodel.TeamWelcome{Message: ""}, nil)

		s.Subject.Call("user-id", "team-id")

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "No welcome message\n\n\nNo default channels\n")
	})

	t.Run("with empty message but some channels", func(t *testing.T) {
		s := setup()

		welcome := pmodel.TeamWelcome{
			Message:    "",
			ChannelIDs: []string{"ci1", "ci2"},
		}
		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(&welcome, nil)

		s.ChannelRepo.On("Get", "ci1").Return(&mmodel.Channel{Name: "cn1"}, nil)
		s.ChannelRepo.On("Get", "ci2").Return(&mmodel.Channel{Name: "cn2"}, nil)

		s.Subject.Call("user-id", "team-id")

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		t.Log(s.CommandMessenger.Calls)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "No welcome message\n\n\nDefault channels:\n- ~cn1\n- ~cn2\n")
	})
}
