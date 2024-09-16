package gopt_test

import (
	"testing"

	"github.com/piotrpersona/gopt"
	"github.com/stretchr/testify/require"
)

func Test_Optional(t *testing.T) {
	t.Parallel()

	t.Run("Test some", func(t *testing.T) {
		t.Parallel()

		opt := gopt.Some("car")
		val, err := opt.Get()
		require.NoError(t, err)
		require.Equal(t, "car", val)
	})
	t.Run("Test none", func(t *testing.T) {
		t.Parallel()

		opt := gopt.None[string]()
		val, err := opt.Get()
		require.Equal(t, "", val)
		require.Error(t, err)
		require.ErrorIs(t, err, gopt.NoneErr)
	})
	t.Run("Test default", func(t *testing.T) {
		t.Parallel()

		opt := gopt.None[string]()
		val := opt.Default("boat")
		require.Equal(t, "boat", val)
	})
}
