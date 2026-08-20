package fortniteapi

import (
	"context"
	"time"
)

type ShopParams LanguageParams

type ShopEntryBundle struct {
	Name  string `json:"name"`
	Info  string `json:"info"`
	Image string `json:"image"`
}

type ShopEntryBanner struct {
	Value        string `json:"value"`
	Intensity    string `json:"intensity"`
	BackendValue string `json:"backendValue"`
}

type ShopEntryOfferTag struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type ShopEntryLayoutMetadata struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ShopEntryLayout struct {
	ID                   string                    `json:"id"`
	Name                 string                    `json:"name"`
	Category             string                    `json:"category,omitempty"`
	Index                int                       `json:"index"`
	Rank                 int                       `json:"rank"`
	ShowIneligibleOffers string                    `json:"showIneligibleOffers"`
	Background           string                    `json:"background,omitempty"`
	UseWidePreview       bool                      `json:"useWidePreview"`
	DisplayType          string                    `json:"displayType"`
	TextureMetadata      []ShopEntryLayoutMetadata `json:"textureMetadata,omitempty"`
	StringMetadata       []ShopEntryLayoutMetadata `json:"stringMetadata,omitempty"`
	TextMetadata         []ShopEntryLayoutMetadata `json:"textMetadata,omitempty"`
}

type ShopEntryColors struct {
	Color1              string `json:"color1,omitempty"`
	Color2              string `json:"color2,omitempty"`
	Color3              string `json:"color3,omitempty"`
	TextBackgroundColor string `json:"textBackgroundColor,omitempty"`
}

// TODO: 40 chars long, might rename

type ShopEntryNewDisplayAssetMaterialInstance struct {
	ID          string            `json:"id"`
	PrimaryMode string            `json:"primaryMode"`
	ProductTag  string            `json:"productTag"`
	Images      map[string]string `json:"images,omitempty"`
	Colors      map[string]string `json:"colors,omitempty"`
	Scalings    map[string]string `json:"scalings,omitempty"`
	Flags       map[string]bool   `json:"flags,omitempty"`
}

type ShopEntryNewDisplayAssetRenderImage struct {
	ProductTag string `json:"productTag"`
	FileName   string `json:"fileName"`
	Image      string `json:"image"`
}

type ShopEntryNewDisplayAsset struct {
	ID                string                                     `json:"id"`
	CosmeticID        string                                     `json:"cosmeticId,omitempty"`
	MaterialInstances []ShopEntryNewDisplayAssetMaterialInstance `json:"materialInstances"`
	RenderImages      []ShopEntryNewDisplayAssetRenderImage      `json:"renderImages"`
}

type ShopEntry struct {
	RegularPrice           int                       `json:"regularPrice"`
	FinalPrice             int                       `json:"finalPrice"`
	DevName                string                    `json:"devName"`
	OfferID                string                    `json:"offerId"`
	InDate                 time.Time                 `json:"inDate"`
	OutDate                time.Time                 `json:"outDate"`
	Bundle                 *ShopEntryBundle          `json:"bundle,omitempty"`
	Banner                 *ShopEntryBanner          `json:"banner,omitempty"`
	OfferTag               *ShopEntryOfferTag        `json:"offerTag,omitempty"`
	Giftable               bool                      `json:"giftable"`
	Refundable             bool                      `json:"refundable"`
	SortPriority           int                       `json:"sortPriority"`
	LayoutID               string                    `json:"layoutId"`
	Layout                 *ShopEntryLayout          `json:"layout,omitempty"`
	Colors                 ShopEntryColors           `json:"colors"`
	TileBackgroundMaterial string                    `json:"tileBackgroundMaterial,omitempty"`
	TileSize               string                    `json:"tileSize"`
	DisplayAssetPath       string                    `json:"displayAssetPath,omitempty"`
	NewDisplayAssetPath    string                    `json:"newDisplayAssetPath"`
	NewDisplayAsset        *ShopEntryNewDisplayAsset `json:"newDisplayAsset,omitempty"`
	BRItems                []BRCosmetic              `json:"brItems,omitempty"`
	Tracks                 []Track                   `json:"tracks,omitempty"`
	Instruments            []Instrument              `json:"instruments,omitempty"`
	Cars                   []Car                     `json:"cars,omitempty"`
	LegoKits               []LegoKit                 `json:"legoKits,omitempty"`
}

type ShopResponse struct {
	Hash      string      `json:"hash"`
	Date      time.Time   `json:"date"`
	VBuckIcon string      `json:"vbuckIcon"`
	Entries   []ShopEntry `json:"entries"`
}

type ShopService struct {
	client *Client
}

func (s *ShopService) Get(ctx context.Context, params *ShopParams) (*ShopResponse, error) {
	return s.client.get[*ShopResponse](ctx, "/v2/shop", params)
}
