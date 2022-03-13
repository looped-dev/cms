package model

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UnmarshalMongoTime verifies that the input is a valid ISO time and parses it
// to primitive.Timestamp
func UnmarshalMongoTime(value interface{}) (primitive.Timestamp, error) {
	timestamp, ok := value.(string)
	if !ok {
		return primitive.Timestamp{}, fmt.Errorf("Time must be a string")
	}
	t, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		return primitive.Timestamp{}, fmt.Errorf("Time must be a valid ISO time")
	}
	fmt.Println(t)
	return primitive.Timestamp{
		T: uint32(t.Unix()),
	}, nil
}

// MarshalMongoTime passes the primitive.Timestamp to gqlgen MongoTime scalar
func MarshalMongoTime(timestamp primitive.Timestamp) graphql.Marshaler {
	return graphql.WriterFunc(func(writer io.Writer) {
		timeValue := time.Unix(int64(timestamp.T), 0)
		_, err := writer.Write([]byte(strconv.Quote(timeValue.Format(time.RFC3339))))
		if err != nil {
			log.Printf("%v", err)
		}
	})
}
