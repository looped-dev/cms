package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/looped-dev/cms/api/graph/generated"
	"github.com/looped-dev/cms/api/graph/model"
)

func (r *queryResolver) GetPosts(ctx context.Context, page *int, perPage *int) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPost(ctx context.Context, slug string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPage(ctx context.Context, slug string) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPostByID(ctx context.Context, id string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetPageByID(ctx context.Context, id string) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) SiteSettings(ctx context.Context) (*model.SiteSettings, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) IsSetup(ctx context.Context) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
