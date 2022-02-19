package staff

import (
	"context"
	"fmt"

	"github.com/looped-dev/cms/api/auth"
	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/models"
	"golang.org/x/crypto/bcrypt"
)

func (s Staff) StaffLogin(ctx context.Context, input *model.StaffLoginInput) (*model.StaffLoginResponse, error) {
	staffMember := &models.StaffMember{}
	err := s.DBClient.Database("cms").Collection("staff").FindOne(
		context.TODO(),
		models.StaffMember{
			Email: input.Email,
		},
	).Decode(staffMember)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(staffMember.HashedPassword), []byte(input.Password))
	if err != nil {
		return nil, fmt.Errorf("Incorrect password: %v", err)
	}
	// create session i.e. JWT Access Token, JWT Refresh Token and JWT ID Token
	accessToken, err := auth.GenerateStaffAccessToken(staffMember)
	if err != nil {
		return nil, fmt.Errorf("error generating access token: %v", err)
	}
	return &model.StaffLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: "",
		Staff:        staffMember,
	}, nil
}

// StaffLogout invalidates the current logged in staff refresh token.
func (s Staff) StaffLogout(ctx context.Context) (*models.StaffMember, error) {
	panic("not implemented")
}

func (s Staff) StaffResetPassword(ctx context.Context, input *model.StaffResetPasswordInput) (*models.StaffMember, error) {
	panic("not implemented")
}

func (s Staff) StaffForgotPassword(ctx context.Context, input *model.StaffForgotPasswordInput) (*models.StaffMember, error) {
	panic("not implemented")
}
