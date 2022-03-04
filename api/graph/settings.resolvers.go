package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/setting"
)

func (r *mutationResolver) UpdateSiteSettings(ctx context.Context, input model.SiteSettingsInput) (*model.SiteSettings, error) {
	setting := setting.NewSetting(r.DB)
	return setting.SaveSettings(ctx, input)
}

func (r *queryResolver) Settings(ctx context.Context) (*model.SiteSettings, error) {
	panic(fmt.Errorf("not implemented"))
}
