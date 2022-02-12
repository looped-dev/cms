package graph

import (
	mail "github.com/xhit/go-simple-mail/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate go run github.com/99designs/gqlgen
//go:generate go run github.com/looped-dev/cms/api/graph/plugins/modelgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB         *mongo.Client
	SMTPClient *mail.SMTPClient
}
