package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/litmus-zhang/rss-feed/internal/config"
	"github.com/litmus-zhang/rss-feed/util"
	"github.com/stretchr/testify/require"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("cannot setup a new config:", err)
	}

	testDB, err = sql.Open(cfg.DbDriver, cfg.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
	testDB.Close()
}

func clearTestTables(db *sql.DB) {
	tables := []string{"feeds"} // Add your table names
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", table))
		if err != nil {
			log.Fatalf("Failed to truncate table %s: %v", table, err)
		}
	}
}

func CreateTestFeed(t *testing.T) Feed {

	args := CreateFeedParams{
		FeedName:    util.RandomString(6),
		Url:         fmt.Sprintf("http://test.com/feed-%s", util.RandomString(3)),
		Description: sql.NullString{Valid: true, String: util.RandomString(6) + " " + util.RandomString(6)},
	}

	feed, err := testQueries.CreateFeed(context.Background(), args)
	require.NoError(t, err)

	require.NotEmpty(t, feed)
	require.Equal(t, args.FeedName, feed.FeedName)
	require.Equal(t, args.Url, feed.Url)
	require.Equal(t, args.Description, feed.Description)
	require.NotZero(t, feed.FeedID)
	return feed
}
