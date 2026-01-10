package fortniteapi_test

import (
	"os"
	"testing"

	"github.com/bur4ky/go-fortniteapi"
)

var client *fortniteapi.Client

func TestMain(m *testing.M) {
	apiKey := os.Getenv("FORTNITE_API_KEY")
	client = fortniteapi.New(fortniteapi.LanguageEnglish, apiKey)

	os.Exit(m.Run())
}
