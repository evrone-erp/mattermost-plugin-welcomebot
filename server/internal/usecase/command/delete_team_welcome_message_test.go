package command

import (
	"testing"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	mmodel "github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/mock"
)

func TestDeleteTeamWelcomeMessage(t *testing.T) {
	type Setup struct {
		CommandMessenger *usecase.MockCommandMessenger
		TeamWelcomeRepo  *usecase.MockTeamWelcomeRepo
		Subject          *DeleteTeamWelcomeMessage
	}

	setup := func() *Setup {
		commandMessenger := new(usecase.MockCommandMessenger)
		teamWelcomeRepo := new(usecase.MockTeamWelcomeRepo)

		commandMessenger.On("PostCommandResponse", mock.Anything).Return()

		subject := &DeleteTeamWelcomeMessage{
			CommandMessenger: commandMessenger,
			TeamWelcomeRepo:  teamWelcomeRepo,
		}

		return &Setup{
			CommandMessenger: commandMessenger,
			TeamWelcomeRepo:  teamWelcomeRepo,
			Subject:          subject,
		}
	}

	t.Run("happy path without previous welcome", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(nil, nil)

		updatedWelcome := &pmodel.TeamWelcome{
			Message: "",
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(nil)

		s.Subject.Call("team-id")

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome message was deleted")
	})

	t.Run("with previouis welcome", func(t *testing.T) {
		s := setup()

		initialWelcome := &pmodel.TeamWelcome{
			Message: "",
		}

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(initialWelcome, nil)

		updatedWelcome := &pmodel.TeamWelcome{
			Message: "",
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(nil)

		s.Subject.Call("team-id")

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome message was deleted")
	})

	t.Run("happy path with bad previous welcome messsage", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(nil, &mmodel.AppError{Message: "broken"})

		updatedWelcome := &pmodel.TeamWelcome{
			Message: "",
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(nil)

		s.Subject.Call("team-id")

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 2)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Error while fetching current welcome team-id: broken, trying to rewrite...")
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome message was deleted")
	})

	t.Run("with a saving error", func(t *testing.T) {
		s := setup()

		s.TeamWelcomeRepo.On("GetTeamWelcome", "team-id").Return(nil, nil)

		updatedWelcome := &pmodel.TeamWelcome{
			Message: "",
		}
		s.TeamWelcomeRepo.On("SetTeamWelcome", "team-id", updatedWelcome).Return(&mmodel.AppError{Message: "fail"})

		s.Subject.Call("team-id")

		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Error while saving team team-id: fail")
	})
}
