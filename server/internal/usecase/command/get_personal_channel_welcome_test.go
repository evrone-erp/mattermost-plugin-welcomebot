package command

import (
	"testing"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	mmodel "github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/mock"
)

func TestGetPersonalChanelWelcome(t *testing.T) {
	channelID := "test-channel"

	type Setup struct {
		CommandMessenger        *usecase.MockCommandMessenger
		ChannelWelcomeRepo      *usecase.MockChannelWelcomeRepo
		WelcomeMessagePresenter *usecase.MockWelcomeMessagePresenter
		Subject                 *GetPersonalChanelWelcome
	}

	setup := func() *Setup {
		commandMessenger := new(usecase.MockCommandMessenger)
		channelWelcomeRepo := new(usecase.MockChannelWelcomeRepo)
		welcomeMessagePresenter := new(usecase.MockWelcomeMessagePresenter)

		commandMessenger.On("PostCommandResponse", mock.Anything).Return()

		subject := &GetPersonalChanelWelcome{
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
		s.ChannelWelcomeRepo.On("GetPersonalChanelWelcome", channelID).Return(&model.ChannelWelcome{Message: "Hello"}, nil)
		s.WelcomeMessagePresenter.On("Render", "Hello", "user-id").Return("Hello, friend!", nil)

		s.Subject.Call("user-id", channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Welcome message is:\nHello, friend!")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
	})

	t.Run("store returns an error", func(t *testing.T) {
		s := setup()
		s.ChannelWelcomeRepo.On("GetPersonalChanelWelcome", channelID).Return(nil, &mmodel.AppError{Message: "some error"})

		s.Subject.Call("user-id", channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "error occurred while retrieving the welcome message for the chanel: `some error`")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
	})

	t.Run("message not set", func(t *testing.T) {
		s := setup()
		s.ChannelWelcomeRepo.On("GetPersonalChanelWelcome", channelID).Return(nil, nil)

		s.Subject.Call("user-id", channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "welcome message has not been set yet")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
	})
}
