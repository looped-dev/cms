package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/looped-dev/cms/api/auth"
	"github.com/looped-dev/cms/api/constants"
	"github.com/looped-dev/cms/api/graph/generated"
	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/models"
	"github.com/looped-dev/cms/api/staff"
	"github.com/looped-dev/cms/api/utils"
)

func (r *mutationResolver) StaffLogin(ctx context.Context, input model.StaffLoginInput) (*model.StaffLoginResponse, error) {
	staff := staff.NewStaffRepository(r.SMTPClient, r.DB)
	return staff.StaffLogin(ctx, &input)
}

func (r *mutationResolver) StaffInvite(ctx context.Context, input model.StaffInviteInput) (*models.StaffMember, error) {
	staff := staff.NewStaffRepository(r.SMTPClient, r.DB)
	return staff.StaffSendInvite(ctx, &input)
}

func (r *mutationResolver) StaffAcceptInvite(ctx context.Context, input model.StaffAcceptInviteInput) (*models.StaffMember, error) {
	staff := staff.NewStaffRepository(r.SMTPClient, r.DB)
	return staff.StaffAcceptInvite(ctx, &input)
}

func (r *mutationResolver) StaffUpdate(ctx context.Context, input model.StaffUpdateInput) (*models.StaffMember, error) {
	staff := staff.NewStaffRepository(r.SMTPClient, r.DB)
	return staff.StaffUpdate(ctx, &input)
}

func (r *mutationResolver) StaffDelete(ctx context.Context, input model.StaffDeleteInput) (*models.StaffMember, error) {
	staff := staff.NewStaffRepository(r.SMTPClient, r.DB)
	return staff.StaffDelete(ctx, &input)
}

func (r *mutationResolver) StaffChangePassword(ctx context.Context, input model.StaffChangePasswordInput) (*models.StaffMember, error) {
	staff := staff.NewStaffRepository(r.SMTPClient, r.DB)
	return staff.StaffChangePassword(ctx, &input)
}

func (r *mutationResolver) StaffResetPassword(ctx context.Context, input model.StaffResetPasswordInput) (*models.StaffMember, error) {
	staff := staff.NewStaffRepository(r.SMTPClient, r.DB)
	return staff.StaffResetPassword(ctx, &input)
}

func (r *mutationResolver) StaffForgotPassword(ctx context.Context, input model.StaffForgotPasswordInput) (*models.StaffMember, error) {
	staff := staff.NewStaffRepository(r.SMTPClient, r.DB)
	return staff.StaffForgotPassword(ctx, &input)
}

func (r *mutationResolver) StaffLogout(ctx context.Context) (bool, error) {
	staff := staff.NewStaffRepository(r.SMTPClient, r.DB)
	_, err := staff.StaffLogout(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) StaffRefreshToken(ctx context.Context, input model.StaffRefreshTokenInput) (*model.StaffLoginResponse, error) {
	refreshTokenRepository := auth.NewStaffRefreshToken(r.DB)
	// get user from context
	user := ctx.Value(constants.CurrentlyAuthenticatedUserContextKey)
	if user == nil {
		return nil, utils.NewGraphQLErrorWithError(401, fmt.Errorf("Access Denied!"))
	}
	userDetails, ok := user.(auth.StaffJWTClaims)
	if !ok {
		return nil, utils.NewGraphQLErrorWithError(500, fmt.Errorf("Internal Error"))
	}
	refreshToken, err := refreshTokenRepository.VerifyStaffRefreshToken(ctx, userDetails.ID, input.RefreshToken)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500, fmt.Errorf("Error validating refresh token: %v", err))
	}
	// invalidate token and create a new one
	_, err = refreshTokenRepository.InvalidateRefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, utils.NewGraphQLError(500,
			fmt.Sprintf("Error validating refresh token: %v", err),
		)
	}
	// create new session
	jwt := auth.NewJWTRepository(r.DB)
	userRepository := staff.NewStaffRepository(r.SMTPClient, r.DB)
	staffMember, err := userRepository.FetchStaffFromDB(ctx, userDetails.Email)
	if err != nil {
		return nil, utils.NewGraphQLError(500,
			fmt.Sprintf("Error fetching user details: %v", err),
		)
	}
	newAccessToken, err := jwt.GenerateStaffAccessToken(staffMember)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500,
			fmt.Errorf("Error generating access token: %v", err),
		)
	}
	newRefreshToken, err := jwt.CreateStaffRefreshTokenSession(ctx, staffMember)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500,
			fmt.Errorf("Error generating refresh token: %v", err),
		)
	}
	newTokens := model.StaffLoginResponse{
		AccessToken:  newAccessToken,
		Staff:        staffMember,
		RefreshToken: newRefreshToken,
	}
	return &newTokens, nil
}

func (r *staffResolver) Role(ctx context.Context, obj *models.StaffMember) (models.StaffRole, error) {
	if obj.Role.IsValid() {
		return "", fmt.Errorf("Your role is invalid")
	}
	return models.StaffRole(obj.Role.String()), nil
}

// Staff returns generated.StaffResolver implementation.
func (r *Resolver) Staff() generated.StaffResolver { return &staffResolver{r} }

type staffResolver struct{ *Resolver }
