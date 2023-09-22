package domain

import (
	"reflect"
	"testing"
	"time"
)

func TestNewOrder(t *testing.T) {
	type args struct {
		number      string
		status      string
		accrual     float64
		upploadedAt time.Time
		userID      int
	}
	tests := []struct {
		name    string
		args    args
		want    *Order
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
			args:    args{number: "8463473g"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "positive test",
			args:    args{number: "1235235"},
			want:    &Order{number: "1235235"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOrder(tt.args.number, tt.args.status, tt.args.accrual, tt.args.upploadedAt, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
