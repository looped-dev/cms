package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/looped-dev/cms/api/graph/model"
)

// UpdatePostStatus is the resolver for the updatePostStatus field.
func (r *mutationResolver) UpdatePostStatus(ctx context.Context, input model.UpdatePostStatusInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// UpdatePost is the resolver for the updatePost field.
func (r *mutationResolver) UpdatePost(ctx context.Context, input model.UpdatePostInput) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetPosts is the resolver for the getPosts field.
func (r *queryResolver) GetPosts(ctx context.Context, page *int, perPage *int) ([]*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetPost is the resolver for the getPost field.
func (r *queryResolver) GetPost(ctx context.Context, slug string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}

// GetPostByID is the resolver for the getPostByID field.
func (r *queryResolver) GetPostByID(ctx context.Context, id string) (*model.Post, error) {
	panic(fmt.Errorf("not implemented"))
}
