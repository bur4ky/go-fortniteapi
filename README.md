# Go Wrapper for [Fortnite-API.com](https://fortnite-api.com)

[![Go Reference](https://pkg.go.dev/badge/github.com/bur4ky/go-fortniteapi.svg)](https://pkg.go.dev/github.com/bur4ky/go-fortniteapi)
[![Release](https://img.shields.io/github/v/release/bur4ky/go-fortniteapi)](https://github.com/bur4ky/go-fortniteapi/releases)
[![License](https://img.shields.io/github/license/bur4ky/go-fortniteapi)](LICENSE)
[![Discord](https://img.shields.io/discord/621452110558527502?label=Discord&logo=discord)](https://discord.gg/eysmvFT2rV)

A simple Go wrapper for [Fortnite-API.com](https://fortnite-api.com), with support for all available endpoints.

## Installation

```sh
go get github.com/bur4ky/go-fortniteapi
```

## Quick Example

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bur4ky/go-fortniteapi"
)

func main() {
	client := fortniteapi.New(fortniteapi.LanguageEnglish, os.Getenv("FORTNITE_API_KEY"))

	flags := fortniteapi.CombineFlags(
		fortniteapi.FlagIncludePaths,
		fortniteapi.FlagIncludeGameplayTags,
	)

	searchParams := &fortniteapi.SearchBRCosmeticParams{
		Name:          "Peely",
		ResponseFlags: flags,
	}

	cosmetic, err := client.Cosmetics.SearchBRCosmetic(context.TODO(), searchParams)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(cosmetic.ID)
}
```

## Links

- [Fortnite-API Documentation](https://dash.fortnite-api.com)
- [pkg.go.dev](https://pkg.go.dev/github.com/bur4ky/go-fortniteapi)
- [Discord](https://discord.gg/eysmvFT2rV)

## License

This project is licensed under the [MIT License](LICENSE).