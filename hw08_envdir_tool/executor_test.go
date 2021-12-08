package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	t.Run("Test empty command", func(t *testing.T) {
		require.Equal(t, 1, RunCmd([]string{}, Environment{}))
	})

	t.Run("Test error code", func(t *testing.T) {
		require.Equal(t, 127, RunCmd([]string{"sh", "NilFile"}, Environment{}))
	})

	t.Run("Test set env", func(t *testing.T) {
		require.Equal(t, 0, RunCmd([]string{"sh"}, Environment{"Test": EnvValue{Value: "value"}}))

		env, ok := os.LookupEnv("Test")
		require.True(t, ok)
		require.Equal(t, "value", env)
	})

	t.Run("Test unset env", func(t *testing.T) {
		require.Equal(t, 0, RunCmd([]string{"sh"}, Environment{"Test": EnvValue{Value: ""}}))

		env, ok := os.LookupEnv("Test")
		require.False(t, ok)
		require.Empty(t, env)
	})
}
