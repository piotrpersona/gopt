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

		opt := gopt.NewSome("car")
		val, err := opt.Get()
		require.NoError(t, err)
		require.Equal(t, "car", val)
	})

	t.Run("Test none", func(t *testing.T) {
		t.Parallel()

		opt := gopt.NewNone()
		val, err := opt.Get()
		require.Error(t, err)
		require.ErrorIs(t, err, gopt.NoneErr)
	})
}
