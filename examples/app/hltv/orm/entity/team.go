package entity

import "github.com/dotdevgo/nest/pkg/crud"

type (
	Team struct {
		crud.Model

		name string
	}
)
