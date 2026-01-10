package fortniteapi

import (
	"context"
	"time"
)

type PlaylistsParams LanguageParams
type PlaylistByIDParams PlaylistsParams

type PlaylistsImages struct {
	Showcase    string `json:"showcase"`
	MissionIcon string `json:"missionIcon"`
}

type Playlist struct {
	ID                       string          `json:"id"`
	Name                     string          `json:"name"`
	SubName                  string          `json:"subName"`
	Description              string          `json:"description"`
	GameType                 string          `json:"gameType"`
	RatingType               string          `json:"ratingType"`
	MinPlayers               int             `json:"minPlayers"`
	MaxPlayers               int             `json:"maxPlayers"`
	MaxTeams                 int             `json:"maxTeams"`
	MaxTeamSize              int             `json:"maxTeamSize"`
	MaxSquads                int             `json:"maxSquads"`
	MaxSquadSize             int             `json:"maxSquadSize"`
	IsDefault                bool            `json:"isDefault"`
	IsTournament             bool            `json:"isTournament"`
	IsLimitedTimeMode        bool            `json:"isLimitedTimeMode"`
	IsLargeTeamGame          bool            `json:"isLargeTeamGame"`
	AccumulateToProfileStats bool            `json:"accumulateToProfileStats"`
	Images                   PlaylistsImages `json:"images"`
	GameplayTags             []string        `json:"gameplayTags"`
	Path                     string          `json:"path"`
	Added                    time.Time       `json:"added"`
}

type PlaylistsResponse []Playlist
type PlaylistByIDResponse Playlist

type PlaylistsService struct {
	client *Client
}

func (s *PlaylistsService) All(ctx context.Context, params *PlaylistsParams) (PlaylistsResponse, error) {
	return getJSONSlice[PlaylistsResponse](ctx, s.client, "/v1/playlists", params)
}

func (s *PlaylistsService) ByID(ctx context.Context, playlistID string, params *PlaylistByIDParams) (*PlaylistByIDResponse, error) {
	if playlistID == "" {
		return nil, emptyParamErr("playlistID")
	}

	return getJSON[PlaylistByIDResponse](ctx, s.client, "/v1/playlists/"+playlistID, params)
}
