package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/looped-dev/cms/api/graph/generated"
	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/models"
	"github.com/looped-dev/cms/api/staff"
)

func (r *mutationResolver) StaffLogin(ctx context.Context, input model.StaffLoginInput) (*model.StaffLoginResponse, error) {
	staff := staff.NewStaff(r.SMTPClient, r.DB)
	return staff.StaffLogin(ctx, &input)
}

func (r *mutationResolver) StaffInvite(ctx context.Context, input model.StaffInviteInput) (*models.StaffMember, error) {
	staff := staff.NewStaff(r.SMTPClient, r.DB)
	return staff.StaffSendInvite(ctx, &input)
}

func (r *mutationResolver) StaffAcceptInvite(ctx context.Context, input model.StaffAcceptInviteInput) (*models.StaffMember, error) {
	staff := staff.NewStaff(r.SMTPClient, r.DB)
	return staff.StaffAcceptInvite(ctx, &input)
}

func (r *mutationResolver) StaffUpdate(ctx context.Context, input model.StaffUpdateInput) (*models.StaffMember, error) {
	staff := staff.NewStaff(r.SMTPClient, r.DB)
	return staff.StaffUpdate(ctx, &input)
}

func (r *mutationResolver) StaffDelete(ctx context.Context, input model.StaffDeleteInput) (*models.StaffMember, error) {
	staff := staff.NewStaff(r.SMTPClient, r.DB)
	return staff.StaffDelete(ctx, &input)
}

func (r *mutationResolver) StaffChangePassword(ctx context.Context, input model.StaffChangePasswordInput) (*models.StaffMember, error) {
	staff := staff.NewStaff(r.SMTPClient, r.DB)
	return staff.StaffChangePassword(ctx, &input)
}

func (r *mutationResolver) StaffResetPassword(ctx context.Context, input model.StaffResetPasswordInput) (*models.StaffMember, error) {
	staff := staff.NewStaff(r.SMTPClient, r.DB)
	return staff.StaffResetPassword(ctx, &input)
}

func (r *mutationResolver) StaffForgotPassword(ctx context.Context, input model.StaffForgotPasswordInput) (*models.StaffMember, error) {
	staff := staff.NewStaff(r.SMTPClient, r.DB)
	return staff.StaffForgotPassword(ctx, &input)
}

func (r *mutationResolver) StaffLogout(ctx context.Context) (bool, error) {
	staff := staff.NewStaff(r.SMTPClient, r.DB)
	_, err := staff.StaffLogout(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *staffResolver) Role(ctx context.Context, obj *models.StaffMember) (models.StaffRole, error) {
	panic(fmt.Errorf("not implemented"))
}

// Staff returns generated.StaffResolver implementation.
func (r *Resolver) Staff() generated.StaffResolver { return &staffResolver{r} }

type staffResolver struct{ *Resolver }
