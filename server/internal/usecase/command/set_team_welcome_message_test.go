package command

import (
	"testing"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	mmodel "github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/mock"
)

func TestSetTeamWelcomeMessage(t *testing.T) {
	type Setup struct {
		Messenger       *usecase.MockCommandMessenger
		TeamWelcomeRepo *usecase.MockTeamWelcomeRepo
		Subject         *SetTeamWelcomeMessage
	}

	setup := func() *Setup {
		messenger := new(usecase.MockCommandMessenger)
		teamWelcomeRepo := new(usecase.MockTeamWelcomeRepo)

		subject := &SetTeamWelcomeMessage{
			Messenger:       messenger,
			TeamWelcomeRepo: teamWelcomeRepo,
		}

		messenger.On("PostCommandResponse", mock.Anything).Return()

		return &Setup{
			Subject:         subject,
			Messenger:       messenger,
			TeamWelcomeRepo: teamWelcomeRepo,
		}
	}

	t.Run("happy path without previous welcome", func(t *testing.T) {
		s := setup()
		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(nil, nil)

		updatedWelcome := &pmodel.TeamWelcome{
			Message: "foo bar kek",
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(nil)

		s.Subject.Call("team-id", "foo bar kek")

		s.Messenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.Messenger.AssertCalled(t, "PostCommandResponse", "Welcome message was updated")
	})

	t.Run("happy path with bad previous welcome messsage", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(nil, &mmodel.AppError{Message: "broken"})

		updatedWelcome := &pmodel.TeamWelcome{
			Message: "foo bar kek",
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(nil)

		s.Subject.Call("team-id", "foo bar kek")

		s.Messenger.AssertNumberOfCalls(t, "PostCommandResponse", 2)
		s.Messenger.AssertCalled(t, "PostCommandResponse", "Error while fetching current welcome team-id: broken, trying to rewrite...")
		s.Messenger.AssertCalled(t, "PostCommandResponse", "Welcome message was updated")
	})

	t.Run("with a saving error", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(nil, nil)

		updatedWelcome := &pmodel.TeamWelcome{
			Message: "foo bar kek",
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(&mmodel.AppError{Message: "fail"})

		s.Subject.Call("team-id", "foo bar kek")

		s.Messenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.Messenger.AssertCalled(t, "PostCommandResponse", "Error while saving team team-id: fail")
	})
}
