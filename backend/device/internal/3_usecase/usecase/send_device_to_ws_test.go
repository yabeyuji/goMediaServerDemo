package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock "device/internal/3_usecase/usecase/mock_repository"
	"device/pkg/shared"
)

func TestUseCase_SendDeviceToWs(t *testing.T) {
	type request struct {
		Address string
		Cc      *shared.CommonContent
	}

	type response struct {
		Value string
		Err   error
	}

	type mocks struct {
		request  *request
		response *response
	}

	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		mocks   mocks
	}{
		{
			name: "normal1",
			args: args{
				value: "{\"bed\":{\"Aircon\":{\"Status\":\"airconStop\",\"WarmTemperature\":23,\"CoolTemperature\":18},\"Light\":{\"Status\":\"lightPower\"}},\"living\":{\"Aircon\":{\"Status\":\"airconStop\",\"WarmTemperature\":23,\"CoolTemperature\":18},\"Light\":{\"Status\":\"lightPower\"}}}",
			},
			mocks: mocks{
				request: &request{
					Address: shared.GRPCAddressWs,
					Cc: &shared.CommonContent{
						Room:   shared.DataRoomCommon,
						Object: shared.DataObjectDevice,
						Key:    shared.DataKeyDevices,
						Value:  "{\"bed\":{\"Aircon\":{\"Status\":\"airconStop\",\"WarmTemperature\":23,\"CoolTemperature\":18},\"Light\":{\"Status\":\"lightPower\"}},\"living\":{\"Aircon\":{\"Status\":\"airconStop\",\"WarmTemperature\":23,\"CoolTemperature\":18},\"Light\":{\"Status\":\"lightPower\"}}}",
					},
				},
				response: &response{
					Value: "ok",
					Err:   nil,
				},
			},
			want:    "ok",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock.NewMockToService(ctrl)

			mockService.EXPECT().
				SvSendContent(tt.mocks.request.Address, tt.mocks.request.Cc).
				Return(tt.mocks.response.Value, tt.mocks.response.Err)

			uc := &UseCase{ToService: mockService}

			got, err := uc.SendDeviceToWs(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.SendDeviceToWs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UseCase.SendDeviceToWs() = %v, want %v", got, tt.want)
			}
		})
	}
}
