package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/looped-dev/cms/api/auth"
	"github.com/looped-dev/cms/api/constants"
	"go.mongodb.org/mongo-driver/mongo"
)

func authenticationMiddleware(client *mongo.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get("Authorization")
			spew.Dump(authorizationHeader)
			// allow unauthenticated users through
			if authorizationHeader == "" {
				next.ServeHTTP(w, r)
				return
			}
			// validate token and get user details from token
			jwtRepository := auth.NewJWTRepository(client)
			// before verifying, might want to check whether it's basic or bearer
			// token, this is because, loggin into the console use Staff Credentials
			// i.e. a Bearer Token while the frontend will use Basic authentication
			// with an API Key
			authTokenFromHeader := strings.Split(authorizationHeader, "Bearer ")
			spew.Dump(authTokenFromHeader)
			if len(authTokenFromHeader) != 2 {
				http.Error(w, "Invalid Authorization", http.StatusUnauthorized)
				return
			}
			user, err := jwtRepository.VerifyStaffAccessToken(authTokenFromHeader[1])
			spew.Dump(err)
			if err != nil {
				http.Error(w, "Invalid Authorization", http.StatusUnauthorized)
			}
			ctx := context.WithValue(r.Context(), constants.CurrentlyAuthenticatedUserContextKey, user)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
