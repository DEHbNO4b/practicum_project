package domain

import (
	"reflect"
	"testing"
	"time"
)

func TestNewOrder(t *testing.T) {
	type args struct {
		number  int
		accrual int
		status  string
		time    time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    *Order
		wantErr bool
	}{
		{
			name:    "negative test",
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "positive test",
			args:    args{number: 1},
			want:    &Order{number: 1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewOrder(tt.args.number, tt.args.accrual, tt.args.status, tt.args.time)
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
