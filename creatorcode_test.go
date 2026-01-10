package fortniteapi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

const creatorCode = "ninja"

func TestCreatorCode_Get(t *testing.T) {
	t.Parallel()

	resp, err := client.CreatorCode.Get(context.Background(), creatorCode, nil)
	require.NoError(t, err)
	require.Equal(t, creatorCode, resp.Code)
}
