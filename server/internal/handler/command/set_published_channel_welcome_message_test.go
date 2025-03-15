package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetPublishedChanelWelcomeMessageValidate(t *testing.T) {
	setup := func() *SetPublishedChanelWelcomeMessage {
		return &SetPublishedChanelWelcomeMessage{}
	}

	t.Run("happy path", func(t *testing.T) {
		c := setup()
		err := c.Validate([]string{"divided", "welcome"})

		assert.NoError(t, err)
	})

	t.Run("with extra parameters", func(t *testing.T) {
		c := setup()
		err := c.Validate([]string{})

		assert.Error(t, err)
	})
}
