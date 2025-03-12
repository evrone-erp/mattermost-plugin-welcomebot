package utils

import (
	"testing"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/assert"
)

func TestIsChannelWithWelcomeSupport(t *testing.T) {
	t.Run("supported types", func(t *testing.T) {
		var supportedChannels = []struct {
			msg     string
			channel *model.Channel
		}{
			{"Open channel", &model.Channel{Type: model.ChannelTypeOpen}},
			{"Private channel", &model.Channel{Type: model.ChannelTypePrivate}},
		}

		for _, tc := range supportedChannels {
			assert.True(t, IsChannelWithWelcomeSupport(tc.channel), tc.msg)
		}
	})

	t.Run("unsupported types", func(t *testing.T) {
		var supportedChannels = []struct {
			msg     string
			channel *model.Channel
		}{
			{"Direct channel", &model.Channel{Type: model.ChannelTypeDirect}},
		}

		for _, tc := range supportedChannels {
			assert.False(t, IsChannelWithWelcomeSupport(tc.channel), tc.msg)
		}
	})
}
