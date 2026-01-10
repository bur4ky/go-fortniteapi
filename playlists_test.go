package fortniteapi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

const playlistID = "Playlist_DefaultSolo"

func TestPlaylists_All(t *testing.T) {
	t.Parallel()

	_, err := client.Playlists.All(context.Background(), nil)
	require.NoError(t, err)
}

func TestPlaylists_ByID(t *testing.T) {
	t.Parallel()

	resp, err := client.Playlists.ByID(context.Background(), playlistID, nil)
	require.NoError(t, err)
	require.Equal(t, playlistID, resp.ID)
}
