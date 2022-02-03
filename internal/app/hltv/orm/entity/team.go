package entity

import "dotdev.io/pkg/crud"

type (
	Team struct {
		crud.Model

		name string
	}
)
