package fortniteapi

import "context"

type BRMapParams LanguageParams
type BRMapImages struct {
	Blank string `json:"blank"`
	Pois  string `json:"pois"`
}

type BRMapPOILocation struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type BRMapPOI struct {
	ID       string           `json:"id"`
	Name     string           `json:"name"`
	Location BRMapPOILocation `json:"location"`
}

type BRMapResponse struct {
	Images BRMapImages `json:"images"`
	POIs   []BRMapPOI  `json:"pois"`
}

type MapService struct {
	client *Client
}

func (s *MapService) BR(ctx context.Context, params *BRMapParams) (*BRMapResponse, error) {
	return getJSON[BRMapResponse](ctx, s.client, "/v1/map", params)
}
