package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/looped-dev/cms/api/graph/generated"
	"github.com/looped-dev/cms/api/models"
)

func (r *staffResolver) Role(ctx context.Context, obj *models.Staff) (models.StaffRole, error) {
	return obj.Role, nil
}

// Staff returns generated.StaffResolver implementation.
func (r *Resolver) Staff() generated.StaffResolver { return &staffResolver{r} }

type staffResolver struct{ *Resolver }
