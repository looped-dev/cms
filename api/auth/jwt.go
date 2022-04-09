package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/looped-dev/cms/api/db"
	"github.com/looped-dev/cms/api/models"
	"go.mongodb.org/mongo-driver/mongo"
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

type StaffJWTRefreshTokenClaims struct {
	models.RefreshToken
	jwt.StandardClaims
}

func NewJWTRepository(dbClient *mongo.Client) JWTRepository {
	return JWTRepository{
		DBClient: dbClient,
		dbName:   db.GetDatabaseName(),
	}
}

type JWTRepository struct {
	DBClient *mongo.Client
	dbName   string
}

func (webTokens JWTRepository) GenerateStaffAccessToken(staff *models.StaffMember) (string, error) {
	claims := StaffJWTClaims{
		ID:            staff.ID.Hex(),
		Name:          staff.Name,
		Email:         staff.Email,
		Role:          staff.Role.String(),
		EmailVerified: staff.EmailVerified,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    "looped-cms",
			Audience:  "looped-cms-admin",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signInString))
}

func (webTokens JWTRepository) VerifyStaffAccessToken(tokenString string) (*StaffJWTClaims, error) {
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
func (webTokens JWTRepository) CreateStaffRefreshTokenSession(client *mongo.Client, ctx context.Context, staff *models.StaffMember) (string, error) {
	src := NewStaffRefreshToken(webTokens.DBClient)
	refreshTokenData, err := src.CreateStaffRefreshTokenSession(ctx, staff)
	if err != nil {
		return "", fmt.Errorf("Error creating new refresh token: %v", err)
	}
	claims := &StaffJWTRefreshTokenClaims{
		RefreshToken: *refreshTokenData,
		StandardClaims: jwt.StandardClaims{
			// stay active for up to 30 Days
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			Issuer:    "looped-cms",
			Audience:  "looped-cms-admin",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(signInString))
}

func (webTokens JWTRepository) GenerateStaffRefreshToken(ctx context.Context, staff *models.StaffMember) (string, error) {
	panic("Not Implemented")
}

func (webTokens JWTRepository) VerifyStaffRefreshToken(ctx context.Context, tokenString string) error {
	panic("Not Implemented")
}

func (webTokens JWTRepository) RevokeStaffRefreshToken(ctx context.Context, tokenString string) error {
	panic("Not Implemented")
}
