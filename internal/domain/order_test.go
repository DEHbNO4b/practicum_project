package domain

import (
	"reflect"
	"testing"
)

func TestNewOrder(t *testing.T) {
	type args struct {
		number int
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
			got, err := NewOrder(tt.args.number)
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
