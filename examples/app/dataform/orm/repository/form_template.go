package repository

import (
	"github.com/dotdevgo/gosymfony/examples/app/dataform/dto"
	"github.com/dotdevgo/gosymfony/examples/app/dataform/orm/entity"
	"github.com/dotdevgo/gosymfony/pkg/crud"
	"github.com/dotdevgo/gosymfony/pkg/goutils"
	"github.com/goava/di"
)

type (
	FormTemplateRepo struct {
		di.Inject
		Crud *crud.Service
	}
)

// Find godoc
func (r *FormTemplateRepo) FindOrNew(input *dto.FormTemplateDto) (*entity.FormTemplate, error) {
	var data entity.FormTemplate
	if err := r.Crud.Find(&data, input.UUID); err != nil {
		return nil, err
	}

	goutils.Copy(&data, input)

	return &data, nil
}
