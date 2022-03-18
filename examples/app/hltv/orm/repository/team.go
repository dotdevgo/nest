package repository

import (
	"github.com/dotdevgo/nest/examples/app/hltv/dto"
	"github.com/dotdevgo/nest/examples/app/hltv/orm/entity"
	"github.com/dotdevgo/nest/pkg/crud"
	"github.com/dotdevgo/nest/pkg/goutils"
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
