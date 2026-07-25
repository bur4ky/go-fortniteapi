package fortniteapi

import "context"

type BannersAllParams LanguageParams

type BannerImages struct {
	SmallIcon string `json:"smallIcon"`
	Icon      string `json:"icon"`
}

type Banner struct {
	ID              string                `json:"id"`
	DevName         string                `json:"devName"`
	Name            string                `json:"name"`
	Description     string                `json:"description"`
	Category        string                `json:"category"`
	FullUsageRights bool                  `json:"fullUsageRights"`
	Rarity          *CosmeticRarity       `json:"rarity,omitempty"`
	Series          *CosmeticSeries       `json:"series,omitempty"`
	Set             *CosmeticSet          `json:"set,omitempty"`
	Introduction    *CosmeticIntroduction `json:"introduction,omitempty"`
	Images          BannerImages          `json:"images"`
}

type BannerColors struct {
	ID               string `json:"id"`
	Color            string `json:"color"`
	Category         string `json:"category"`
	SubCategoryGroup int    `json:"subCategoryGroup"`
}

type BannersService struct {
	client *Client
}

func (s *BannersService) All(ctx context.Context, params *BannersAllParams) ([]Banner, error) {
	return getJSON[[]Banner](ctx, s.client, "/v1/banners", params)
}

func (s *BannersService) Colors(ctx context.Context, params *ResponseFlagsParams) ([]BannerColors, error) {
	return getJSON[[]BannerColors](ctx, s.client, "/v1/banners/colors", params)
}
