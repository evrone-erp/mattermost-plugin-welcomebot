package command

import (
	"testing"

	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/mock"
)

func TestListChannelWelcomes(t *testing.T) {
	type Setup struct {
		CommandMessenger   *usecase.MockCommandMessenger
		ChannelWelcomeRepo *usecase.MockChannelWelcomeRepo
		ChannelRepo        *usecase.MockChannelRepo
		Subject            *ListChannelWelcomes
	}

	setup := func() *Setup {
		messenger := new(usecase.MockCommandMessenger)
		channelWelcomeRepo := new(usecase.MockChannelWelcomeRepo)
		channelRepo := new(usecase.MockChannelRepo)

		messenger.On("PostCommandResponse", mock.Anything).Return()

		subject := &ListChannelWelcomes{
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

	t.Run("complete happy path", func(t *testing.T) {
		s := setup()

		personalIDs := []string{"personalid-1", "personalid-2"}
		publishedIDs := []string{"published-1", "garbage"}

		s.ChannelWelcomeRepo.On("ListChannelsWithWelcome").Return(personalIDs, publishedIDs, nil)
		s.ChannelRepo.On("Get", "personalid-1").Return(&model.Channel{Id: "personalid-1", Name: "pe1"}, nil)
		s.ChannelRepo.On("Get", "personalid-2").Return(&model.Channel{Id: "personalid-2", Name: "pe2"}, nil)
		s.ChannelRepo.On("Get", "published-1").Return(&model.Channel{Id: "published-1", Name: "pu1"}, nil)
		s.ChannelRepo.On("Get", "garbage").Return(nil, &model.AppError{Message: "whatever"})

		s.Subject.Call()
		expectedMessage := "Channels with published welcome message:\n~pu1\nChannels with personal welcome message:\n~pe1\n~pe2\n"

		s.CommandMessenger.AssertCalled(t, "PostCommandResponse", expectedMessage)
	})
}
