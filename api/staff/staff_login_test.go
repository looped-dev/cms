package staff

import (
	"testing"

	"github.com/looped-dev/cms/api/graph/model"
	"github.com/stretchr/testify/assert"
)

func TestStaffLogin(t *testing.T) {
	staff, err := StaffLogin(db, &model.StaffLoginInput{
		Email:    "login_test@example.com",
		Password: "password",
	})
	assert.NoError(t, err)
	assert.Equal(t, staff.Email, "login_test@example.com")
	assert.Equal(t, staff.Name, "test")
}
