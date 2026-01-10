package fortniteapi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAES_Keys(t *testing.T) {
	t.Parallel()

	_, err := client.AES.Keys(context.Background(), nil)
	require.NoError(t, err)
}
