package repo

import (
	"testing"

	mockAPI "github.com/evrone-erp/mattermost-plugin-welcomebot/server/mocks/mattermost/server/public/plugin/API"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/assert"
)

func setupUserRepo() (*UserRepo, *mockAPI.MockAPI) {
	api := new(mockAPI.MockAPI)
	r := UserRepo{
		botUserID: "bot-user-id",
		api:       api,
	}

	return &r, api
}

func TestUserRepoGetByID(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		r, api := setupUserRepo()

		user := &model.User{Id: "user-id"}
		api.On("GetUser", "user-id").Return(user, nil)

		result, err := r.GetByID("user-id")

		assert.Equal(t, user, result)
		assert.Nil(t, err)
	})
}
