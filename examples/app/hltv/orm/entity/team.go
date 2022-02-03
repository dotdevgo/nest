package entity

import "github.com/dotdevgo/gosymfony/pkg/crud"

type (
	Team struct {
		crud.Model

		name string
	}
)
