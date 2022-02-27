package model

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUnmarshalMongoID_ThrowError(t *testing.T) {
	_, err := UnmarshalMongoID("")
	assert.Error(t, err, "ID must be a string")
}

func TestUnmarshalMongoID_ParseCorrectDate(t *testing.T) {
	objectID := primitive.NewObjectID().Hex()
	unmarshalObjectID, err := UnmarshalMongoID(objectID)
	assert.NoError(t, err)
	assert.Equal(t, objectID, unmarshalObjectID.Hex())
}

func TestUnMarshalMongoID(t *testing.T) {
	buffer := &bytes.Buffer{}
	objectID := primitive.NewObjectID()
	MarshalMongoID(objectID).MarshalGQL(buffer)
	assert.Equal(t, objectID.Hex(), buffer.String())
}
