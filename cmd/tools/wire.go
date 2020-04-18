//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/tarxzfv/go-deepl/internal/pkg/translation"

	http2 "github.com/tarxzfv/go-deepl/internal/pkg/http"
	"github.com/tarxzfv/go-deepl/pkg/deepl"
)

func InitializeTranslator(config *deepl.Config) (translation.Translator, error) {
	wire.Build(
		translation.NewTranslator,
		deepl.NewDeepLClient,
		http2.NewDefaultHTTPClient,
	)
	return nil, nil
}
