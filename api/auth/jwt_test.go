package auth

import (
	"testing"

	"github.com/looped-dev/cms/api/models"
	"github.com/stretchr/testify/assert"
)

func TestJWT_GenerateStaffAccessToken(t *testing.T) {
	jwtToken := JWTRepository{}
	token, err := jwtToken.GenerateStaffAccessToken(&models.StaffMember{})
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}
