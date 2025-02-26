package hook

import (
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/presenter"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/hook"
	"github.com/mattermost/mattermost/server/public/model"
)

type PublishedWelcomeNotifier struct{}

func (n *PublishedWelcomeNotifier) Execute(p handler.BotAPIProvider, channelMember *model.ChannelMember) {
	uc := hook.NotifyWithPublishedWelcome{
		Messenger:               p.Container().Messenger(),
		ChannelWelcomeRepo:      p.Container().ChannelWelcomeRepo(),
		ChannelRepo:             p.Container().ChannelRepo(),
		WelcomeMessagePresenter: &presenter.WelcomeMessagePresenter{UserRepo: p.Container().UserRepo()},
	}

	uc.Call(channelMember)
}
