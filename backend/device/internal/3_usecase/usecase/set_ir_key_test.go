package usecase

import (
	"device/pkg/shared"
	"testing"
)

func TestUseCase_SetIrKey(t *testing.T) {

	type args struct {
		Cc *shared.CommonContent
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal1",
			args: args{
				Cc: &shared.CommonContent{
					Room:   shared.DataRoomCommon,
					Object: shared.DataObjectLight,
					Key:    shared.DataKeyStatus,
					Value:  "lightPower",
				},
			},
			want: "lightPower",
		},
		{
			name: "normal2",
			args: args{
				Cc: &shared.CommonContent{
					Room:   shared.DataRoomCommon,
					Object: shared.DataObjectAircon,
					Key:    shared.DataKeyCoolTemperature,
					Value:  "25",
				},
			},
			want: "coolTemperature25",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			uc := &UseCase{}
			if got := uc.SetIrKey(tt.args.Cc); got != tt.want {
				t.Errorf("UseCase.SetIrKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
