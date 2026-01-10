package fortniteapi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNews_All(t *testing.T) {
	t.Parallel()

	_, err := client.News.All(context.Background(), nil)
	require.NoError(t, err)
}

func TestNews_BR(t *testing.T) {
	t.Parallel()

	_, err := client.News.BR(context.Background(), nil)
	require.NoError(t, err)
}

func TestNews_STW(t *testing.T) {
	t.Parallel()

	_, err := client.News.STW(context.Background(), nil)
	require.NoError(t, err)
}

func TestNews_Creative(t *testing.T) {
	t.Parallel()
	t.Skip("Creative news are not available anymore")

	_, err := client.News.Creative(context.Background(), nil)
	require.NoError(t, err)
}
