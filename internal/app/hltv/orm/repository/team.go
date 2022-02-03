package repository

import (
	"dotdev.io/internal/app/hltv/dto"
	"dotdev.io/internal/app/hltv/orm/entity"
	"dotdev.io/pkg/crud"
	"dotdev.io/pkg/goutils"
	"github.com/goava/di"
)

type (
	TeamRepo struct {
		di.Inject
		Crud *crud.Service
	}
)

// Find godoc
func (r *TeamRepo) FindOrNew(input *dto.TeamDto) (*entity.Team, error) {
	var data entity.Team
	if err := r.Crud.Find(&data, input.UUID); err != nil {
		return nil, err
	}

	goutils.Copy(&data, input)

	return &data, nil
}
