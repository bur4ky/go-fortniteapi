package fortniteapi_test

import (
	"context"
	"testing"

	"github.com/bur4ky/go-fortniteapi"
	"github.com/stretchr/testify/require"
)

const (
	cosmeticName = "Peely"
	cosmeticID1  = "CID_349_Athena_Commando_M_Banana"
	cosmeticID2  = "CID_049_Athena_Commando_M_HolidayGingerbread"
)

func TestCosmetics_All(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.All(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestCosmetics_New(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.New(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestCosmetics_BRCosmeticsAll(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.BRCosmeticsAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestCosmetics_TrackCosmeticsAll(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.TrackCosmeticsAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestCosmetics_InstrumentCosmeticsAll(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.InstrumentCosmeticsAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestCosmetics_CarCosmeticsAll(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.CarCosmeticsAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestCosmetics_LegoCosmeticsAll(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.LegoCosmeticsAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestCosmetics_LegoKitCosmeticsAll(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.LegoKitCosmeticsAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestCosmetics_BeanCosmeticsAll(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.BeanCosmeticsAll(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestCosmetics_BRCosmeticByID(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.BRCosmeticByID(context.Background(), cosmeticID1, nil)
	require.NoError(t, err)
	require.Equal(t, cosmeticID1, resp.ID)
}

func TestCosmetics_SearchBRCosmetic(t *testing.T) {
	t.Parallel()

	params := &fortniteapi.SearchBRCosmeticParams{Name: cosmeticName}
	resp, err := client.Cosmetics.SearchBRCosmetic(context.Background(), params)
	require.NoError(t, err)
	require.Equal(t, cosmeticName, resp.Name)
}

func TestCosmetics_SearchBRCosmetics(t *testing.T) {
	t.Parallel()

	params := &fortniteapi.SearchBRCosmeticsParams{Name: cosmeticName}
	resp, err := client.Cosmetics.SearchBRCosmetics(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, resp)

	for _, c := range resp {
		require.Contains(t, c.Name, cosmeticName)
	}
}

func TestCosmetics_SearchBRCosmeticsByIDs(t *testing.T) {
	t.Parallel()

	ids := []string{cosmeticID1, cosmeticID2}
	resp, err := client.Cosmetics.SearchBRCosmeticsByIDs(context.Background(), ids, nil)

	require.NoError(t, err)
	require.Len(t, resp, 2)

	for _, c := range resp {
		require.Contains(t, ids, c.ID)
	}
}
