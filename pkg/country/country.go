package country

import (
	"dotdev/nest/pkg/crud"
)

const (
	DBTableCountries = "countries"
)

type Country struct {
	crud.Model

	Code string `json:"code"`
	Name string `json:"name"`
}
