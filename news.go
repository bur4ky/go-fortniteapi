package fortniteapi

import (
	"context"
	"time"
)

type NewsParams LanguageParams
type BRNewsParams LanguageParams
type STWNewsParams LanguageParams
type CreativeNewsParams LanguageParams

type NewsMOTD struct {
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
	MOTDs    []NewsMOTD    `json:"motds"`
	Messages []NewsMessage `json:"messages"`
}

type NewsResponse struct {
	BR       News `json:"br"`
	STW      News `json:"stw"`
	Creative News `json:"creative"`
}

type NewsService struct {
	client *Client
}

func (s *NewsService) All(ctx context.Context, params *NewsParams) (*NewsResponse, error) {
	return s.client.get[*NewsResponse](ctx, "/v2/news", params)
}

func (s *NewsService) BR(ctx context.Context, params *BRNewsParams) (*News, error) {
	return s.client.get[*News](ctx, "/v2/news/br", params)
}

func (s *NewsService) STW(ctx context.Context, params *STWNewsParams) (*News, error) {
	return s.client.get[*News](ctx, "/v2/news/stw", params)
}

func (s *NewsService) Creative(ctx context.Context, params *CreativeNewsParams) (*News, error) {
	return s.client.get[*News](ctx, "/v2/news/creative", params)
}
