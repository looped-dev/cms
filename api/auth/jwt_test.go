package auth

import (
	"testing"

	"github.com/looped-dev/cms/api/models"
	"github.com/stretchr/testify/assert"
)

func TestGenerateStaffJWTToken(t *testing.T) {
	type args struct {
		staff *models.Staff
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "should generate a valid JWT token",
			args: args{
				staff: &models.Staff{
					Name:  "John Doe",
					Email: "johndoe@example.com",
					ID:    "123456789",
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
