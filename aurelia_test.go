package aurelia_test

import (
	"testing"

	"github.com/kodefluence/aurelia"
	"github.com/stretchr/testify/assert"
)

func TestAurelia(t *testing.T) {
	hash := aurelia.Hash("credential", "key")

	t.Run("Authentication should success", func(t *testing.T) {

		assert.True(t, aurelia.Authenticate("credential", "key", hash))
	})

	t.Run("Authentication should not success", func(t *testing.T) {

		assert.False(t, aurelia.Authenticate("credential", "wrong_key", hash))
	})
}
