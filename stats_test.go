package fortniteapi_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testStatsName = "BurakYhs"
	testStatsID   = "05006cb489c347beaad83551a1b9544e"
)

func TestStats_BRByName(t *testing.T) {
	t.Parallel()
	requireAPIKey(t)

	_, err := client.Stats.BRByName(context.Background(), testStatsName, nil)
	require.NoError(t, err)
}

func TestStats_BRByID(t *testing.T) {
	t.Parallel()
	requireAPIKey(t)

	resp, err := client.Stats.BRByID(context.Background(), testStatsID, nil)
	require.NoError(t, err)
	require.Equal(t, testStatsID, resp.Account.ID)
}

func requireAPIKey(t *testing.T) {
	t.Helper()

	if os.Getenv("FORTNITE_API_KEY") == "" {
		t.Skip("FORTNITE_API_KEY is not set")
	}
}
