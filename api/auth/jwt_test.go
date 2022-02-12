package auth

import (
	"testing"

	"github.com/looped-dev/cms/api/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGenerateStaffJWTToken(t *testing.T) {
	type args struct {
		staff *models.StaffMember
	}
	id, _ := primitive.ObjectIDFromHex("123456789")
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should generate a valid JWT token",
			args: args{
				staff: &models.StaffMember{
					Name:  "John Doe",
					Email: "johndoe@example.com",
					ID:    id,
					Role:  models.StaffRoleAdministrator,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateStaffAccessToken(tt.args.staff)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
			assert.NotNil(t, got)
		})
	}
}
