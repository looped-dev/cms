package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/setting"
)

// UpdateSEOSettings is the resolver for the updateSEOSettings field.
func (r *mutationResolver) UpdateSEOSettings(ctx context.Context, input model.UpdateSEOSettingsInput) (*model.SiteSettings, error) {
	settingRepo := setting.NewSettingRepository(r.DB)
	return settingRepo.UpdateSEOSettings(ctx, input)
}
