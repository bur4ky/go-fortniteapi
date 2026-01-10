package fortniteapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	querypkg "github.com/google/go-querystring/query"
)

const (
	version = "v1.0.0"
	baseURL = "https://fortnite-api.com"
)

var (
	ErrNoAPIKey       = fmt.Errorf("missing API key")
	ErrEmptyParameter = fmt.Errorf("missing required parameter")
)

type APIResponse[T any] struct {
	Status int `json:"status"`
	Data   T   `json:"data"`
}

type APIError struct {
	Status  int    `json:"status"`
	Message string `json:"error"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("api error: %d - %s", e.Status, e.Message)
}

type Client struct {
	httpClient *http.Client
	apiKey     string
	language   Language

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

func New(language Language, apiKey string) *Client {
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
		httpClient: client,
		apiKey:     apiKey,
		language:   language,
	}

	c.AES = &AESService{client: c}
	c.Banners = &BannersService{client: c}
	c.Cosmetics = &CosmeticsService{client: c}
	c.CreatorCode = &CreatorCodeService{client: c}
	c.Map = &MapService{client: c}
	c.News = &NewsService{client: c}
	c.Playlists = &PlaylistsService{client: c}
	c.Shop = &ShopService{client: c}
	c.Stats = &StatsService{client: c}

	return c
}

func (c *Client) fetch(ctx context.Context, method, path string, query, body, out any) error {
	fullURL, err := c.buildURL(path, query)
	if err != nil {
		return err
	}

	request, err := c.newRequest(ctx, method, fullURL.String(), body)
	if err != nil {
		return err
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		var apiError APIError
		if err := decoder.Decode(&apiError); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}

		return &apiError
	}

	var apiResponse APIResponse[json.RawMessage]
	if err := decoder.Decode(&apiResponse); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if out != nil {
		if err := json.Unmarshal(apiResponse.Data, out); err != nil {
			return fmt.Errorf("failed to unmarshal data from response: %w", err)
		}
	}

	return nil
}

func (c *Client) newRequest(ctx context.Context, method, urlStr string, body any) (*http.Request, error) {
	var bodyReader io.Reader

	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}

		bodyReader = bytes.NewReader(jsonBytes)
	}

	request, err := http.NewRequestWithContext(ctx, method, urlStr, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
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

func (c *Client) buildURL(path string, query any) (*url.URL, error) {
	fullURL, err := url.Parse(baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	params := url.Values{}
	if query != nil {
		if values, ok := query.(url.Values); ok {
			params = values
		} else {
			params, err = querypkg.Values(query)
			if err != nil {
				return nil, fmt.Errorf("failed to encode query: %w", err)
			}
		}
	}

	if c.language != "" && !params.Has("language") {
		params.Set("language", string(c.language))
	}

	fullURL.RawQuery = params.Encode()
	return fullURL, nil
}

// getJSON is used internally to avoid repetitive code when making GET requests that return JSON data.
func getJSON[T any](ctx context.Context, c *Client, path string, query any) (*T, error) {
	var out T
	if err := c.fetch(ctx, http.MethodGet, path, query, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// getJSONSlice is used internally to avoid repetitive code when making GET requests that return JSON slice data.
func getJSONSlice[T any](ctx context.Context, c *Client, path string, query any) (T, error) {
	var out T
	err := c.fetch(ctx, http.MethodGet, path, query, nil, &out)
	return out, err
}

func emptyParamErr(name string) error {
	return fmt.Errorf("%w: %s", ErrEmptyParameter, name)
}
