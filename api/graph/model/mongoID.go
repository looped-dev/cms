package model

import (
	"fmt"
	"io"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UnmarshalMongoID takes a GraphQL string and parses it to primitive.ObjectID
func UnmarshalMongoID(value interface{}) (primitive.ObjectID, error) {
	id, ok := value.(string)
	if !ok {
		return primitive.NewObjectID(), fmt.Errorf("ID must be a string")
	}
	return primitive.ObjectIDFromHex(id)
}

// MarshalMongoID takes primitive.ObjectID and passes it to gqlgen MongoID scalar
func MarshalMongoID(timestamp primitive.ObjectID) graphql.Marshaler {
	return graphql.WriterFunc(func(writer io.Writer) {
		_, err := writer.Write([]byte(timestamp.Hex()))
		if err != nil {
			log.Printf("%v", err)
		}
	})
}
