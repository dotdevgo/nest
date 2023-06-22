package config

import (
	"github.com/matcornic/hermes/v2"
)

func Hermes() *hermes.Hermes {
	return &hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			Name:      "DotDev",
			Link:      "https://dotdevio.com/",
			Logo:      "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
			Copyright: "Copyright © 2022 DotDev. All rights reserved",
		},
	}
}
