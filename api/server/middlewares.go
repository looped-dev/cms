package main

import (
	"context"
	"net/http"

	"github.com/looped-dev/cms/api/constants"
	"go.mongodb.org/mongo-driver/mongo"
)

func authenticationMiddleware(client *mongo.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authorizationHeader := r.Header.Get("Authorization")
			// allow unauthenticated users through
			if authorizationHeader == "" {
				next.ServeHTTP(w, r)
				return
			}
			// put the header into context, authorization will happen as needed
			// instead of all routes. This will use isAuthorized directives among
			// other directives for authorization
			ctx := context.WithValue(r.Context(), constants.BearerAuthorizationHeader, authorizationHeader)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
