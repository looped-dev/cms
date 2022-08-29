package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/setting"
)

// UpdateSiteSettings is the resolver for the updateSiteSettings field.
func (r *mutationResolver) UpdateSiteSettings(ctx context.Context, input model.UpdateSiteSettingsInput) (*model.SiteSettings, error) {
	setting := setting.NewSettingRepository(r.DB)
	return setting.SaveSettings(ctx, input)
}

// Settings is the resolver for the settings field.
func (r *queryResolver) Settings(ctx context.Context) (*model.SiteSettings, error) {
	setting := setting.NewSettingRepository(r.DB)
	return setting.Details(ctx)
}
