package fortniteapi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap_BR(t *testing.T) {
	t.Parallel()

	_, err := client.Map.BR(context.Background(), nil)
	require.NoError(t, err)
}
