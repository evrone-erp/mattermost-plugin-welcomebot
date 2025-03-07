package repo

import (
	"github.com/stretchr/testify/assert"

	"testing"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	mockAPI "github.com/evrone-erp/mattermost-plugin-welcomebot/server/mocks/mattermost/server/public/plugin/API"
	mmodel "github.com/mattermost/mattermost/server/public/model"
)

func setupTeamWelcomeRepo() (*TeamWelcomeRepo, *mockAPI.MockAPI) {
	api := new(mockAPI.MockAPI)
	r := TeamWelcomeRepo{
		api: api,
	}

	return &r, api
}

func TestGetTeamWelcome(t *testing.T) {
	t.Run("happy path with just a message", func(t *testing.T) {
		r, api := setupTeamWelcomeRepo()
		data := `{"message": "hello"}`
		api.On("KVGet", "teamwelcome:team-id").Return([]byte(data), nil)
		result, err := r.GetTeamWelcome("team-id")

		expected := &pmodel.TeamWelcome{
			ID:      "teamwelcome:team-id",
			TeamID:  "team-id",
			Message: "hello",
		}

		assert.Equal(t, expected, result)
		assert.Nil(t, err)
	})

	t.Run("happy path with message and channels", func(t *testing.T) {
		r, api := setupTeamWelcomeRepo()
		data := `{"message": "hello", "channelIDs": ["one", "two"]}`
		api.On("KVGet", "teamwelcome:team-id").Return([]byte(data), nil)
		result, err := r.GetTeamWelcome("team-id")

		expected := &pmodel.TeamWelcome{
			ID:         "teamwelcome:team-id",
			TeamID:     "team-id",
			Message:    "hello",
			ChannelIDs: []string{"one", "two"},
		}

		assert.Equal(t, expected, result)
		assert.Nil(t, err)
	})

	t.Run("empty struct", func(t *testing.T) {
		r, api := setupTeamWelcomeRepo()
		api.On("KVGet", "teamwelcome:team-id").Return([]byte{}, nil)
		result, err := r.GetTeamWelcome("team-id")

		assert.Nil(t, result)
		assert.Nil(t, err)
	})

	t.Run("broken json", func(t *testing.T) {
		r, api := setupTeamWelcomeRepo()
		data := `{"message": "hello"`
		api.On("KVGet", "teamwelcome:team-id").Return([]byte(data), nil)
		result, err := r.GetTeamWelcome("team-id")

		assert.Nil(t, result)
		assert.Equal(t, "Broken JSON format in KVStore", err.Message)
	})

	t.Run("error", func(t *testing.T) {
		r, api := setupTeamWelcomeRepo()
		api.On("KVGet", "teamwelcome:team-id").Return(nil, &mmodel.AppError{Message: "foo"})
		_, err := r.GetTeamWelcome("team-id")

		assert.Equal(t, "foo", err.Error())
	})
}

func TestSetTeamWelcome(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		r, api := setupTeamWelcomeRepo()
		data := `{"message":"hello"}`
		api.On("KVSet", "teamwelcome:team-id", []byte(data)).Once().Return(nil)
		welcome := pmodel.TeamWelcome{
			Message: "hello",
		}

		err := r.SetTeamWelcome("team-id", &welcome)

		assert.Nil(t, err)
	})

	t.Run("happy path with channels", func(t *testing.T) {
		r, api := setupTeamWelcomeRepo()
		data := `{"message":"hello","channelIDs":["foo","bar"]}`
		api.On("KVSet", "teamwelcome:team-id", []byte(data)).Once().Return(nil)
		welcome := pmodel.TeamWelcome{
			Message:    "hello",
			ChannelIDs: []string{"foo", "bar"},
		}

		err := r.SetTeamWelcome("team-id", &welcome)

		assert.Nil(t, err)
	})

	t.Run("kvset error", func(t *testing.T) {
		r, api := setupTeamWelcomeRepo()
		data := `{"message":"hello"}`
		api.On("KVSet", "teamwelcome:team-id", []byte(data)).Once().Return(&mmodel.AppError{Message: "something"})
		welcome := pmodel.TeamWelcome{
			Message: "hello",
		}

		err := r.SetTeamWelcome("team-id", &welcome)
		assert.Equal(t, "something", err.Message)
	})
}
