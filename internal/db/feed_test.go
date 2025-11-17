package db

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateFeed(t *testing.T) {
	// Create a new feed
	f := CreateTestFeed(t)

	require.NotEmpty(t, f)
}
