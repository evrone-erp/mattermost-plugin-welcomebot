package hook

import (
	"testing"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	mmodel "github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/mock"
)

func TestNotifyWithPersonalWelcome(t *testing.T) {
	userID := "user-id"

	groupChannel := &mmodel.Channel{
		Id:   "group-channel-id",
		Type: mmodel.ChannelTypeOpen,
	}

	directChannel := &mmodel.Channel{
		Id:   "direct-channel-id",
		Type: mmodel.ChannelTypeDirect,
	}

	channelMember := &mmodel.ChannelMember{
		ChannelId: "group-channel-id",
		UserId:    userID,
	}

	type Setup struct {
		Messenger               *usecase.MockMessenger
		ChannelWelcomeRepo      *usecase.MockChannelWelcomeRepo
		ChannelRepo             *usecase.MockChannelRepo
		WelcomeMessagePresenter *usecase.MockWelcomeMessagePresenter
		Subject                 *NotifyWithPersonalWelcome
	}

	setup := func() *Setup {
		messenger := new(usecase.MockMessenger)
		channelWelcomeRepo := new(usecase.MockChannelWelcomeRepo)
		channelRepo := new(usecase.MockChannelRepo)
		welcomeMessagePresenter := new(usecase.MockWelcomeMessagePresenter)

		messenger.On("PostChannelEphemeral", mock.Anything, mock.Anything, mock.Anything).Return()

		subject := &NotifyWithPersonalWelcome{
			Messenger:               messenger,
			ChannelWelcomeRepo:      channelWelcomeRepo,
			ChannelRepo:             channelRepo,
			WelcomeMessagePresenter: welcomeMessagePresenter,
		}

		return &Setup{
			Messenger:               messenger,
			ChannelWelcomeRepo:      channelWelcomeRepo,
			ChannelRepo:             channelRepo,
			WelcomeMessagePresenter: welcomeMessagePresenter,
			Subject:                 subject,
		}
	}

	t.Run("happy path", func(t *testing.T) {
		s := setup()

		s.Messenger.On("PostDirect", mock.Anything, mock.Anything).Return(nil)
		s.ChannelRepo.On("Get", "group-channel-id").Return(groupChannel, nil)
		s.ChannelRepo.On("GetDirect", userID).Return(directChannel, nil)
		s.ChannelWelcomeRepo.On("GetPersonalChanelWelcome", "group-channel-id").Return(&pmodel.ChannelWelcome{Message: "hello"}, nil)
		s.WelcomeMessagePresenter.On("Render", "hello", "user-id").Return("Hello, friend!", nil)

		s.Subject.Call(channelMember)

		s.Messenger.AssertCalled(t, "PostDirect", "direct-channel-id", "Hello, friend!")
		s.Messenger.AssertCalled(t, "PostChannelEphemeral", "group-channel-id", "user-id", "Hello, friend!")

		s.Messenger.AssertNumberOfCalls(t, "PostDirect", 1)
		s.Messenger.AssertNumberOfCalls(t, "PostChannelEphemeral", 1)
	})

	t.Run("no stored message", func(t *testing.T) {
		s := setup()
		s.Messenger.On("PostDirect", mock.Anything, mock.Anything).Return(nil)
		s.ChannelRepo.On("Get", "group-channel-id").Return(groupChannel, nil)
		s.ChannelWelcomeRepo.On("GetPersonalChanelWelcome", "group-channel-id").Return(nil, nil)
		s.ChannelRepo.On("GetDirect", userID).Return(directChannel, nil)

		s.Subject.Call(channelMember)

		s.Messenger.AssertNumberOfCalls(t, "PostDirect", 0)
		s.Messenger.AssertNumberOfCalls(t, "PostChannelEphemeral", 0)
	})

	t.Run("error during direct message", func(t *testing.T) {
		s := setup()
		s.ChannelRepo.On("Get", "group-channel-id").Return(groupChannel, nil)
		s.ChannelRepo.On("GetDirect", userID).Return(directChannel, nil)
		s.WelcomeMessagePresenter.On("Render", "Hello", "user-id").Return("Hello, friend!", nil)
		s.ChannelWelcomeRepo.On("GetPersonalChanelWelcome", "group-channel-id").Return(&pmodel.ChannelWelcome{Message: "Hello"}, nil)
		s.Messenger.On("PostDirect", "direct-channel-id", "Hello, friend!").Return(&mmodel.AppError{Message: "foo"})

		s.Subject.Call(channelMember)

		s.Messenger.AssertCalled(t, "PostChannelEphemeral", "group-channel-id", "user-id", "Hello, friend!")

		s.Messenger.AssertNumberOfCalls(t, "PostDirect", 1)
		s.Messenger.AssertNumberOfCalls(t, "PostChannelEphemeral", 1)
	})

	t.Run("error fetching direct", func(t *testing.T) {
		s := setup()
		s.ChannelRepo.On("Get", "group-channel-id").Return(groupChannel, nil)
		s.ChannelRepo.On("GetDirect", userID).Return(nil, &mmodel.AppError{Message: "foo"})
		s.ChannelWelcomeRepo.On("GetPersonalChanelWelcome", "group-channel-id").Return(&pmodel.ChannelWelcome{Message: "Hello, friend!"}, nil)

		s.Subject.Call(channelMember)

		s.Messenger.AssertNumberOfCalls(t, "PostDirect", 0)
		s.Messenger.AssertNumberOfCalls(t, "PostChannelEphemeral", 0)
	})

	t.Run("error while fetching personal welcome", func(t *testing.T) {
		s := setup()
		s.ChannelRepo.On("Get", "group-channel-id").Return(groupChannel, nil)
		s.ChannelWelcomeRepo.On("GetPersonalChanelWelcome", "group-channel-id").Return(nil, &mmodel.AppError{Message: "foo"})

		s.Subject.Call(channelMember)

		s.Messenger.AssertNumberOfCalls(t, "PostDirect", 0)
		s.Messenger.AssertNumberOfCalls(t, "PostChannelEphemeral", 0)
	})

	t.Run("errro while fetching channel", func(t *testing.T) {
		s := setup()
		s.ChannelRepo.On("Get", "group-channel-id").Return(nil, &mmodel.AppError{Message: "foo"})

		s.Subject.Call(channelMember)

		s.Messenger.AssertNumberOfCalls(t, "PostDirect", 0)
		s.Messenger.AssertNumberOfCalls(t, "PostChannelEphemeral", 0)
	})
}
