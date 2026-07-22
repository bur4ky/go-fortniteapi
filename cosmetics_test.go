package fortniteapi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/bur4ky/go-fortniteapi"
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

func TestCosmetics_AllBRCosmetics(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.AllBRCosmetics(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestCosmetics_AllTracks(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.AllTracks(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestCosmetics_AllInstruments(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.AllInstruments(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestCosmetics_AllCars(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.AllCars(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestCosmetics_AllLego(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.AllLego(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestCosmetics_AllLegoKits(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.AllLegoKits(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, resp)
}

func TestCosmetics_AllBeans(t *testing.T) {
	t.Parallel()

	resp, err := client.Cosmetics.AllBeans(context.Background(), nil)
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
