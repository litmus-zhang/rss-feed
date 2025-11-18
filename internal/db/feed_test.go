package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFeedOperations(t *testing.T) {
	t.Cleanup(func() {
		clearTestTables(testDB) // Clean after each test
	})

	// Test cases
	t.Run("Create feed", func(t *testing.T) {
		// Test implementation
		// Create a new feed
		f := CreateTestFeed(t)

		require.NotEmpty(t, f)
		require.NotEmpty(t, f.FeedName)
		require.NotEmpty(t, f.Url)
		require.NotEmpty(t, f.Description)
	})
	t.Run("Get one feed", func(t *testing.T) {
		f1 := CreateTestFeed(t)

		f2, err := testQueries.GetOneFeedById(context.Background(), f1.FeedID)

		require.NoError(t, err)

		require.Equal(t, f1, f2)
		require.Equal(t, f2.FeedName, f1.FeedName)
		require.NotEmpty(t, f2.Url, f1.Url)
		require.NotEmpty(t, f2.Description, f1.Description)
	})
	t.Run("Update feed", func(t *testing.T) {
		f1 := CreateTestFeed(t)

		arg := UpdateFeedParams{
			FeedID:      f1.FeedID,
			FeedName:    f1.FeedName + "Updated",
			Url:         f1.Url + "updated",
			Description: f1.Description,
		}

		f2, err := testQueries.UpdateFeed(context.Background(), arg)

		require.NoError(t, err)

		require.NotEqual(t, f1, f2)
		require.Equal(t, f2.FeedName, arg.FeedName)
		require.Equal(t, f2.Description, arg.Description)
	})
	t.Run("Get all feeds", func(t *testing.T) {
		// Test implementation
		// Create a new feed
		for i := 0; i < 5; i++ {
			CreateTestFeed(t)
		}
		all, err := testQueries.GetAllFeeds(context.Background(), GetAllFeedsParams{
			Limit:  5,
			Offset: 0,
		})
		require.NoError(t, err)
		require.Equal(t, len(all), 5)

	})
	t.Run("Delete a feed", func(t *testing.T) {
		// Test implementation
		// Create a new feed
		f := CreateTestFeed(t)

		err := testQueries.DeleteFeed(context.Background(), f.FeedID)
		require.NoError(t, err)

		// Check if the feed is deleted
		_, err = testQueries.GetOneFeedById(context.Background(), f.FeedID)
		require.Error(t, err)

	})
}
