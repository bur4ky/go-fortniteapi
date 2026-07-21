package fortniteapi

type ResponseFlags uint32

type ResponseFlagsParams struct {
	ResponseFlags ResponseFlags `url:"responseFlags,omitempty"`
}

const (
	FlagNone ResponseFlags = 0

	FlagIncludePaths ResponseFlags = 1 << (iota - 1)
	FlagIncludeGameplayTags
	FlagIncludeShopHistory
)
