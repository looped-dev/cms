package staff

import (
	"context"
	"testing"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestStaff_StaffLogin(t *testing.T) {
	// create a login user for testing
	staffInput := &model.StaffRegisterInput{
		Name:     "test",
		Email:    "login_test@example.com",
		Password: "password",
	}
	s := Staff{
		DBClient: dbClient,
	}
	_, err := s.StaffRegister(context.TODO(), staffInput)
	assert.ErrorIs(t, err, nil, "shouldn't return error when creating staff registered")

	staff := Staff{
		DBClient: dbClient,
	}
	staffLogin, err := staff.StaffLogin(context.Background(), &model.StaffLoginInput{
		Email:    "login_test@example.com",
		Password: "password",
	})
	assert.NoError(t, err)
	assert.NotNil(t, staffLogin)
}
