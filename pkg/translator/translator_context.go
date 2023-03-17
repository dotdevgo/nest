package translator

import (
	"dotdev/nest/pkg/nest"

	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// LocaleContext godoc
type LocaleContext struct {
	nest.Context
}

// Context godoc
func Context(ctx nest.Context) (*LocaleContext, *i18n.Localizer) {
	cc := &LocaleContext{Context: ctx}
	localizer := cc.Get("localizer").(*i18n.Localizer)

	return cc, localizer
}
