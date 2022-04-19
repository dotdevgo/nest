package config

import (
	hr "github.com/matcornic/hermes/v2"
)

func Hermes() *hr.Hermes {
	return &hr.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hr.Product{
			Name:      "DotDevio",
			Link:      "https://dotdevio.com/",
			Logo:      "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
			Copyright: "Copyright © 2022 DotDevio. All rights reserved",
		},
	}
}
