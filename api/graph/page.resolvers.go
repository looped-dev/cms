package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/looped-dev/cms/api/graph/model"
)

func (r *mutationResolver) UpdatePageStatus(ctx context.Context, input model.UpdatePageStatusInput) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePage(ctx context.Context, input model.UpdatePageInput) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPage(ctx context.Context, slug string) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPageByID(ctx context.Context, id string) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}
