package fortniteapi

type ResponseFlag uint32

type ResponseFlagsParams struct {
	ResponseFlags ResponseFlag `url:"responseFlags,omitempty"`
}

const (
	FlagIncludePaths ResponseFlag = 1 << iota
	FlagIncludeGameplayTags
	FlagIncludeShopHistory

	FlagAll = FlagIncludePaths | FlagIncludeGameplayTags | FlagIncludeShopHistory
)
