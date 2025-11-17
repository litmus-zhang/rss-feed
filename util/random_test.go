package util

import (
	"testing"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	randStr := RandomString(10)
	require.Len(t, randStr, 10)
}

func TestRandomInt(t *testing.T) {
	randInt := RandomInt(1, 100)
	require.GreaterOrEqual(t, randInt, 1)
	require.LessOrEqual(t, randInt, 100)
}
