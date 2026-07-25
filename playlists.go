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
	RatingType               string          `json:"ratingType,omitempty"`
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
	Images                   PlaylistsImages `json:"images,omitzero"`
	GameplayTags             []string        `json:"gameplayTags,omitempty"`
	Path                     string          `json:"path"`
	Added                    time.Time       `json:"added"`
}

type PlaylistsService struct {
	client *Client
}

func (s *PlaylistsService) All(ctx context.Context, params *PlaylistsParams) ([]Playlist, error) {
	return getJSON[[]Playlist](ctx, s.client, "/v1/playlists", params)
}

func (s *PlaylistsService) ByID(ctx context.Context, id string, params *PlaylistByIDParams) (*Playlist, error) {
	if id == "" {
		return nil, emptyParamErr("id")
	}

	return getJSON[*Playlist](ctx, s.client, "/v1/playlists/"+id, params)
}
