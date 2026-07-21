package fortniteapi

import "context"

type BannersAllParams LanguageParams

type BannerImages struct {
	SmallIcon string `json:"smallIcon"`
	Icon      string `json:"icon"`
}

type BannerRarity BRCosmeticRarity
type BannerSeries BRCosmeticSeries
type BannerSet BRCosmeticSet
type BannerIntroduction BRCosmeticIntroduction

type Banner struct {
	ID              string              `json:"id"`
	DevName         string              `json:"devName"`
	Name            string              `json:"name"`
	Description     string              `json:"description"`
	Category        string              `json:"category"`
	FullUsageRights bool                `json:"fullUsageRights"`
	Rarity          *BannerRarity       `json:"rarity,omitempty"`
	Series          *BannerSeries       `json:"series,omitempty"`
	Set             *BannerSet          `json:"set,omitempty"`
	Introduction    *BannerIntroduction `json:"introduction,omitempty"`
	Images          BannerImages        `json:"images"`
}

type BannersAllResponse []Banner

type BannerColors struct {
	ID               string `json:"id"`
	Color            string `json:"color"`
	Category         string `json:"category"`
	SubCategoryGroup int    `json:"subCategoryGroup"`
}

type BannerColorsResponse []BannerColors

type BannersService struct {
	client *Client
}

func (s *BannersService) All(ctx context.Context, params *BannersAllParams) (BannersAllResponse, error) {
	return getJSON[BannersAllResponse](ctx, s.client, "/v1/banners", params)
}

func (s *BannersService) Colors(ctx context.Context, params *ResponseFlagsParams) (BannerColorsResponse, error) {
	return getJSON[BannerColorsResponse](ctx, s.client, "/v1/banners/colors", params)
}
