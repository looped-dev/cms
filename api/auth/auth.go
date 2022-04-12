package auth

import (
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/looped-dev/cms/api/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewAuthRepository(db *mongo.Client) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

type AuthRepository struct {
	db *mongo.Client
}

func (a AuthRepository) IsAuthorizationHeaderValid(authorizationHeader string) (*StaffJWTClaims, error) {
	// validate token and get user details from token
	jwtRepository := NewJWTRepository(a.db)
	// before verifying, might want to check whether it's basic or bearer
	// token, this is because, loggin into the console use Staff Credentials
	// i.e. a Bearer Token while the frontend will use Basic authentication
	// with an API Key
	authTokenFromHeader := strings.Split(authorizationHeader, "Bearer ")
	spew.Dump(authTokenFromHeader)
	if len(authTokenFromHeader) != 2 {
		return nil, utils.NewGraphQLError(http.StatusUnauthorized, "Invalid Authorization")
	}
	user, err := jwtRepository.VerifyStaffAccessToken(authTokenFromHeader[1])
	if err != nil {
		return nil, utils.NewGraphQLError(http.StatusUnauthorized, err.Error())
	}
	return user, nil
}
