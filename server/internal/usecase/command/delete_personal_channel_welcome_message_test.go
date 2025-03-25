package command

import (
	"testing"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	mmodel "github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/mock"
)

func TestDeletePersonalChanelWelcome(t *testing.T) {
	channelID := "test-channel"

	type Setup struct {
		CommandMessenger   *usecase.MockCommandMessenger
		ChannelWelcomeRepo *usecase.MockChannelWelcomeRepo
		Subject            *DeletePersonalChanelWelcomeMessage
	}

	setup := func() *Setup {
		commandMessenger := new(usecase.MockCommandMessenger)
		channelWelcomeRepo := new(usecase.MockChannelWelcomeRepo)

		commandMessenger.On("PostCommandResponse", mock.Anything).Return()

		subject := &DeletePersonalChanelWelcomeMessage{
			CommandMessenger:   commandMessenger,
			ChannelWelcomeRepo: channelWelcomeRepo,
		}

		return &Setup{
			CommandMessenger:   commandMessenger,
			ChannelWelcomeRepo: channelWelcomeRepo,
			Subject:            subject,
		}
	}

	t.Run("happy path", func(t *testing.T) {
		s := setup()
		s.ChannelWelcomeRepo.On("GetPersonalChanelWelcome", channelID).Return(&pmodel.ChannelWelcome{Message: "Hello, friend"}, nil, nil)
		s.ChannelWelcomeRepo.On("DeletePersonalChanelWelcome", channelID).Return(nil)

		s.Subject.Call(channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "welcome message has been deleted")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
	})

	t.Run("no message", func(t *testing.T) {
		s := setup()
		s.ChannelWelcomeRepo.On("GetPersonalChanelWelcome", channelID).Return(nil, nil)

		s.Subject.Call(channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "welcome message has not been set yet")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
	})

	t.Run("error while deleting from store", func(t *testing.T) {
		s := setup()
		s.ChannelWelcomeRepo.On("GetPersonalChanelWelcome", channelID).Return(&pmodel.ChannelWelcome{Message: "Hello, friend"}, nil)
		s.ChannelWelcomeRepo.On("DeletePersonalChanelWelcome", channelID).Return(&mmodel.AppError{Message: "DAMN"})

		s.Subject.Call(channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "error occurred while deleting the welcome message for the chanel: `DAMN`")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
	})

	t.Run("error while retreiving from store", func(t *testing.T) {
		s := setup()
		s.ChannelWelcomeRepo.On("GetPersonalChanelWelcome", channelID).Return(nil, &mmodel.AppError{Message: "FOO"})

		s.Subject.Call(channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "error occurred while retrieving the welcome message for the chanel: `FOO`")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
	})
}
