package hook

import (
	"testing"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	mmodel "github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/mock"
)

func TestNotifyWithPublishedWelcome(t *testing.T) {
	userID := "user-id"

	groupChannel := &mmodel.Channel{
		Id:   "group-channel-id",
		Type: mmodel.ChannelTypeOpen,
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
		Subject                 *NotifyWithPublishedWelcome
	}

	setup := func() *Setup {
		messenger := new(usecase.MockMessenger)
		channelWelcomeRepo := new(usecase.MockChannelWelcomeRepo)
		channelRepo := new(usecase.MockChannelRepo)
		welcomeMessagePresenter := new(usecase.MockWelcomeMessagePresenter)

		messenger.On("Post", mock.Anything, mock.Anything, mock.Anything).Return(nil)

		subject := &NotifyWithPublishedWelcome{
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
		s.ChannelRepo.On("Get", "group-channel-id").Return(groupChannel, nil)
		s.ChannelWelcomeRepo.On("GetPublishedChanelWelcome", "group-channel-id").Return(&pmodel.ChannelWelcome{Message: "hello"}, nil)
		s.WelcomeMessagePresenter.On("Render", "hello", "user-id").Return("Hello, friend!", nil)

		s.Subject.Call(channelMember)

		s.Messenger.AssertCalled(t, "Post", "group-channel-id", "Hello, friend!")
		s.Messenger.AssertNumberOfCalls(t, "Post", 1)
	})

	t.Run("no stored message", func(t *testing.T) {
		s := setup()
		s.ChannelRepo.On("Get", "group-channel-id").Return(groupChannel, nil)
		s.ChannelWelcomeRepo.On("GetPublishedChanelWelcome", "group-channel-id").Return(nil, nil)

		s.Subject.Call(channelMember)

		s.Messenger.AssertNumberOfCalls(t, "Post", 0)
	})

	t.Run("error while fetching published welcome", func(t *testing.T) {
		s := setup()
		s.ChannelRepo.On("Get", "group-channel-id").Return(groupChannel, nil)
		s.ChannelWelcomeRepo.On("GetPublishedChanelWelcome", "group-channel-id").Return(nil, &mmodel.AppError{Message: "foo"})

		s.Subject.Call(channelMember)

		s.Messenger.AssertNumberOfCalls(t, "Post", 0)
	})

	t.Run("errro while fetching channel", func(t *testing.T) {
		s := setup()
		s.ChannelRepo.On("Get", "group-channel-id").Return(nil, &mmodel.AppError{Message: "foo"})

		s.Subject.Call(channelMember)

		s.Messenger.AssertNumberOfCalls(t, "Post", 0)
	})
}
