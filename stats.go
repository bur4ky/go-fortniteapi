package fortniteapi

import "context"

// StatsAccountType represents the type of account for fetching stats.
//
// Default: StatsAccountEpic
type StatsAccountType string

const (
	StatsAccountEpic StatsAccountType = "epic"
	StatsAccountPSN  StatsAccountType = "psn"
	StatsAccountXBL  StatsAccountType = "xbl"
)

// StatsTimeWindow represents the time window for fetching stats.
//
// Default: StatsTimeWindowLifetime
type StatsTimeWindow string

const (
	StatsTimeWindowSeason   StatsTimeWindow = "season"
	StatsTimeWindowLifetime StatsTimeWindow = "lifetime"
)

// StatsImage represents the input device type for fetching stats.
//
// Default: StatsImageNone
type StatsImage string

const (
	StatsImageAll           StatsImage = "all"
	StatsImageKeyboardMouse StatsImage = "keyboardMouse"
	StatsImageGamepad       StatsImage = "gamepad"
	StatsImageTouch         StatsImage = "touch"
	StatsImageNone          StatsImage = "none"
)

type BRStatsByNameParams struct {
	Name          string           `url:"name"`
	AccountType   StatsAccountType `url:"accountType,omitempty"`
	TimeWindow    StatsTimeWindow  `url:"timeWindow,omitempty"`
	Image         StatsImage       `url:"image,omitempty"`
	ResponseFlags ResponseFlag     `url:"responseFlags,omitempty"`
}

type BRStatsByIDParams struct {
	TimeWindow    StatsTimeWindow `url:"timeWindow,omitempty"`
	Image         StatsImage      `url:"image,omitempty"`
	ResponseFlags ResponseFlag    `url:"responseFlags,omitempty"`
}

type BRStatsAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type BRStatsBattlePass struct {
	Level    int `json:"level"`
	Progress int `json:"progress"`
}

type BRStatsData struct {
	Score           int     `json:"score"`
	ScorePerMin     float64 `json:"scorePerMin"`
	ScorePerMatch   float64 `json:"scorePerMatch"`
	Wins            int     `json:"wins"`
	Top3            int     `json:"top3"`
	Top5            int     `json:"top5"`
	Top6            int     `json:"top6"`
	Top10           int     `json:"top10"`
	Top12           int     `json:"top12"`
	Top25           int     `json:"top25"`
	Kills           int     `json:"kills"`
	KillsPerMin     float64 `json:"killsPerMin"`
	KillsPerMatch   float64 `json:"killsPerMatch"`
	Deaths          int     `json:"deaths"`
	KD              float64 `json:"kd"`
	Matches         int     `json:"matches"`
	WinRate         float64 `json:"winRate"`
	MinutesPlayed   int     `json:"minutesPlayed"`
	PlayersOutlived int     `json:"playersOutlived"`
	LastModified    string  `json:"lastModified"`
}

type BRStatsModes struct {
	Overall BRStatsData `json:"overall"`
	Solo    BRStatsData `json:"solo"`
	Duo     BRStatsData `json:"duo"`
	Trio    BRStatsData `json:"trio"`
	Squad   BRStatsData `json:"squad"`
	LTM     BRStatsData `json:"ltm"`
}

type BRStatsAll BRStatsModes
type BRStatsKeyboardMouse BRStatsModes
type BRStatsGamepad BRStatsModes
type BRStatsTouch BRStatsModes
type BRStatsStats struct {
	All           BRStatsAll           `json:"all"`
	KeyboardMouse BRStatsKeyboardMouse `json:"keyboardMouse"`
	Gamepad       BRStatsGamepad       `json:"gamepad"`
	Touch         BRStatsTouch         `json:"touch"`
}

type BRStatsResponse struct {
	Account    BRStatsAccount    `json:"account"`
	BattlePass BRStatsBattlePass `json:"battlePass"`
	Image      string            `json:"image"`
	Stats      BRStatsStats      `json:"stats"`
}

type StatsService struct {
	client *Client
}

func (s *StatsService) BRByName(ctx context.Context, name string, params *BRStatsByNameParams) (*BRStatsResponse, error) {
	if s.client.apiKey == "" {
		return nil, ErrNoAPIKey
	}

	if name == "" {
		return nil, emptyParamErr("name")
	}

	if params == nil {
		params = &BRStatsByNameParams{}
	}

	params.Name = name

	return getJSON[BRStatsResponse](ctx, s.client, "/v2/stats/br/v2", params)
}

func (s *StatsService) BRByID(ctx context.Context, id string, params *BRStatsByIDParams) (*BRStatsResponse, error) {
	if s.client.apiKey == "" {
		return nil, ErrNoAPIKey
	}

	if id == "" {
		return nil, emptyParamErr("id")
	}

	return getJSON[BRStatsResponse](ctx, s.client, "/v2/stats/br/v2/"+id, params)
}
