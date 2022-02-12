package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/looped-dev/cms/api/graph/generated"
	"github.com/looped-dev/cms/api/models"
)

func (r *staffResolver) Role(ctx context.Context, obj *models.StaffMember) (models.StaffRole, error) {
	panic(fmt.Errorf("not implemented"))
}

// Staff returns generated.StaffResolver implementation.
func (r *Resolver) Staff() generated.StaffResolver { return &staffResolver{r} }

type staffResolver struct{ *Resolver }
