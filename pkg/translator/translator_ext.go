package translator

import (
	"encoding/json"

	"github.com/BurntSushi/toml"
	"github.com/goava/di"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// New godoc
func New() di.Option {
	return di.Options(
		di.Provide(func() *i18n.Bundle {
			var bundle = i18n.NewBundle(language.English)
			bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
			bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

			return bundle
		}),
		di.Provide(func(bundle *i18n.Bundle) *i18n.Localizer {
			var localizer = i18n.NewLocalizer(bundle)
			return localizer
		}),
	)
}
