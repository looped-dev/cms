package staff

import (
	"context"
	"fmt"
	"time"

	"github.com/looped-dev/cms/api/constants"
	"github.com/looped-dev/cms/api/db"
	"github.com/looped-dev/cms/api/emails"
	"github.com/looped-dev/cms/api/graph/model"
	"github.com/looped-dev/cms/api/models"
	"github.com/looped-dev/cms/api/utils"
	mail "github.com/xhit/go-simple-mail/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type StaffRepository struct {
	SMTPClient *mail.SMTPClient
	DBClient   *mongo.Client
	dbName     string
}

func NewStaffRepository(smtpClient *mail.SMTPClient, dbClient *mongo.Client) *StaffRepository {
	return &StaffRepository{
		SMTPClient: smtpClient,
		DBClient:   dbClient,
		dbName:     db.GetDatabaseName(),
	}
}

// StaffRegister creates a new staff (admin users) and returns the Staff object.
func (s StaffRepository) StaffRegister(ctx context.Context, input *model.StaffRegisterInput) (*models.StaffMember, error) {
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	createdAt := primitive.Timestamp{
		T: uint32(time.Now().Unix()),
	}
	staff := &models.StaffMember{
		Name:           input.Name,
		Email:          input.Email,
		HashedPassword: hashedPassword,
		Role:           input.Role,
		EmailVerified:  true,
		CreatedAt:      createdAt,
		UpdatedAt:      createdAt,
	}
	result, err := s.DBClient.Database(s.dbName).
		Collection(constants.StaffCollectionName).
		InsertOne(ctx, staff)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return nil, utils.NewGraphQLErrorWithError(400, fmt.Errorf("Email already exists"))
		}
		return nil, utils.NewGraphQLErrorWithError(500, err)
	}
	staff.ID = result.InsertedID.(primitive.ObjectID)
	return staff, nil
}

// StaffSendInvite creates a new staff, with a specific role and creates an invite
// code and sends an email to the staff member.
func (s StaffRepository) StaffSendInvite(ctx context.Context, input *model.StaffInviteInput) (*models.StaffMember, error) {
	code := utils.GenerateInviteCode()
	staffMember := &models.StaffMember{
		Email: input.Email,
		Role:  input.Role,
		InviteCode: models.InviteCode{
			Code: code,
			Expiry: primitive.Timestamp{
				T: uint32(time.Now().Unix() + 60*60*24),
			},
		},
	}
	staffMember, err := s.addNewStaffToDB(ctx, staffMember)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500, fmt.Errorf("Error saving users: %v", err))
	}
	// send email
	err = emails.SendEmail(s.SMTPClient, emails.SendMailConfig{
		EmailTo:   staffMember.Email,
		EmailFrom: "info@looped.dev",
		Subject:   "Invite to Looped CMS",
		HtmlBody:  "Hi,<br><br>You have been invited to Looped CMS. Please click the link below to register.<br><br><a href=\"https://looped.dev/staff/register?code=" + code + "\">Register</a>",
		PlainBody: "Hi,\n\nYou have been invited to Looped CMS. Please click the link below to register.\n\nhttps://looped.dev/staff/register?code=" + code,
	})
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(500, fmt.Errorf("Error sending emails: %v", err))
	}
	return staffMember, nil
}

// StaffAcceptInvite verify invite code and set the new staff password and email
// as verified.
func (s StaffRepository) StaffAcceptInvite(ctx context.Context, input *model.StaffAcceptInviteInput) (*models.StaffMember, error) {
	if input.ConfirmPassword != input.Password {
		return nil, utils.NewGraphQLErrorWithError(400,
			fmt.Errorf("Password and confirm password do not match"),
		)
	}
	staffMember, err := s.fetchStaffFromDB(ctx, input.Email)
	if err != nil {
		return nil, utils.NewGraphQLErrorWithError(400, fmt.Errorf("Error fetching staff: %v", err))
	}
	// check if invite code is valid
	if err := validateInviteCode(input.Code, staffMember.InviteCode); err != nil {
		return nil, utils.NewGraphQLErrorWithError(400, err)
	}
	// update staff in database
	if err := s.updateStaffInDB(ctx, staffMember, input); err != nil {
		return nil, utils.NewGraphQLErrorWithError(500, fmt.Errorf("Error updating staff: %v", err))
	}
	return staffMember, nil
}

// StaffUpdate updates the details of the staff i.e. Name, Email, Role.
func (s StaffRepository) StaffUpdate(ctx context.Context, input *model.StaffUpdateInput) (*models.StaffMember, error) {
	panic("not implemented")
}

// StaffDelete soft deletes the staff from the database by adding a delatedAt field.
func (s StaffRepository) StaffDelete(ctx context.Context, input *model.StaffDeleteInput) (*models.StaffMember, error) {
	panic("not implemented")
}

// StaffChangePassword update the staff password.
func (s StaffRepository) StaffChangePassword(ctx context.Context, input *model.StaffChangePasswordInput) (*models.StaffMember, error) {
	panic("not implemented")
}

func (s StaffRepository) StaffExists(ctx context.Context) (bool, error) {
	count, err := s.DBClient.Database(s.dbName).
		Collection(constants.StaffCollectionName).
		CountDocuments(ctx, bson.M{})
	if err != nil {
		return false, utils.NewGraphQLErrorWithError(500, err)
	}
	if count < 1 {
		return false, nil
	}
	return true, nil
}
