package utils

import (
	"fmt"
)

func NewGraphQLError(code int, message string) *GraphQLError {
	return &GraphQLError{
		Code:    code,
		Message: message,
	}
}

func NewGraphQLErrorWithError(code int, err error) *GraphQLError {
	return &GraphQLError{
		Code:    code,
		Message: err.Error(),
	}
}

type GraphQLError struct {
	Code    int
	Message string
}

func (err GraphQLError) Error() string {
	return fmt.Sprintf("Error: code: %s, message: %d", err.Message, err.Code)
}
