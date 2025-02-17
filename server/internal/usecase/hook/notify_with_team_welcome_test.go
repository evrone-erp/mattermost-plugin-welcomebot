package hook

import (
	"testing"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	mmodel "github.com/mattermost/mattermost/server/public/model"
)

func TestNotifyWithTeamWelcome(t *testing.T) {
	directChannel := &mmodel.Channel{
		Id:   "direct-channel-id",
		Type: mmodel.ChannelTypeDirect,
	}

	teamMember := &mmodel.TeamMember{
		UserId: "user-id",
		TeamId: "team-id",
	}

	type Setup struct {
		Messenger               *usecase.MockMessenger
		TeamWelcomeRepo         *usecase.MockTeamWelcomeRepo
		ChannelRepo             *usecase.MockChannelRepo
		WelcomeMessagePresenter *usecase.MockWelcomeMessagePresenter
		Subject                 *NotifyWithTeamWelcome
	}

	setup := func() *Setup {
		messenger := new(usecase.MockMessenger)
		teamWelcomeRepo := new(usecase.MockTeamWelcomeRepo)
		channelRepo := new(usecase.MockChannelRepo)
		welcomeMessagePresenter := new(usecase.MockWelcomeMessagePresenter)

		subject := &NotifyWithTeamWelcome{
			Messenger:               messenger,
			TeamWelcomeRepo:         teamWelcomeRepo,
			ChannelRepo:             channelRepo,
			WelcomeMessagePresenter: welcomeMessagePresenter,
		}

		return &Setup{
			Messenger:               messenger,
			TeamWelcomeRepo:         teamWelcomeRepo,
			ChannelRepo:             channelRepo,
			WelcomeMessagePresenter: welcomeMessagePresenter,
			Subject:                 subject,
		}
	}

	t.Run("happy path", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(&pmodel.TeamWelcome{Message: "Hello"}, nil)
		s.ChannelRepo.On("GetDirect", "user-id").Return(directChannel, nil)
		s.Messenger.On("PostDirect", "direct-channel-id", "Hello darkness").Once().Return(nil)
		s.WelcomeMessagePresenter.On("Render", "Hello", "user-id").Return("Hello darkness", nil)

		s.Subject.Call(teamMember)

		s.Messenger.AssertNumberOfCalls(t, "PostDirect", 1)
	})

	t.Run("no message but channels", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(&pmodel.TeamWelcome{ChannelIDs: []string{"ci1", "ci2"}}, nil)

		s.ChannelRepo.On("AddMemberByUserID", "ci1", "user-id").Once().Return(nil, &mmodel.AppError{Message: "some"})
		s.ChannelRepo.On("AddMemberByUserID", "ci2", "user-id").Once().Return(&mmodel.ChannelMember{}, nil)
		s.Subject.Call(teamMember)

		s.Messenger.AssertNumberOfCalls(t, "PostDirect", 0)
		s.ChannelRepo.AssertNumberOfCalls(t, "AddMemberByUserID", 2)
	})

	t.Run("no message struct", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(nil, nil)

		s.Subject.Call(teamMember)

		s.Messenger.AssertNumberOfCalls(t, "PostDirect", 0)
	})

	t.Run("empty message struct", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(&pmodel.TeamWelcome{}, nil)

		s.Subject.Call(teamMember)

		s.Messenger.AssertNumberOfCalls(t, "PostDirect", 0)
	})

	t.Run("error while getting welcome message", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(nil, &mmodel.AppError{Message: "foo"})

		s.Subject.Call(teamMember)

		s.Messenger.AssertNumberOfCalls(t, "PostDirect", 0)
	})

	t.Run("error while fetching direct channel", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(&pmodel.TeamWelcome{Message: "Hello"}, nil)
		s.WelcomeMessagePresenter.On("Render", "Hello", "user-id").Return("Hello darkness", nil)
		s.ChannelRepo.On("GetDirect", "user-id").Return(nil, &mmodel.AppError{Message: "some"})

		s.Subject.Call(teamMember)

		s.Messenger.AssertNumberOfCalls(t, "PostDirect", 0)
	})

	t.Run("error while sending message", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(&pmodel.TeamWelcome{Message: "Hello"}, nil)
		s.WelcomeMessagePresenter.On("Render", "Hello", "user-id").Return("Hello darkness", nil)
		s.ChannelRepo.On("GetDirect", "user-id").Return(directChannel, nil)
		s.Messenger.On("PostDirect", "direct-channel-id", "Hello darkness").Once().Return(&mmodel.AppError{Message: "foo"})

		s.Subject.Call(teamMember)

		s.Messenger.AssertNumberOfCalls(t, "PostDirect", 1)
	})
}
