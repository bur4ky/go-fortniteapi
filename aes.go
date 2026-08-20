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

type AESKeyParams struct {
	KeyFormat     AESKeyFormat  `url:"keyFormat"`
	ResponseFlags ResponseFlags `url:"responseFlags,omitempty"`
}

type AESDynamicKey struct {
	PakFilename string `json:"pakFilename"`
	PakGUID     string `json:"pakGuid"`
	Key         string `json:"key"`
}

type AESKeyResponse struct {
	Build       string          `json:"build"`
	MainKey     string          `json:"mainKey"`
	DynamicKeys []AESDynamicKey `json:"dynamicKeys"`
	Updated     time.Time       `json:"updated"`
}

type AESService struct {
	client *Client
}

func (s *AESService) Key(ctx context.Context, params *AESKeyParams) (*AESKeyResponse, error) {
	return s.client.get[*AESKeyResponse](ctx, "/v2/aes", params)
}
