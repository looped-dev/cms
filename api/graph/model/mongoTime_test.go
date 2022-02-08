package model

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUnmarshalMongoTime_ThrowError(t *testing.T) {
	_, err := UnmarshalMongoTime("")
	assert.Error(t, err, "Time must be a string")
}

func TestUnmarshalMongoTime_ParseCorrectDate(t *testing.T) {
	timestamp, err := UnmarshalMongoTime("2022-02-08T20:00:00.000Z")
	assert.NoError(t, err)
	println(timestamp.T)
	assert.Nil(t, err)
	assert.Equal(t, uint32(1644350400), timestamp.T)
}

func TestMarshalMongoTime(t *testing.T) {
	// set timezone to UTC
	os.Setenv("TZ", "UTC")
	buffer := &bytes.Buffer{}
	MarshalMongoTime(primitive.Timestamp{T: uint32(1644350400)}).MarshalGQL(buffer)
	assert.Equal(t, "2022-02-08T20:00:00Z", buffer.String())
}
