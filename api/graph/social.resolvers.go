package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/setting"
)

// UpdateTwitterCardSettings is the resolver for the updateTwitterCardSettings field.
func (r *mutationResolver) UpdateTwitterCardSettings(ctx context.Context, input model.UpdateTwitterCardSettingsInput) (*model.SiteSettings, error) {
	settingRepo := setting.NewSettingRepository(r.DB)
	return settingRepo.UpdateTwitterCardSettings(ctx, input)
}

// UpdateFacebookCardSettings is the resolver for the updateFacebookCardSettings field.
func (r *mutationResolver) UpdateFacebookCardSettings(ctx context.Context, input model.UpdateFacebookCardSettingsInput) (*model.SiteSettings, error) {
	settingRepo := setting.NewSettingRepository(r.DB)
	return settingRepo.UpdateFacebookCardSettings(ctx, input)
}
