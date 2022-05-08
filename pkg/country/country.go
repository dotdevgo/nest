package country

import (
	"dotdev/nest/pkg/crud"
)

const (
	DBTableCountry = "countries"
)

type Country struct {
	crud.Model

	Code string `json:"code"`
	Name string `json:"name"`
}
