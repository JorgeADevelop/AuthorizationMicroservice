package utils

import (
	"encoding/json"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func InitTranslator() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)

	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("translations/en.json")
	bundle.MustLoadMessageFile("translations/es.json")

	return bundle
}
