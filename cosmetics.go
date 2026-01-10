package fortniteapi

import (
	"context"
	"net/http"
	"time"
)

// SearchMatchMethod sets the method used to match search terms.
//
// Default: SearchMatchMethodFull
type SearchMatchMethod string

const (
	SearchMatchMethodFull     SearchMatchMethod = "full"
	SearchMatchMethodContains SearchMatchMethod = "contains"
	SearchMatchMethodStarts   SearchMatchMethod = "starts"
	SearchMatchMethodEnds     SearchMatchMethod = "ends"
)

type BRCosmeticType struct {
	Value        string `json:"value"`
	DisplayValue string `json:"displayValue"`
	BackendValue string `json:"backendValue"`
}

type BRCosmeticRarity struct {
	Value        string `json:"value"`
	DisplayValue string `json:"displayValue"`
	BackendValue string `json:"backendValue"`
}

type BRCosmeticSeries struct {
	Value        string   `json:"value"`
	Image        string   `json:"image"`
	Colors       []string `json:"colors"`
	BackendValue string   `json:"backendValue"`
}

type BRCosmeticSet struct {
	Value        string `json:"value"`
	Text         string `json:"text"`
	BackendValue string `json:"backendValue"`
}

type BRCosmeticIntroduction struct {
	Chapter      string `json:"chapter"`
	Season       string `json:"season"`
	Text         string `json:"text"`
	BackendValue int    `json:"backendValue"`
}

type BRCosmeticLegoImages struct {
	Small string `json:"small,omitempty"`
	Large string `json:"large,omitempty"`
	Wide  string `json:"wide,omitempty"`
}

type BRCosmeticBeanImages struct {
	Small string `json:"small,omitempty"`
	Large string `json:"large,omitempty"`
}

type BRCosmeticImages struct {
	SmallIcon string               `json:"smallIcon,omitempty"`
	Icon      string               `json:"icon,omitempty"`
	Featured  string               `json:"featured,omitempty"`
	Lego      BRCosmeticLegoImages `json:"lego,omitzero"`
	Bean      BRCosmeticBeanImages `json:"bean,omitzero"`
	Other     map[string]string    `json:"Other,omitzero"`
}

type BRCosmeticVariantOption struct {
	Tag                string `json:"tag"`
	Name               string `json:"name"`
	UnlockRequirements string `json:"unlockRequirements"`
	Image              string `json:"image"`
}

type BRCosmeticItemVariant struct {
	Channel string                    `json:"channel"`
	Type    string                    `json:"type"`
	Options []BRCosmeticVariantOption `json:"options"`
}

type BRCosmetic struct {
	ID                     string                  `json:"id"`
	Name                   string                  `json:"name"`
	Description            string                  `json:"description"`
	ExclusiveDescription   string                  `json:"exclusiveDescription,omitempty"`
	UnlockRequirements     string                  `json:"unlockRequirements,omitempty"`
	CustomExclusiveCallout string                  `json:"customExclusiveCallout,omitempty"`
	Type                   BRCosmeticType          `json:"type"`
	Rarity                 BRCosmeticRarity        `json:"rarity"`
	Series                 BRCosmeticSeries        `json:"series,omitzero"`
	Set                    BRCosmeticSet           `json:"set,omitzero"`
	Introduction           BRCosmeticIntroduction  `json:"introduction"`
	Images                 BRCosmeticImages        `json:"images"`
	Variants               []BRCosmeticItemVariant `json:"variants,omitzero"`
	BuiltInEmoteIDs        []string                `json:"builtInEmoteIds,omitzero"`
	SearchTags             []string                `json:"searchTags,omitempty"`
	GameplayTags           []string                `json:"gameplayTags,omitempty"`
	MetaTags               []string                `json:"metaTags,omitempty"`
	ShowcaseVideo          string                  `json:"showcaseVideo"`
	DynamicPakID           string                  `json:"dynamicPakId,omitempty"`
	ItemPreviewHeroPath    string                  `json:"itemPreviewHeroPath,omitempty"`
	DisplayAssetPath       string                  `json:"displayAssetPath,omitempty"`
	DefinitionPath         string                  `json:"definitionPath,omitempty"`
	Path                   string                  `json:"path,omitempty"`
	Added                  time.Time               `json:"added"`
	ShopHistory            []string                `json:"shopHistory,omitempty"`
}

type TrackDifficulty struct {
	Vocals       int `json:"vocals"`
	Guitar       int `json:"guitar"`
	Bass         int `json:"bass"`
	PlasticBass  int `json:"plasticBass"`
	Drums        int `json:"drums"`
	PlasticDrums int `json:"plasticDrums"`
}

type Track struct {
	ID           string          `json:"id"`
	DevName      string          `json:"devName"`
	Title        string          `json:"title"`
	Artist       string          `json:"artist"`
	Album        string          `json:"album"`
	ReleaseYear  int             `json:"releaseYear"`
	BPM          int             `json:"bpm"`
	Duration     int             `json:"duration"`
	Difficulty   TrackDifficulty `json:"difficulty"`
	GameplayTags []string        `json:"gameplayTags,omitempty"`
	Genres       []string        `json:"genres"`
	AlbumArt     string          `json:"albumArt"`
	Added        time.Time       `json:"added,omitempty"`
	ShopHistory  []string        `json:"shopHistory,omitempty"`
}

type InstrumentImages struct {
	Small string `json:"small,omitempty"`
	Large string `json:"large,omitempty"`
}

type Instrument struct {
	ID            string           `json:"id"`
	Name          string           `json:"name"`
	Description   string           `json:"description"`
	Type          BRCosmeticType   `json:"type"`
	Rarity        BRCosmeticRarity `json:"rarity"`
	Images        InstrumentImages `json:"images"`
	Series        BRCosmeticSeries `json:"series,omitzero"`
	GameplayTags  []string         `json:"gameplayTags,omitempty"`
	Path          string           `json:"path"`
	ShowcaseVideo string           `json:"showcaseVideo"`
	Added         time.Time        `json:"added"`
	ShopHistory   []string         `json:"shopHistory,omitempty"`
}

type CarImages struct {
	Small string `json:"small,omitempty"`
	Large string `json:"large,omitempty"`
}

type Car struct {
	ID            string           `json:"id"`
	VehicleID     string           `json:"vehicleId"`
	Name          string           `json:"name"`
	Description   string           `json:"description"`
	Type          BRCosmeticType   `json:"type"`
	Rarity        BRCosmeticRarity `json:"rarity"`
	Images        CarImages        `json:"images"`
	Series        BRCosmeticSeries `json:"series,omitzero"`
	GameplayTags  []string         `json:"gameplayTags,omitempty"`
	Path          string           `json:"path,omitempty"`
	ShowcaseVideo string           `json:"showcaseVideo"`
	Added         time.Time        `json:"added"`
	ShopHistory   []string         `json:"shopHistory,omitempty"`
}

type LegoImages struct {
	Small string `json:"small,omitempty"`
	Large string `json:"large,omitempty"`
	Wide  string `json:"wide,omitempty"`
}

type Lego struct {
	ID               string     `json:"id"`
	Name             string     `json:"name"`
	SoundLibraryTags []string   `json:"soundLibraryTags"`
	Images           LegoImages `json:"images"`
	Path             string     `json:"path"`
	Added            time.Time  `json:"added"`
}

type LegoKitsImages struct {
	Small string `json:"small,omitempty"`
	Large string `json:"large,omitempty"`
	Wide  string `json:"wide,omitempty"`
}

type LegoKit struct {
	ID           string           `json:"id"`
	Name         string           `json:"name"`
	Type         BRCosmeticType   `json:"type"`
	Series       BRCosmeticSeries `json:"series,omitzero"`
	GameplayTags []string         `json:"gameplayTags,omitempty"`
	Images       LegoKitsImages   `json:"images"`
	Path         string           `json:"path,omitempty"`
	Added        time.Time        `json:"added"`
	ShopHistory  []string         `json:"shopHistory,omitempty"`
}

type BeanImages struct {
	Small string `json:"small,omitempty"`
	Large string `json:"large,omitempty"`
}

type Bean struct {
	ID           string     `json:"id"`
	CosmeticID   string     `json:"cosmeticId"`
	Name         string     `json:"name"`
	Gender       string     `json:"gender"`
	GameplayTags []string   `json:"gameplayTags"`
	Images       BeanImages `json:"images"`
	Path         string     `json:"path"`
	Added        time.Time  `json:"added"`
}

type AllCosmeticsParams LanguageParams
type AllCosmeticsResponse struct {
	BR          []BRCosmetic `json:"br"`
	Tracks      []Track      `json:"tracks"`
	Instruments []Instrument `json:"instruments"`
	Cars        []Car        `json:"cars"`
	Lego        []Lego       `json:"lego"`
	LegoKits    []LegoKit    `json:"legoKits"`
	Beans       []Bean       `json:"beans"`
}

type NewCosmeticsParams LanguageParams
type NewCosmeticsHashes struct {
	All         string `json:"all"`
	BR          string `json:"br"`
	Tracks      string `json:"tracks"`
	Instruments string `json:"instruments"`
	Cars        string `json:"cars"`
	Lego        string `json:"lego"`
	LegoKits    string `json:"legoKits"`
	Beans       string `json:"beans"`
}

type NewCosmeticsLastAdditions struct {
	All         string `json:"all"`
	BR          string `json:"br"`
	Tracks      string `json:"tracks"`
	Instruments string `json:"instruments"`
	Cars        string `json:"cars"`
	Lego        string `json:"lego"`
	LegoKits    string `json:"legoKits"`
	Beans       string `json:"beans"`
}

type NewCosmeticsResponse struct {
	Date          time.Time                 `json:"date"`
	Build         string                    `json:"build"`
	PreviousBuild string                    `json:"previousBuild"`
	Hashes        NewCosmeticsHashes        `json:"hashes"`
	LastAdditions NewCosmeticsLastAdditions `json:"lastAdditions"`
	Items         AllCosmeticsResponse      `json:"items"`
}

type BRCosmeticsAllParams LanguageParams
type BRCosmeticsAllResponse []BRCosmetic

type TrackCosmeticsAllParams ResponseFlagsParams
type TrackCosmeticsAllResponse []Track

type InstrumentCosmeticsAllParams LanguageParams
type InstrumentCosmeticsAllResponse []Instrument

type CarCosmeticsAllParams LanguageParams
type CarCosmeticsAllResponse []Car

type LegoCosmeticsAllParams ResponseFlagsParams
type LegoCosmeticsAllResponse []Lego

type LegoKitCosmeticsAllParams LanguageParams
type LegoKitCosmeticsAllResponse []LegoKit

type BeanCosmeticsAllParams LanguageParams
type BeanCosmeticsAllResponse []Bean

type BRCosmeticByIDParams LanguageParams
type BRCosmeticByIDResponse BRCosmetic

type SearchBRCosmeticParams struct {
	Language            Language          `url:"language,omitempty"`
	SearchLanguage      Language          `url:"searchLanguage,omitempty"`
	MatchMethod         SearchMatchMethod `url:"matchMethod,omitempty"`
	ID                  string            `url:"id,omitempty"`
	Name                string            `url:"name,omitempty"`
	Description         string            `url:"description,omitempty"`
	Type                string            `url:"type,omitempty"`
	DisplayType         string            `url:"displayType,omitempty"`
	BackendType         string            `url:"backendType,omitempty"`
	Rarity              string            `url:"rarity,omitempty"`
	DisplayRarity       string            `url:"displayRarity,omitempty"`
	BackendRarity       string            `url:"backendRarity,omitempty"`
	HasSeries           bool              `url:"hasSeries,omitempty"`
	Series              string            `url:"series,omitempty"`
	BackendSeries       string            `url:"backendSeries,omitempty"`
	HasSet              bool              `url:"hasSet,omitempty"`
	Set                 string            `url:"set,omitempty"`
	SetText             string            `url:"setText,omitempty"`
	BackendSet          string            `url:"backendSet,omitempty"`
	HasIntroduction     bool              `url:"hasIntroduction,omitempty"`
	BackendIntroduction int               `url:"backendIntroduction,omitempty"`
	IntroductionChapter string            `url:"introductionChapter,omitempty"`
	IntroductionSeason  string            `url:"introductionSeason,omitempty"`
	HasFeaturedImage    bool              `url:"hasFeaturedImage,omitempty"`
	HasVariants         bool              `url:"hasVariants,omitempty"`
	HasGameplayTags     bool              `url:"hasGameplayTags,omitempty"`
	GameplayTag         string            `url:"gameplayTag,omitempty"`
	HasMetaTags         bool              `url:"hasMetaTags,omitempty"`
	MetaTag             string            `url:"metaTag,omitempty"`
	HasDynamicPakID     bool              `url:"hasDynamicPakId,omitempty"`
	DynamicPakID        string            `url:"dynamicPakId,omitempty"`
	Added               int64             `url:"added,omitempty"`
	AddedSince          int64             `url:"addedSince,omitempty"`
	UnseenFor           int               `url:"unseenFor,omitempty"`
	LastAppearance      int64             `url:"lastAppearance,omitempty"`
	ResponseFlags       ResponseFlag      `url:"responseFlags,omitempty"`
}

type SearchBRCosmeticResponse BRCosmetic

type SearchBRCosmeticsParams SearchBRCosmeticParams
type SearchBRCosmeticsResponse []SearchBRCosmeticResponse

type BRCosmeticsSearchByIDsParams LanguageParams
type BRCosmeticsSearchByIDsResponse []BRCosmetic

type CosmeticsService struct {
	client *Client
}

func (s *CosmeticsService) All(ctx context.Context, params *AllCosmeticsParams) (*AllCosmeticsResponse, error) {
	return getJSON[AllCosmeticsResponse](ctx, s.client, "/v2/cosmetics", params)
}

func (s *CosmeticsService) New(ctx context.Context, params *NewCosmeticsParams) (*NewCosmeticsResponse, error) {
	return getJSON[NewCosmeticsResponse](ctx, s.client, "/v2/cosmetics/new", params)
}

func (s *CosmeticsService) BRCosmeticsAll(ctx context.Context, params *BRCosmeticsAllParams) (BRCosmeticsAllResponse, error) {
	return getJSONSlice[BRCosmeticsAllResponse](ctx, s.client, "/v2/cosmetics/br", params)
}

func (s *CosmeticsService) TrackCosmeticsAll(ctx context.Context, params *TrackCosmeticsAllParams) (TrackCosmeticsAllResponse, error) {
	return getJSONSlice[TrackCosmeticsAllResponse](ctx, s.client, "/v2/cosmetics/tracks", params)
}

func (s *CosmeticsService) InstrumentCosmeticsAll(ctx context.Context, params *InstrumentCosmeticsAllParams) (InstrumentCosmeticsAllResponse, error) {
	return getJSONSlice[InstrumentCosmeticsAllResponse](ctx, s.client, "/v2/cosmetics/instruments", params)
}

func (s *CosmeticsService) CarCosmeticsAll(ctx context.Context, params *CarCosmeticsAllParams) (CarCosmeticsAllResponse, error) {
	return getJSONSlice[CarCosmeticsAllResponse](ctx, s.client, "/v2/cosmetics/cars", params)
}

func (s *CosmeticsService) LegoCosmeticsAll(ctx context.Context, params *LegoCosmeticsAllParams) (LegoCosmeticsAllResponse, error) {
	return getJSONSlice[LegoCosmeticsAllResponse](ctx, s.client, "/v2/cosmetics/lego", params)
}

func (s *CosmeticsService) LegoKitCosmeticsAll(ctx context.Context, params *LegoKitCosmeticsAllParams) (LegoKitCosmeticsAllResponse, error) {
	return getJSONSlice[LegoKitCosmeticsAllResponse](ctx, s.client, "/v2/cosmetics/lego/kits", params)
}

func (s *CosmeticsService) BeanCosmeticsAll(ctx context.Context, params *BeanCosmeticsAllParams) (BeanCosmeticsAllResponse, error) {
	return getJSONSlice[BeanCosmeticsAllResponse](ctx, s.client, "/v2/cosmetics/beans", params)
}

func (s *CosmeticsService) BRCosmeticByID(ctx context.Context, id string, params *BRCosmeticByIDParams) (*BRCosmeticByIDResponse, error) {
	if id == "" {
		return nil, emptyParamErr("id")
	}

	return getJSON[BRCosmeticByIDResponse](ctx, s.client, "/v2/cosmetics/br/"+id, params)
}

func (s *CosmeticsService) SearchBRCosmetic(ctx context.Context, params *SearchBRCosmeticParams) (*SearchBRCosmeticResponse, error) {
	if params == nil {
		params = &SearchBRCosmeticParams{}
	}

	if params.SearchLanguage == "" {
		params.SearchLanguage = s.client.language
	}

	return getJSON[SearchBRCosmeticResponse](ctx, s.client, "/v2/cosmetics/br/search", params)
}

func (s *CosmeticsService) SearchBRCosmetics(ctx context.Context, params *SearchBRCosmeticsParams) (SearchBRCosmeticsResponse, error) {
	if params == nil {
		params = &SearchBRCosmeticsParams{}
	}

	if params.SearchLanguage == "" {
		params.SearchLanguage = s.client.language
	}

	return getJSONSlice[SearchBRCosmeticsResponse](ctx, s.client, "/v2/cosmetics/br/search/all", params)
}

func (s *CosmeticsService) SearchBRCosmeticsByIDs(ctx context.Context, ids []string, params *BRCosmeticsSearchByIDsParams) (BRCosmeticsSearchByIDsResponse, error) {
	if len(ids) == 0 {
		return nil, emptyParamErr("ids")
	}

	var out BRCosmeticsSearchByIDsResponse
	err := s.client.fetch(ctx, http.MethodPost, "/v2/cosmetics/br/search/ids", params, ids, &out)
	return out, err
}
