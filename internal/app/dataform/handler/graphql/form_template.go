package graphql

import (
	"dotdev.io/internal/app/dataform/dto"
	"dotdev.io/internal/app/dataform/orm/entity"
	"dotdev.io/internal/app/dataform/orm/repository"
	"dotdev.io/pkg/crud"
	"dotdev.io/pkg/nest/httpkernel"
	"errors"
	"gorm.io/gorm/clause"
)

type FormTemplateResolver struct {
	httpkernel.Controller
	Crud *crud.Service
}

// Save godoc
func (ctrl *FormTemplateResolver) Save(input dto.FormTemplateDto) (*entity.FormTemplate, error) {
	var data = repository.NewFormTemplate(&input)
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
