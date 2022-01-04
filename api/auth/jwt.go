package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/looped-dev/cms/api/models"
)

// todo: dynamically generate this during setup and store in config
const signInString = "AllYourBase"

type StaffJWTClaims struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Role          string `json:"role"`
	EmailVerified bool   `json:"emailVerified"`
	jwt.StandardClaims
}

func GenerateStaffJWTToken(staff *models.Staff) (string, error) {
	claims := StaffJWTClaims{
		ID:            staff.ID,
		Name:          staff.Name,
		Email:         staff.Email,
		Role:          staff.Role.String(),
		EmailVerified: staff.EmailVerified,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: int64(time.Hour.Seconds()),
			Issuer:    "looped-cms",
			Audience:  "looped-cms-admin",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signInString))
}
