package policy

import (
	"testing"

	mockAPI "github.com/evrone-erp/mattermost-plugin-welcomebot/server/mocks/mattermost/server/public/plugin/API"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin/plugintest/mock"
	"github.com/stretchr/testify/require"
)

func TestCanManageTeam(t *testing.T) {
	type TestCase struct {
		api    *mockAPI.MockAPI
		policy *Policy
	}

	setup := func() *TestCase {
		mockAPI := new(mockAPI.MockAPI)
		policy := Policy{api: mockAPI}

		return &TestCase{
			api:    mockAPI,
			policy: &policy,
		}
	}

	userID := "usser-id"
	teamID := "team-id"

	t.Run("when the user is a system_admin", func(t *testing.T) {
		tc := setup()

		tc.api.On("GetUser", userID).Return(&model.User{Roles: model.SystemAdminRoleId}, nil)
		require.True(t, tc.policy.CanManageTeam(userID, teamID))
	})

	t.Run("when the user is not an admin but has a team permission", func(t *testing.T) {
		tc := setup()

		tc.api.On("GetUser", userID).Return(&model.User{Roles: model.SystemUserRoleId}, nil)
		tc.api.On("HasPermissionToTeam", userID, teamID, mock.Anything).Return(true)
		require.True(t, tc.policy.CanManageTeam(userID, teamID))
	})

	t.Run("when the user is not an admin and has no additional permissions", func(t *testing.T) {
		tc := setup()

		tc.api.On("GetUser", userID).Return(&model.User{Roles: model.SystemUserRoleId}, nil)
		tc.api.On("HasPermissionToTeam", userID, teamID, mock.Anything).Return(false)
		require.False(t, tc.policy.CanManageTeam(userID, teamID))
	})
}
