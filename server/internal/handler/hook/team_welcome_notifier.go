package hook

import (
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/handler"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/presenter"
	"github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/usecase/hook"
	"github.com/mattermost/mattermost/server/public/model"
)

type TeamWelcomeNotifier struct{}

func (n *TeamWelcomeNotifier) Execute(p handler.BotAPIProvider, member *model.TeamMember, _ *model.User) {
	uc := hook.NotifyWithTeamWelcome{
		Messenger:               p.Container().Messenger(),
		TeamWelcomeRepo:         p.Container().TeamWelcomeRepo(),
		ChannelRepo:             p.Container().ChannelRepo(),
		WelcomeMessagePresenter: &presenter.WelcomeMessagePresenter{UserRepo: p.Container().UserRepo()},
	}

	uc.Call(member)
}
