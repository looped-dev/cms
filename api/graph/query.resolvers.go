package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/looped-dev/cms/api/graph/generated"
	"github.com/looped-dev/cms/api/setting"
	"github.com/looped-dev/cms/api/staff"
)

func (r *queryResolver) IsSiteSetup(ctx context.Context) (bool, error) {
	staff := staff.NewStaff(r.SMTPClient, r.DB)

	// check if settings exist
	setting := setting.NewSetting(r.DB)
	settingExists, err := setting.Exists(ctx)
	if err != nil {
		return false, err
	}
	if settingExists {
		return true, nil
	}

	// check if staff exists
	return staff.StaffExists(ctx)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
