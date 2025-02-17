package repo

import (
	"testing"

	mockAPI "github.com/evrone-erp/mattermost-plugin-welcomebot/server/mocks/mattermost/server/public/plugin/API"
	"github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/assert"
)

func setupTeamRepo() (*TeamRepo, *mockAPI.MockAPI) {
	api := new(mockAPI.MockAPI)
	r := TeamRepo{
		botUserID: "bot-user-id",
		api:       api,
	}

	return &r, api
}

func TestTeamRepoGetByTeamID(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		r, api := setupTeamRepo()

		team := &model.Team{Id: "team-id"}
		api.On("GetTeam", "team-id").Return(team, nil)

		result, err := r.GetByTeamID("team-id")

		assert.Equal(t, team, result)
		assert.Nil(t, err)
	})
}

func TestTeamRepoGetByTeamName(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		r, api := setupTeamRepo()

		team := &model.Team{Name: "team-name"}
		api.On("GetTeamByName", "team-name").Return(team, nil)

		result, err := r.GetByTeamName("team-name")

		assert.Equal(t, team, result)
		assert.Nil(t, err)
	})
}
