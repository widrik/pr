package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig(t *testing.T) {
	_, err := Init()

	t.Run("init has no errors", func(t *testing.T) {
		require.Equal(t, nil, err)
	})
}
