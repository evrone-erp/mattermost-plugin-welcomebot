package command

import (
	"testing"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/stretchr/testify/require"
)

func TestRemoveTeamDefaultChannelsValidate(t *testing.T) {
	setup := func() *RemoveTeamDefaultChannels {
		return &RemoveTeamDefaultChannels{}
	}

	t.Run("happy path", func(t *testing.T) {
		cmd := setup()
		err := cmd.Validate([]string{"divided", "welcome"})

		require.NoError(t, err)
	})

	t.Run("with no parameters", func(t *testing.T) {
		cmd := setup()
		err := cmd.Validate([]string{})

		require.Error(t, err)
	})
}

func TestRemoveTeamDefaultChannelsParse(t *testing.T) {
	setup := func() *RemoveTeamDefaultChannels {
		return &RemoveTeamDefaultChannels{}
	}

	t.Run("happy path", func(t *testing.T) {
		cmd := setup()

		result, err := cmd.parse(&model.CommandArgs{Command: "/welcomebot remove_team_default_channels ~ch1 ~ch2"})

		require.NoError(t, err)
		require.NotEmpty(t, result)
		require.Equal(t, []string{"ch1", "ch2"}, result.channelNames)
	})

	t.Run("broken command", func(t *testing.T) {
		cmd := setup()

		result, err := cmd.parse(&model.CommandArgs{Command: "/welcomebot_remove_team_default_channels_team_~ch1_~ch2"})

		require.Error(t, err)
		require.Empty(t, result)
		require.Equal(t, "Unable to parse command /welcomebot_remove_team_default_channels_team_~ch1_~ch2", err.Error())
	})
}
