package graphql

import (
	"github.com/dotdevgo/gosymfony/examples/app/dataform/dto"
	"github.com/dotdevgo/gosymfony/examples/app/dataform/orm/entity"
	"github.com/dotdevgo/gosymfony/examples/app/dataform/orm/repository"
	"github.com/dotdevgo/gosymfony/pkg/crud"
	"github.com/dotdevgo/gosymfony/pkg/nest/kernel"
	"errors"
	"gorm.io/gorm/clause"
)

// FormTemplateResolver godoc
type FormTemplateResolver struct {
	kernel.Controller
	Crud *crud.Service
	Repo *repository.FormTemplateRepo
}

// Save godoc
func (ctrl *FormTemplateResolver) Save(input dto.FormTemplateDto) (*entity.FormTemplate, error) {
	data, err := ctrl.Repo.FindOrNew(&input)
	if err != nil {
		return nil, err
	}

	if err := ctrl.Crud.Save(data); err != nil {
		return nil, err
	}
	return data, nil
}

// ContentList godoc
func (ctrl *FormTemplateResolver) List(cursor *crud.PaginatorCursor) (*dto.FormTemplatePaginator, error) {
	var model []*entity.FormTemplate
	result, err := ctrl.Crud.Paginate(
		&model,
		crud.WithCursor(cursor),
		crud.WithPreload(clause.Associations),
	)
	if err != nil {
		return nil, err
	}

	records, ok := result.Records.(*[]*entity.FormTemplate)
	if !ok {
		return nil, errors.New("invalid records")
	}

	return &dto.FormTemplatePaginator{Records: *records, Cursor: result}, nil
}
