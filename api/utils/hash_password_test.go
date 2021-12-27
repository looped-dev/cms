package utils

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
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
			got, err := HashPassword(tt.args.password)
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
