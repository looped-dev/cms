package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/looped-dev/cms/api/graph/generated"
	"github.com/looped-dev/cms/api/graph/model"
)

func (r *mutationResolver) UpdatePostStatus(ctx context.Context, input model.UpdatePostStatusInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePost(ctx context.Context, input model.UpdatePostInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePageStatus(ctx context.Context, input model.UpdatePostStatusInput) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePage(ctx context.Context, input model.UpdatePostInput) (*model.Page, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input *model.LoginInput) (*model.LoginResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Register(ctx context.Context, input *model.RegisterInput) (*model.RegisterResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Logout(ctx context.Context) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateSiteSettings(ctx context.Context, input model.SiteSettingsInput) (*model.SiteSettings, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
