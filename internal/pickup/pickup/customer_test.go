package pickup

import (
	"context"
	"errors"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/entity"
	"github.com/agadilkhan/pickup-point-service/internal/pickup/repository/mocks"
	"reflect"
	"testing"
)

func TestService_GetCustomerByID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.Customer
		wantErr bool
	}{
		{
			name: "success: customer exists",
			args: args{id: 1},
			want: &entity.Customer{
				ID:    1,
				Email: "customer1@gmail.com",
			},
			wantErr: false,
		},
		{
			name:    "fail: customer not exists",
			args:    args{id: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success: customer exists",
			args: args{id: 2},
			want: &entity.Customer{
				ID:    2,
				Email: "customer2@gmail.com",
			},
			wantErr: false,
		},
		{
			name: "success: customer exists",
			args: args{id: 3},
			want: &entity.Customer{
				ID:    3,
				Email: "customer3@gmail.com",
			},
			wantErr: false,
		},
		{
			name:    "fail: customer not exists",
			args:    args{id: 2},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repomock := &mocks.Repository{}
			s := &Service{
				repo: repomock,
			}
			if !tt.wantErr {
				repomock.On("GetCustomerByID", context.TODO(), tt.args.id).Return(tt.want, nil)
			} else {
				repomock.On("GetCustomerByID", context.TODO(), tt.args.id).Return(nil, errors.New("error"))
			}

			got, err := s.GetCustomerByID(context.TODO(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomerByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomerByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
