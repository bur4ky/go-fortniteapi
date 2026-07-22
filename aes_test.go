package fortniteapi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAES_Key(t *testing.T) {
	t.Parallel()

	_, err := client.AES.Key(context.Background(), nil)
	require.NoError(t, err)
}
