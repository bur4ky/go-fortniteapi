package fortniteapi

import (
	"context"
	"time"
)

type NewsParams LanguageParams
type BRNewsParams LanguageParams
type STWNewsParams LanguageParams
type CreativeNewsParams LanguageParams

type NewsMotd struct {
	ID              string `json:"id"`
	Title           string `json:"title"`
	TabTitle        string `json:"tabTitle"`
	Body            string `json:"body"`
	Image           string `json:"image"`
	TileImage       string `json:"tileImage"`
	SortingPriority int    `json:"sortingPriority"`
	Hidden          bool   `json:"hidden"`
	WebsiteURL      string `json:"websiteUrl"`
	VideoString     string `json:"videoString"`
	VideoID         string `json:"videoId"`
}

type NewsMessage struct {
	Title   string `json:"title"`
	Body    string `json:"body"`
	Image   string `json:"image"`
	Adspace string `json:"adspace"`
}

type News struct {
	Hash     string        `json:"hash"`
	Date     time.Time     `json:"date"`
	Image    string        `json:"image"`
	Motds    []NewsMotd    `json:"motds"`
	Messages []NewsMessage `json:"messages"`
}

type NewsResponse struct {
	BR       News `json:"br"`
	STW      News `json:"stw"`
	Creative News `json:"creative"`
}

type BRNewsResponse News
type STWNewsResponse News
type CreativeNewsResponse News

type NewsService struct {
	client *Client
}

func (s *NewsService) All(ctx context.Context, params *NewsParams) (*NewsResponse, error) {
	return getJSON[NewsResponse](ctx, s.client, "/v2/news", params)
}

func (s *NewsService) BR(ctx context.Context, params *BRNewsParams) (*BRNewsResponse, error) {
	return getJSON[BRNewsResponse](ctx, s.client, "/v2/news/br", params)
}

func (s *NewsService) STW(ctx context.Context, params *STWNewsParams) (*STWNewsResponse, error) {
	return getJSON[STWNewsResponse](ctx, s.client, "/v2/news/stw", params)
}

func (s *NewsService) Creative(ctx context.Context, params *CreativeNewsParams) (*CreativeNewsResponse, error) {
	return getJSON[CreativeNewsResponse](ctx, s.client, "/v2/news/creative", params)
}
