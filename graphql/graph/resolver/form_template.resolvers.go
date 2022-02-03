package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/dotdevgo/gosymfony/examples/app/dataform/dto"
	"github.com/dotdevgo/gosymfony/examples/app/dataform/orm/entity"
	"github.com/dotdevgo/gosymfony/graphql/graph/generated"
	"github.com/dotdevgo/gosymfony/pkg/crud"
)

func (r *formTemplateResolver) ID(ctx context.Context, obj *entity.FormTemplate) (string, error) {
	return obj.UUID, nil
}

func (r *mutationResolver) CreateFormTemplate(ctx context.Context, input dto.FormTemplateDto) (*entity.FormTemplate, error) {
	return r.FormTemplateResolver.Save(input)
}

func (r *queryResolver) FormTemplateList(ctx context.Context, cursor *crud.PaginatorCursor) (*dto.FormTemplatePaginator, error) {
	return r.FormTemplateResolver.List(cursor)
}

// FormTemplate returns generated.FormTemplateResolver implementation.
func (r *Resolver) FormTemplate() generated.FormTemplateResolver { return &formTemplateResolver{r} }

type formTemplateResolver struct{ *Resolver }
