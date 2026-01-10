package fortniteapi

import (
	"context"
	"time"
)

// AESKeyFormat represents the format of the AES key.
//
// Default: AESKeyFormatHex
type AESKeyFormat string

const (
	AESKeyFormatBase64 AESKeyFormat = "base64"
	AESKeyFormatHex    AESKeyFormat = "hex"
)

type AESKeysParams struct {
	KeyFormat     AESKeyFormat `url:"keyFormat"`
	ResponseFlags ResponseFlag `url:"responseFlags,omitempty"`
}

type AESDynamicKey struct {
	PakFilename string `json:"pakFilename"`
	PakGUID     string `json:"pakGuid"`
	Key         string `json:"key"`
}

type AESKeysResponse struct {
	Build       string          `json:"build"`
	MainKey     string          `json:"mainKey"`
	DynamicKeys []AESDynamicKey `json:"dynamicKeys"`
	Updated     time.Time       `json:"updated"`
}

type AESService struct {
	client *Client
}

func (s *AESService) Keys(ctx context.Context, params *AESKeysParams) (*AESKeysResponse, error) {
	return getJSON[AESKeysResponse](ctx, s.client, "/v2/aes", params)
}
