package staff

import (
	"context"
	"fmt"

	"github.com/looped-dev/cms/api/auth"
	"github.com/looped-dev/cms/api/constants"
	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/models"
	"github.com/looped-dev/cms/api/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

func (s StaffRepository) StaffLogin(ctx context.Context, input *model.StaffLoginInput) (*model.StaffLoginResponse, error) {
	staffMember := &models.StaffMember{}
	filter := models.StaffMember{Email: input.Email}
	err := s.DBClient.Database(s.dbName).
		Collection(constants.StaffCollectionName).
		FindOne(ctx, filter).
		Decode(staffMember)
	if err != nil {
		// if there are no documents, it means email address is not available
		if err == mongo.ErrNoDocuments {
			return nil, utils.NewGraphQLErrorWithError(400,
				fmt.Errorf("Invalid email or password"),
			)
		}
		return nil, err
	}
	// check if user account is verified
	if !staffMember.EmailVerified {
		return nil, utils.NewGraphQLErrorWithError(400,
			fmt.Errorf("Email address has not been verified. Check in your email inbox."),
		)
	}
	err = bcrypt.CompareHashAndPassword([]byte(staffMember.HashedPassword), []byte(input.Password))
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(400,
			fmt.Errorf("Invalid email or password"),
		)
	}
	// create session i.e. JWT Access Token, JWT Refresh Token and JWT ID Token
	jwt := auth.NewJWTRepository(s.DBClient)
	accessToken, err := jwt.GenerateStaffAccessToken(staffMember)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500,
			fmt.Errorf("Error generating access token: %v", err),
		)
	}
	refreshToken, err := jwt.CreateStaffRefreshTokenSession(ctx, staffMember)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500,
			fmt.Errorf("Error generating refresh token: %v", err),
		)
	}
	return &model.StaffLoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Staff:        staffMember,
	}, nil
}

// StaffLogout invalidates the current logged in staff refresh token.
func (s StaffRepository) StaffLogout(ctx context.Context) (*models.StaffMember, error) {
	panic("not implemented")
}

func (s StaffRepository) StaffResetPassword(ctx context.Context, input *model.StaffResetPasswordInput) (*models.StaffMember, error) {
	panic("not implemented")
}

func (s StaffRepository) StaffForgotPassword(ctx context.Context, input *model.StaffForgotPasswordInput) (*models.StaffMember, error) {
	panic("not implemented")
}
