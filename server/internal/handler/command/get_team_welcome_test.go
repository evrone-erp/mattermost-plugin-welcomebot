package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTeamWelcomeValidate(t *testing.T) {
	setup := func() *GetTeamWelcomeSettings {
		return &GetTeamWelcomeSettings{}
	}

	t.Run("happy path", func(t *testing.T) {
		cmd := setup()
		err := cmd.Validate([]string{})

		assert.NoError(t, err)
	})

	t.Run("with many parameters", func(t *testing.T) {
		cmd := setup()
		err := cmd.Validate([]string{"some", "thing"})

		assert.Error(t, err)
	})
}
