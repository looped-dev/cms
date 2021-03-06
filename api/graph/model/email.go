package model

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
)

// VerifyEmailAddress verifies that the input is a valid email address
func VerifyEmailAddress(email string) bool {
	const emailRegexString = `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	emailRegex := regexp.MustCompile(emailRegexString)
	return emailRegex.MatchString(email)
}

// UnMarshalEmail verifies that the input is a valid email address
func UnmarshalEmail(value interface{}) (string, error) {
	email, ok := value.(string)
	if !ok {
		return "", fmt.Errorf("Email must be a string")
	}
	if !VerifyEmailAddress(email) {
		return "", fmt.Errorf("%s is not a valid email", email)
	}
	return email, nil
}

// MarshalEmail passes email address to gqlgen email scalar
func MarshalEmail(email string) graphql.Marshaler {
	return graphql.WriterFunc(func(writer io.Writer) {
		if email != "" {
			if VerifyEmailAddress(email) {
				_, err := writer.Write([]byte(strconv.Quote(email)))
				if err != nil {
					log.Printf("%v", err)
				}
			}
		}
	})
}
