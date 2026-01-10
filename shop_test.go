package fortniteapi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShop_Get(t *testing.T) {
	t.Parallel()

	_, err := client.Shop.Get(context.Background(), nil)
	require.NoError(t, err)
}
