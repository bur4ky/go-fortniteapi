package fortniteapi

// Language represents languages supported by the API.
type Language string

const (
	LanguageArabic               Language = "ar"
	LanguageGerman               Language = "de"
	LanguageEnglish              Language = "en"
	LanguageSpanish              Language = "es"
	LanguageLatinAmericanSpanish Language = "es-419"
	LanguageFrench               Language = "fr"
	LanguageIndonesian           Language = "id"
	LanguageItalian              Language = "it"
	LanguageJapanese             Language = "ja"
	LanguageKorean               Language = "ko"
	LanguagePolish               Language = "pl"
	LanguageBrazilianPortuguese  Language = "pt-BR"
	LanguageRussian              Language = "ru"
	LanguageThai                 Language = "th"
	LanguageTurkish              Language = "tr"
	LanguageVietnamese           Language = "vi"
	LanguageSimplifiedChinese    Language = "zh-Hans"
	LanguageTraditionalChinese   Language = "zh-Hant"
)

type LanguageParams struct {
	Language      Language      `url:"language,omitempty"`
	ResponseFlags ResponseFlags `url:"responseFlags,omitempty"`
}
