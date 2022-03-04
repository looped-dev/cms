package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/looped-dev/cms/api/graph/generated"
	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/setting"
	"github.com/looped-dev/cms/api/staff"
)

func (r *queryResolver) IsSetup(ctx context.Context) (bool, error) {
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) SiteSettings(ctx context.Context) (*model.SiteSettings, error) {
	panic(fmt.Errorf("not implemented"))
}
