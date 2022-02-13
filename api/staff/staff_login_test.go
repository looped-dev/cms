package staff

import (
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

	if _, err := StaffRegister(db, staffInput); err != nil {
		log.Fatalf("Unable to create users for testing")
	}

	_, err := StaffLogin(db, &model.StaffLoginInput{
		Email:    "login_test@example.com",
		Password: "password",
	})
	assert.NoError(t, err)
}
