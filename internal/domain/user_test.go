package domain

import (
	"reflect"
	"testing"

	"github.com/gofrs/uuid"
)

func TestNewUser(t *testing.T) {
	id, _ := uuid.NewV4()
	type args struct {
		id       uuid.UUID
		login    string
		password string
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
			args:    args{id: id, login: "login"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "negative test #3",
			args:    args{password: "password", login: "login"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "negative test #4",
			args:    args{id: id, password: "password"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "positive test #1",
			args:    args{id: id, password: "password", login: "login"},
			want:    &User{id: id, password: "password", login: "login"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(tt.args.id, tt.args.login, tt.args.password)
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
