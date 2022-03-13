package model

import (
	"bytes"
	"strconv"
	"testing"
)

func TestVerifyEmailAddress(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid email address: johndoe@example.com",
			args: args{
				email: "johndoe@example.com",
			},
			want: true,
		},
		{
			name: "Valid email address: janedoe@example.com",
			args: args{
				email: "janedoe@example.com",
			},
			want: true,
		},
		{
			name: "InValid email address: johndoeexample.com",
			args: args{
				email: "johndoeexample.com",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyEmailAddress(tt.args.email); got != tt.want {
				t.Errorf("VerifyEmailAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnMarshalEmail(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Valid email address: johndoe@example.com",
			args: args{
				value: "johndoe@example.com",
			},
			want:    "johndoe@example.com",
			wantErr: false,
		},
		{
			name: "InValid email address: janedoeexample.com",
			args: args{
				value: "janedoeexample.com",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Valid email address: janedoe@example.com",
			args: args{
				value: "janedoe@example.com",
			},
			want:    "janedoe@example.com",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshalEmail(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnMarshalEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UnMarshalEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMarshalEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Valid email address: johndoe@example.com",
			args: args{
				email: "johndoe@example.com",
			},
			want: strconv.Quote("johndoe@example.com"),
		},
		{
			name: "Valid email address: Empty string",
			args: args{
				email: "",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MarshalEmail(tt.args.email)
			buffer := &bytes.Buffer{}
			got.MarshalGQL(buffer)

			if tt.want != "" {
				if buffer.String() != tt.want {
					t.Errorf("MarshalEmail() = %v, want %v", buffer.String(), tt.want)
					return
				}
			} else {
				if buffer.String() != "" {
					t.Errorf("MarshalEmail() = %v, want %v", buffer.String(), tt.want)
					return
				}
			}
		})
	}
}
