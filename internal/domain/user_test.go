package domain

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		id       int
		login    string
		password string
		balance  int
	}
	tests := []struct {
		name    string
		args    args
		want    *User
		wantErr bool
	}{
		{
			name:    "negative test #1",
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "negative test #2",
			args:    args{login: "login"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "negative test #3",
			args:    args{password: "password"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "positive test #1",
			args:    args{password: "password", login: "login"},
			want:    &User{password: "password", login: "login"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.id, tt.args.login, tt.args.password, tt.args.balance)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
