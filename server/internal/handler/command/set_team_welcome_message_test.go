package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetTeamWelcomeMessageValidate(t *testing.T) {
	setup := func() *SetTeamWelcomeMessage {
		return &SetTeamWelcomeMessage{}
	}

	t.Run("happy path", func(t *testing.T) {
		cmd := setup()
		err := cmd.Validate([]string{"divided", "welcome"})

		assert.NoError(t, err)
	})

	t.Run("with no parameters", func(t *testing.T) {
		cmd := setup()
		err := cmd.Validate([]string{})

		assert.Error(t, err)
	})
}
