package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/looped-dev/cms/api/graph/model"
)

// UpdatePageStatus is the resolver for the updatePageStatus field.
func (r *mutationResolver) UpdatePageStatus(ctx context.Context, input model.UpdatePageStatusInput) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdatePage is the resolver for the updatePage field.
func (r *mutationResolver) UpdatePage(ctx context.Context, input model.UpdatePageInput) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetPage is the resolver for the getPage field.
func (r *queryResolver) GetPage(ctx context.Context, slug string) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetPageByID is the resolver for the getPageByID field.
func (r *queryResolver) GetPageByID(ctx context.Context, id string) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}
