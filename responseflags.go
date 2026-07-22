package fortniteapi

type ResponseFlags uint32

const (
	FlagIncludePaths ResponseFlags = 1 << iota
	FlagIncludeGameplayTags
	FlagIncludeShopHistory
)

type ResponseFlagsParams struct {
	ResponseFlags ResponseFlags `url:"responseFlags,omitempty"`
}
