package usecase

import (
	"errors"
	mock "media/internal/3_usecase/usecase/mock_repository"
	"media/pkg/shared"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUseCase_SendProgressToWs(t *testing.T) {
	type request struct {
		address string
		cc      *shared.CommonContent
	}

	type response struct {
		value string
		err   error
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
			name: "normal",
			args: args{
				value: "0.5",
			},
			mocks: mocks{
				request: &request{
					address: shared.GRPCAddressWs,
					cc: &shared.CommonContent{
						Room:   shared.DataRoomCommon,
						Object: shared.DataObjectVlc,
						Key:    shared.DataKeyProgress,
						Value:  "0.5",
					},
				},
				response: &response{
					value: "ok",
					err:   nil,
				},
			},
			want:    "ok",
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				value: "0.5",
			},
			mocks: mocks{
				request: &request{
					address: shared.GRPCAddressWs,
					cc: &shared.CommonContent{
						Room:   shared.DataRoomCommon,
						Object: shared.DataObjectVlc,
						Key:    shared.DataKeyProgress,
						Value:  "0.5",
					},
				},
				response: &response{
					value: "",
					err:   errors.New("error"),
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock.NewMockToService(ctrl)

			mockService.EXPECT().
				SvSendContent(tt.mocks.request.address, tt.mocks.request.cc).
				Return(tt.mocks.response.value, tt.mocks.response.err)

			uc := &UseCase{ToService: mockService}

			got, err := uc.SendProgressToWs(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.SendProgressToWs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UseCase.SendProgressToWs() = %v, want %v", got, tt.want)
			}
		})
	}
}
