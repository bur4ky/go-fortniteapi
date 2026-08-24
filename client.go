package fortniteapi

import (
	"bytes"
	"context"
	"encoding/json/v2"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	querypkg "github.com/google/go-querystring/query"
	"golang.org/x/time/rate"
)

const version = "v1.0.0"

var (
	ErrMissingAPIKey  = fmt.Errorf("missing API key")
	ErrInvalidAPIKey  = fmt.Errorf("invalid API key")
	ErrEmptyParameter = fmt.Errorf("missing required parameter")
)

type Response[T any] struct {
	Status int `json:"status"`
	Data   T   `json:"data"`
}

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"error"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("api error: %d - %s", e.Status, e.Message)
}

type Client struct {
	http     *http.Client
	baseURL  *url.URL
	apiKey   string
	language Language

	AES         *AESService
	Banners     *BannersService
	Cosmetics   *CosmeticsService
	CreatorCode *CreatorCodeService
	Map         *MapService
	News        *NewsService
	Playlists   *PlaylistsService
	Shop        *ShopService
	Stats       *StatsService
}

func NewClient(language Language, apiKey string) *Client {
	return NewWithClient(language, apiKey, nil)
}

func NewWithClient(language Language, apiKey string, client *http.Client) *Client {
	if language == "" {
		language = LanguageEnglish
	}

	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}

	c := &Client{
		http:     client,
		apiKey:   apiKey,
		language: language,
	}

	c.baseURL, _ = url.Parse("https://fortnite-api.com")
	c.AES = &AESService{client: c}
	c.Banners = &BannersService{client: c}
	c.Cosmetics = &CosmeticsService{client: c}
	c.CreatorCode = &CreatorCodeService{client: c}
	c.Map = &MapService{client: c}
	c.News = &NewsService{client: c}
	c.Playlists = &PlaylistsService{client: c}
	c.Shop = &ShopService{client: c}
	c.Stats = &StatsService{
		client:  c,
		limiter: rate.NewLimiter(3, 3),
	}

	return c
}

func (c *Client) get[T any](ctx context.Context, path string, query any) (T, error) {
	return c.do[T](ctx, http.MethodGet, path, query, nil)
}

func (c *Client) do[T any](ctx context.Context, method, path string, query, body any) (T, error) {
	var zero T
	fullURL, err := c.fullURL(path, query)
	if err != nil {
		return zero, err
	}

	request, err := c.newRequest(ctx, method, fullURL.String(), body)
	if err != nil {
		return zero, err
	}

	response, err := c.http.Do(request)
	if err != nil {
		return zero, fmt.Errorf("send request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		var respBody Error
		if err := json.UnmarshalRead(response.Body, &respBody); err != nil {
			return zero, fmt.Errorf("decode error response (status %d): %w", response.StatusCode, err)
		}

		return zero, &respBody
	}

	var respBody Response[T]
	if err := json.UnmarshalRead(response.Body, &respBody); err != nil {
		return zero, fmt.Errorf("decode response: %w", err)
	}

	return respBody.Data, nil
}

func (c *Client) newRequest(ctx context.Context, method, urlStr string, body any) (*http.Request, error) {
	var bodyReader io.Reader
	if body != nil {
		var buf bytes.Buffer
		if err := json.MarshalWrite(&buf, body); err != nil {
			return nil, fmt.Errorf("marshal request body: %w", err)
		}

		bodyReader = &buf
	}

	request, err := http.NewRequestWithContext(ctx, method, urlStr, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	request.Header.Set("User-Agent", "go-fortniteapi/"+version)

	if c.apiKey != "" {
		request.Header.Set("Authorization", c.apiKey)
	}

	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	return request, nil
}

func (c *Client) fullURL(path string, query any) (*url.URL, error) {
	fullURL := c.baseURL.JoinPath(path)

	var err error
	params := url.Values{}
	if values, ok := query.(url.Values); ok {
		params = values
	} else if query != nil {
		params, err = querypkg.Values(query)
		if err != nil {
			return nil, err
		}
	}

	if c.language != "" && !params.Has("language") {
		params.Set("language", string(c.language))
	}

	fullURL.RawQuery = params.Encode()
	return fullURL, nil
}

func emptyParamErr(name string) error {
	return fmt.Errorf("%w: %s", ErrEmptyParameter, name)
}
