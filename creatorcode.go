package fortniteapi

import "context"

type CreatorCodeParams struct {
	Name          string       `url:"name"`
	ResponseFlags ResponseFlag `url:"responseFlags,omitempty"`
}

type CreatorCodeAccount struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreatorCodeResponse struct {
	Code     string             `json:"code"`
	Account  CreatorCodeAccount `json:"account"`
	Status   string             `json:"status"`
	Verified bool               `json:"verified"`
}

type CreatorCodeService struct {
	client *Client
}

func (s *CreatorCodeService) Get(ctx context.Context, name string, params *CreatorCodeParams) (*CreatorCodeResponse, error) {
	if name == "" {
		return nil, emptyParamErr("name")
	}

	if params == nil {
		params = &CreatorCodeParams{}
	}

	params.Name = name

	return getJSON[CreatorCodeResponse](ctx, s.client, "/v2/creatorcode", params)
}
