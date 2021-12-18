package auth

import (
	"testing"

	"golang.org/x/crypto/bcrypt"

	"github.com/looped-dev/cms/api/graph/model"
)

func Test_hashPassword(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "hash a password - John@DoePassword",
			args: args{
				password: "John@DoePassword",
			},
			wantErr: false,
		},
		{
			name: "test with empty password",
			args: args{
				password: "",
			},
			wantErr: true,
		},
		{
			name: "hash password - John#$3Doe292#92$929",
			args: args{
				password: "John#$3Doe292#92$929",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := hashPassword(tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("hashPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if want error is true, then we expect an error and test passes
			if (err != nil) == tt.wantErr {
				return
			}
			// check the password returned and compare against the original to see if
			// they match
			errHash := bcrypt.CompareHashAndPassword([]byte(got), []byte(tt.args.password))
			if errHash != nil {
				t.Errorf("hashed password doesn't match the input password, error = %v", errHash)
				return
			}
		})
	}
}

func TestRegister(t *testing.T) {
	type args struct {
		input model.LoginInput
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Register(tt.args.input)
		})
	}
}
