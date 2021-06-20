package usecase

import (
	"device/internal/4_domain/domain"
	"device/pkg/shared"
	"reflect"
	"testing"
)

func TestUseCase_ChangeDeviceStatus(t *testing.T) {
	type fields struct {
		ToDomain  ToDomain
		ToService ToService
	}
	type args struct {
		cc    *shared.CommonContent
		rooms map[string]*domain.Room
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]*domain.Room
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &UseCase{
				ToDomain:  tt.fields.ToDomain,
				ToService: tt.fields.ToService,
			}
			if got := uc.ChangeDeviceStatus(tt.args.cc, tt.args.rooms); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.ChangeDeviceStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
