package fortniteapi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBanners_All(t *testing.T) {
	t.Parallel()

	_, err := client.Banners.All(context.Background(), nil)
	require.NoError(t, err)
}

func TestBanners_Colors(t *testing.T) {
	t.Parallel()

	_, err := client.Banners.Colors(context.Background(), nil)
	require.NoError(t, err)
}
