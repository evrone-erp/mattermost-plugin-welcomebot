package command

import (
	"testing"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/mock"
)

func TestSetPersonalChanelWelcome(t *testing.T) {
	channelID := "test-channel"

	type Setup struct {
		CommandMessenger   *usecase.MockCommandMessenger
		ChannelWelcomeRepo *usecase.MockChannelWelcomeRepo
		ChannelRepo        *usecase.MockChannelRepo
		Subject            *SetPersonalChanelWelcomeMessage
	}

	setup := func() *Setup {
		messenger := new(usecase.MockCommandMessenger)
		channelWelcomeRepo := new(usecase.MockChannelWelcomeRepo)
		channelRepo := new(usecase.MockChannelRepo)

		messenger.On("PostCommandResponse", mock.Anything).Return()

		subject := &SetPersonalChanelWelcomeMessage{
			CommandMessenger:   messenger,
			ChannelWelcomeRepo: channelWelcomeRepo,
			ChannelRepo:        channelRepo,
		}

		return &Setup{
			CommandMessenger:   messenger,
			ChannelWelcomeRepo: channelWelcomeRepo,
			ChannelRepo:        channelRepo,
			Subject:            subject,
		}
	}

	validChannel := &model.Channel{
		Id:   channelID,
		Type: model.ChannelTypeOpen,
	}

	validCommand := "set_personal_channel_welcome_message foo bar keke   "

	t.Run("happy path", func(t *testing.T) {
		s := setup()

		s.ChannelWelcomeRepo.On("SetPersonalChanelWelcome", mock.Anything, mock.Anything).Return(nil)
		s.ChannelRepo.On("Get", channelID).Return(validChannel, nil)

		s.Subject.Call(validCommand, channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "stored the welcome message:\nfoo bar keke")
		s.ChannelWelcomeRepo.AssertCalled(t, "SetPersonalChanelWelcome", "test-channel", "foo bar keke")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.ChannelWelcomeRepo.AssertNumberOfCalls(t, "SetPersonalChanelWelcome", 1)
	})

	t.Run("with broken command", func(t *testing.T) {
		s := setup()

		s.ChannelWelcomeRepo.On("SetPersonalChanelWelcome", mock.Anything, mock.Anything).Return(nil)
		s.ChannelRepo.On("Get", channelID).Return(validChannel, nil)

		s.Subject.Call("kek", channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "error ocured while parsing command kek")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.ChannelWelcomeRepo.AssertNumberOfCalls(t, "SetPersonalChanelWelcome", 0)
	})

	t.Run("with a command without message", func(t *testing.T) {
		s := setup()

		s.ChannelWelcomeRepo.On("SetPersonalChanelWelcome", mock.Anything, mock.Anything).Return(nil)
		s.ChannelRepo.On("Get", channelID).Return(validChannel, nil)

		s.Subject.Call("set_personal_channel_welcome_message     ", channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "unable to store empty message")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.ChannelWelcomeRepo.AssertNumberOfCalls(t, "SetPersonalChanelWelcome", 0)
	})

	t.Run("direct channel", func(t *testing.T) {
		s := setup()

		s.ChannelWelcomeRepo.On("SetPersonalChanelWelcome", mock.Anything, mock.Anything).Return(nil)
		channel := &model.Channel{
			Id:   channelID,
			Type: model.ChannelTypeDirect,
		}
		s.ChannelRepo.On("Get", channelID).Return(channel, nil)

		s.Subject.Call(validCommand, channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "Channel type is not supported")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.ChannelWelcomeRepo.AssertNumberOfCalls(t, "SetPersonalChanelWelcome", 0)
	})

	t.Run("persist message error", func(t *testing.T) {
		s := setup()

		s.ChannelWelcomeRepo.On("SetPersonalChanelWelcome", mock.Anything, mock.Anything).Return(&model.AppError{Message: "persist error"})
		s.ChannelRepo.On("Get", channelID).Return(validChannel, nil)

		s.Subject.Call(validCommand, channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "error occurred while storing the welcome message for the chanel: `persist error`")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.ChannelWelcomeRepo.AssertNumberOfCalls(t, "SetPersonalChanelWelcome", 1)
	})

	t.Run("retrieve channel error", func(t *testing.T) {
		s := setup()
		s.ChannelRepo.On("Get", channelID).Return(nil, &model.AppError{Message: "receiving error"})

		s.Subject.Call(validCommand, channelID)

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", "error occurred while checking the type of the chanelId `test-channel`: `receiving error`")
		s.CommandMessenger.AssertNumberOfCalls(t, "PostCommandResponse", 1)
		s.ChannelWelcomeRepo.AssertNumberOfCalls(t, "SetPersonalChanelWelcome", 0)
	})
}
