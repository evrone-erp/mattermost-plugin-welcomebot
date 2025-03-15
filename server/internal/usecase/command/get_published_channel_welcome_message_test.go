package command

import (
	"testing"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	mmodel "github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/mock"
)

func TestGetPublishedChanelWelcome(t *testing.T) {
	channelID := "test-channel"

	type Setup struct {
		CommandMessenger        *usecase.MockCommandMessenger
		ChannelWelcomeRepo      *usecase.MockChannelWelcomeRepo
		WelcomeMessagePresenter *usecase.MockWelcomeMessagePresenter
		Subject                 *GetPublishedChanelWelcomeMessage
	}

	setup := func() *Setup {
		commandMessenger := new(usecase.MockCommandMessenger)
		channelWelcomeRepo := new(usecase.MockChannelWelcomeRepo)
		welcomeMessagePresenter := new(usecase.MockWelcomeMessagePresenter)

		commandMessenger.On("PostCommandResponse", mock.Anything).Return()

		subject := &GetPublishedChanelWelcomeMessage{
			CommandMessenger:        commandMessenger,
			ChannelWelcomeRepo:      channelWelcomeRepo,
			WelcomeMessagePresenter: welcomeMessagePresenter,
		}

		return &Setup{
			CommandMessenger:        commandMessenger,
			ChannelWelcomeRepo:      channelWelcomeRepo,
			WelcomeMessagePresenter: welcomeMessagePresenter,
			Subject:                 subject,
		}
	}

	t.Run("message exists", func(t *testing.T) {
		s := setup()
		s.ChannelWelcomeRepo.On("GetPublishedChanelWelcome", channelID).Return(&pmodel.ChannelWelcome{Message: "Hello"}, nil)

		s.WelcomeMessagePresenter.On("Render", "Hello", "user-id").Return("Hello, friend!", nil)
		s.Subject.Call("user-id", channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome message is:\nHello, friend!")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
	})

	t.Run("store returns an error", func(t *testing.T) {
		s := setup()
		s.ChannelWelcomeRepo.On("GetPublishedChanelWelcome", channelID).Return(nil, &mmodel.AppError{Message: "some error"})

		s.Subject.Call("user-id", channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "error occurred while retrieving the welcome message for the chanel: `some error`")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
	})

	t.Run("message not set", func(t *testing.T) {
		s := setup()
		s.ChannelWelcomeRepo.On("GetPublishedChanelWelcome", channelID).Return(nil, nil)

		s.Subject.Call("user-id", channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "welcome message has not been set yet")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
	})
}
