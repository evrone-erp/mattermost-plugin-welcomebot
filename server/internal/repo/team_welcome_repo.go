package repo

import (
	"encoding/json"
	"fmt"

	pmodel "github.com/evrone-erp/mattermost-plugin-welcomebot/server/internal/model"
	mmodel "github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
)

type TeamWelcomeDTO struct {
	Message    string   `json:"message,omitempty"`
	ChannelIDs []string `json:"channelIDs,omitempty"`
}

type TeamWelcomeRepo struct {
	api plugin.API
}

func NewTeamWelcomeRepo(p BotAPIProvider) *TeamWelcomeRepo {
	return &TeamWelcomeRepo{
		api: p.APIHandle(),
	}
}

const (
	generalTeamPrefix = "teamwelcome"
)

func (r *TeamWelcomeRepo) GetTeamWelcome(teamID string) (*pmodel.TeamWelcome, *mmodel.AppError) {
	key := teamKey(teamID)

	data, appErr := r.api.KVGet(key)

	if appErr != nil {
		return nil, appErr
	}

	if len(data) == 0 {
		return nil, nil
	}

	var dto TeamWelcomeDTO

	err := json.Unmarshal(data, &dto)

	if err != nil {
		return nil, &mmodel.AppError{Message: "Broken JSON format in KVStore"}
	}

	return &pmodel.TeamWelcome{
		ID:         key,
		Message:    dto.Message,
		ChannelIDs: dto.ChannelIDs,
		TeamID:     teamID,
	}, nil
}

func (r *TeamWelcomeRepo) SetTeamWelcome(teamID string, welcome *pmodel.TeamWelcome) *mmodel.AppError {
	dto := TeamWelcomeDTO{
		Message:    welcome.Message,
		ChannelIDs: welcome.ChannelIDs,
	}

	data, err := json.Marshal(dto)

	if err != nil {
		return &mmodel.AppError{Message: err.Error()}
	}

	key := teamKey(teamID)
	appErr := r.api.KVSet(key, data)

	if appErr != nil {
		return appErr
	}

	return nil
}

func teamKey(teamID string) string {
	return fmt.Sprintf("%s:%s", generalTeamPrefix, teamID)
}
