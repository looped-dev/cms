package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/looped-dev/cms/api/graph/generated"
	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/models"
	"github.com/looped-dev/cms/api/setting"
	"github.com/looped-dev/cms/api/staff"
)

func (r *mutationResolver) InitialSetup(ctx context.Context, input model.InitialSetupInput) (*model.InitialSetupResponse, error) {
	// first step, configure settings
	setting := setting.NewSettingRepository(r.DB)
	_, err := setting.SaveSettings(ctx, model.SiteSettingsInput{
		SiteName: input.SiteName,
		// todo: probably set this one as well
		BaseURL: "",
	})
	if err != nil {
		return nil, err
	}
	// next, register staff member
	staff := staff.NewStaffRepository(r.SMTPClient, r.DB)
	staffMember, err := staff.StaffRegister(ctx, &model.StaffRegisterInput{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		Role:     models.StaffRoleOwner,
	})
	if err != nil {
		return nil, err
	}
	// login in the user
	logins, err := staff.StaffLogin(ctx, &model.StaffLoginInput{
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}
	response := &model.InitialSetupResponse{
		AccessToken:  logins.AccessToken,
		RefreshToken: logins.RefreshToken,
		Staff:        staffMember,
	}
	return response, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
