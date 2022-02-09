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

func GenerateStaffAccessToken(staff *models.Staff) (string, error) {
	claims := StaffJWTClaims{
		ID:            staff.ID.Hex(),
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

func VerifyStaffAccessToken(tokenString string) (*StaffJWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &StaffJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signInString), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*StaffJWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// CreateStaffRefreshTokenSession creates a new refresh token session for the
// staff and saves in the database. This allows the option to revoke the token
// and also tracking usage of refresh tokens. The refresh tokens will be single
// use and once used, they will be invalidated.
func CreateStaffRefreshTokenSession(staff *models.Staff) {
	panic("Not Implemented")
}

func GenerateStaffRefreshToken(staff *models.Staff) (string, error) {
	panic("Not Implemented")
}

func VerifyStaffRefreshToken(tokenString string) error {
	panic("Not Implemented")
}

func RevokeStaffRefreshToken(tokenString string) error {
	panic("Not Implemented")
}
