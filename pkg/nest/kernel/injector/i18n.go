package injector

import (
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/goava/di"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// I18n godoc
func I18n() di.Option {
	return di.Options(
		di.Provide(func() *i18n.Bundle {
			var bundle *i18n.Bundle
			bundle = i18n.NewBundle(language.English)
			bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
			bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

			return bundle
		}),
	)
}
