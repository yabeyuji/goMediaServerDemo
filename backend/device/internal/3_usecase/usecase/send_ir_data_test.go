package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock "device/internal/3_usecase/usecase/mock_repository"
)

func TestUseCase_SendIRData(t *testing.T) {
	type request struct {
		room  string
		irKey string
	}

	type response struct {
		Err error
	}

	type mocks struct {
		request  *request
		response *response
	}

	type args struct {
		room  string
		irKey string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mocks   mocks
	}{
		{
			name: "normal1",
			args: args{
				room:  "bed",
				irKey: "lightPower",
			},
			mocks: mocks{
				request: &request{
					room:  "bed",
					irKey: "lightPower",
				},
				response: &response{
					Err: nil,
				},
			},
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
				SvSendIRData(tt.mocks.request.room, tt.mocks.request.irKey).
				Return(tt.mocks.response.Err)

			uc := &UseCase{ToService: mockService}

			if err := uc.SendIRData(tt.args.room, tt.args.irKey); (err != nil) != tt.wantErr {
				t.Errorf("UseCase.SendIRData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
