package policy

import (
	"strings"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
)

type Policy struct {
	api plugin.API
}

func NewPolicy(p BotAPIProvider) *Policy {
	return &Policy{
		api: p.APIHandle(),
	}
}

func (p *Policy) CanManageChannel(userID string, channelID string) bool {
	return p.api.HasPermissionToChannel(userID, channelID, model.PermissionManageChannelRoles)
}

func (p *Policy) CanManageTeam(userID string, teamID string) bool {
	if p.hasRole(userID, model.SystemAdminRoleId) {
		return true
	}

	return p.api.HasPermissionToTeam(userID, teamID, model.PermissionManageTeamRoles)
}

func (p *Policy) IsSysadmin(userID string) bool {
	return p.hasRole(userID, model.SystemAdminRoleId)
}

func (p *Policy) hasRole(userID string, role string) bool {
	user, appErr := p.api.GetUser(userID)
	if appErr != nil {
		p.api.LogError("failed to query user", "user_id", userID, "err", appErr)
		return false
	}

	if !strings.Contains(user.Roles, role) {
		return false
	}
	return true
}
