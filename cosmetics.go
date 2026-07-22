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
	SmallIcon string                `json:"smallIcon,omitempty"`
	Icon      string                `json:"icon,omitempty"`
	Featured  string                `json:"featured,omitempty"`
	Lego      *BRCosmeticLegoImages `json:"lego,omitempty"`
	Bean      *BRCosmeticBeanImages `json:"bean,omitempty"`
	Other     map[string]string     `json:"other,omitempty"`
}

type BRCosmeticVariantOption struct {
	Tag                string `json:"tag"`
	Name               string `json:"name"`
	UnlockRequirements string `json:"unlockRequirements,omitempty"`
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
	Series                 *BRCosmeticSeries       `json:"series,omitempty"`
	Set                    *BRCosmeticSet          `json:"set,omitempty"`
	Introduction           BRCosmeticIntroduction  `json:"introduction"`
	Images                 BRCosmeticImages        `json:"images"`
	Variants               []BRCosmeticItemVariant `json:"variants,omitempty"`
	BuiltInEmoteIDs        []string                `json:"builtInEmoteIds,omitempty"`
	SearchTags             []string                `json:"searchTags,omitempty"`
	GameplayTags           []string                `json:"gameplayTags,omitempty"`
	MetaTags               []string                `json:"metaTags,omitempty"`
	ShowcaseVideo          string                  `json:"showcaseVideo,omitempty"`
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
	Album        string          `json:"album,omitempty"`
	ReleaseYear  int             `json:"releaseYear"`
	BPM          int             `json:"bpm"`
	Duration     int             `json:"duration"`
	Difficulty   TrackDifficulty `json:"difficulty"`
	GameplayTags []string        `json:"gameplayTags,omitempty"`
	Genres       []string        `json:"genres,omitempty"`
	AlbumArt     string          `json:"albumArt"`
	Added        time.Time       `json:"added"`
	ShopHistory  []string        `json:"shopHistory,omitempty"`
}

type InstrumentImages struct {
	Small string `json:"small,omitempty"`
	Large string `json:"large,omitempty"`
}

type Instrument struct {
	ID            string            `json:"id"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	Type          BRCosmeticType    `json:"type"`
	Rarity        BRCosmeticRarity  `json:"rarity"`
	Images        InstrumentImages  `json:"images"`
	Series        *BRCosmeticSeries `json:"series,omitempty"`
	GameplayTags  []string          `json:"gameplayTags,omitempty"`
	Path          string            `json:"path"`
	ShowcaseVideo string            `json:"showcaseVideo,omitempty"`
	Added         time.Time         `json:"added"`
	ShopHistory   []string          `json:"shopHistory,omitempty"`
}

type CarImages struct {
	Small string `json:"small,omitempty"`
	Large string `json:"large,omitempty"`
}

type Car struct {
	ID            string            `json:"id"`
	VehicleID     string            `json:"vehicleId"`
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	Type          BRCosmeticType    `json:"type"`
	Rarity        BRCosmeticRarity  `json:"rarity"`
	Images        CarImages         `json:"images"`
	Series        *BRCosmeticSeries `json:"series,omitempty"`
	GameplayTags  []string          `json:"gameplayTags,omitempty"`
	Path          string            `json:"path,omitempty"`
	ShowcaseVideo string            `json:"showcaseVideo,omitempty"`
	Added         time.Time         `json:"added"`
	ShopHistory   []string          `json:"shopHistory,omitempty"`
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
	ID           string            `json:"id"`
	Name         string            `json:"name"`
	Type         BRCosmeticType    `json:"type"`
	Series       *BRCosmeticSeries `json:"series,omitempty"`
	GameplayTags []string          `json:"gameplayTags,omitempty"`
	Images       LegoKitsImages    `json:"images"`
	Path         string            `json:"path,omitempty"`
	Added        time.Time         `json:"added"`
	ShopHistory  []string          `json:"shopHistory,omitempty"`
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

type AllBRCosmeticsParams LanguageParams
type AllTracksParams ResponseFlagsParams
type AllInstrumentsParams LanguageParams
type AllCarsParams LanguageParams
type AllLegoParams ResponseFlagsParams
type AllLegoKitsParams LanguageParams
type AllBeansParams LanguageParams
type BRCosmeticByIDParams LanguageParams

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
	ResponseFlags       ResponseFlags     `url:"responseFlags,omitempty"`
}

type SearchBRCosmeticsParams SearchBRCosmeticParams
type BRCosmeticsSearchByIDsParams LanguageParams

type CosmeticsService struct {
	client *Client
}

func (s *CosmeticsService) All(ctx context.Context, params *AllCosmeticsParams) (*AllCosmeticsResponse, error) {
	return getJSON[*AllCosmeticsResponse](ctx, s.client, "/v2/cosmetics", params)
}

func (s *CosmeticsService) New(ctx context.Context, params *NewCosmeticsParams) (*NewCosmeticsResponse, error) {
	return getJSON[*NewCosmeticsResponse](ctx, s.client, "/v2/cosmetics/new", params)
}

func (s *CosmeticsService) AllBRCosmetics(ctx context.Context, params *AllBRCosmeticsParams) ([]BRCosmetic, error) {
	return getJSON[[]BRCosmetic](ctx, s.client, "/v2/cosmetics/br", params)
}

func (s *CosmeticsService) AllTracks(ctx context.Context, params *AllTracksParams) ([]Track, error) {
	return getJSON[[]Track](ctx, s.client, "/v2/cosmetics/tracks", params)
}

func (s *CosmeticsService) AllInstruments(ctx context.Context, params *AllInstrumentsParams) ([]Instrument, error) {
	return getJSON[[]Instrument](ctx, s.client, "/v2/cosmetics/instruments", params)
}

func (s *CosmeticsService) AllCars(ctx context.Context, params *AllCarsParams) ([]Car, error) {
	return getJSON[[]Car](ctx, s.client, "/v2/cosmetics/cars", params)
}

func (s *CosmeticsService) AllLego(ctx context.Context, params *AllLegoParams) ([]Lego, error) {
	return getJSON[[]Lego](ctx, s.client, "/v2/cosmetics/lego", params)
}

func (s *CosmeticsService) AllLegoKits(ctx context.Context, params *AllLegoKitsParams) ([]LegoKit, error) {
	return getJSON[[]LegoKit](ctx, s.client, "/v2/cosmetics/lego/kits", params)
}

func (s *CosmeticsService) AllBeans(ctx context.Context, params *AllBeansParams) ([]Bean, error) {
	return getJSON[[]Bean](ctx, s.client, "/v2/cosmetics/beans", params)
}

func (s *CosmeticsService) BRCosmeticByID(ctx context.Context, id string, params *BRCosmeticByIDParams) (*BRCosmetic, error) {
	if id == "" {
		return nil, emptyParamErr("id")
	}

	return getJSON[*BRCosmetic](ctx, s.client, "/v2/cosmetics/br/"+id, params)
}

func (s *CosmeticsService) SearchBRCosmetic(ctx context.Context, params *SearchBRCosmeticParams) (*BRCosmetic, error) {
	if params.SearchLanguage == "" {
		if params == nil {
			params = &SearchBRCosmeticParams{}
		}

		params.SearchLanguage = s.client.language
	}

	return getJSON[*BRCosmetic](ctx, s.client, "/v2/cosmetics/br/search", params)
}

func (s *CosmeticsService) SearchBRCosmetics(ctx context.Context, params *SearchBRCosmeticsParams) ([]BRCosmetic, error) {
	if params.SearchLanguage == "" {
		if params == nil {
			params = &SearchBRCosmeticsParams{}
		}

		params.SearchLanguage = s.client.language
	}

	return getJSON[[]BRCosmetic](ctx, s.client, "/v2/cosmetics/br/search/all", params)
}

func (s *CosmeticsService) SearchBRCosmeticsByIDs(ctx context.Context, ids []string, params *BRCosmeticsSearchByIDsParams) ([]BRCosmetic, error) {
	if len(ids) == 0 {
		return nil, emptyParamErr("ids")
	}

	var out []BRCosmetic
	err := s.client.do(ctx, http.MethodPost, "/v2/cosmetics/br/search/ids", params, ids, &out)
	return out, err
}
