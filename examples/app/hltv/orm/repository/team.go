package repository

import (
	"github.com/dotdevgo/gosymfony/examples/app/hltv/dto"
	"github.com/dotdevgo/gosymfony/examples/app/hltv/orm/entity"
	"github.com/dotdevgo/gosymfony/pkg/crud"
	"github.com/dotdevgo/gosymfony/pkg/goutils"
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
