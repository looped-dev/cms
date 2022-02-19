package staff

import (
	"context"
	"log"
	"testing"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestStaffLogin(t *testing.T) {
	// create a login user for testing
	staffInput := &model.StaffRegisterInput{
		Name:     "test",
		Email:    "login_test@example.com",
		Password: "password",
	}
	s := Staff{
		DBClient: dbClient,
	}
	if _, err := s.StaffRegister(context.TODO(), staffInput); err != nil {
		log.Fatalf("Unable to create users for testing")
	}
	staff := Staff{
		DBClient: dbClient,
	}
	_, err := staff.StaffLogin(context.Background(), &model.StaffLoginInput{
		Email:    "login_test@example.com",
		Password: "password",
	})
	assert.NoError(t, err)
}
