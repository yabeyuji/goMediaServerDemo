package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	mock "file/internal/3_usecase/usecase/mock_repository"
	"file/pkg/shared"
)

func TestUseCase_SendFilesToWs(t *testing.T) {
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
				value: "[{\"Vid\":\"a\",\"Name\":\"a\",\"Valid\":true}]",
			},
			mocks: mocks{
				request: &request{
					Address: shared.GRPCAddressWs,
					Cc: &shared.CommonContent{
						Room:   shared.DataRoomCommon,
						Object: shared.DataObjectFile,
						Key:    shared.DataKeyFiles,
						Value:  "[{\"Vid\":\"a\",\"Name\":\"a\",\"Valid\":true}]",
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
		{
			name: "error1",
			args: args{
				value: "[{\"Vid\":\"a\",\"Name\":\"a\",\"Valid\":true}]",
			},
			mocks: mocks{
				request: &request{
					Address: shared.GRPCAddressWs,
					Cc: &shared.CommonContent{
						Room:   shared.DataRoomCommon,
						Object: shared.DataObjectFile,
						Key:    shared.DataKeyFiles,
						Value:  "[{\"Vid\":\"a\",\"Name\":\"a\",\"Valid\":true}]",
					},
				},
				response: &response{
					Value: "",
					Err:   errors.New("error"),
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
				SvSendContent(tt.mocks.request.Address, tt.mocks.request.Cc).
				Return(tt.mocks.response.Value, tt.mocks.response.Err)

			uc := &UseCase{ToService: mockService}

			got, err := uc.SendFilesToWs(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.SendFilesToWs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UseCase.SendFilesToWs() = %v, want %v", got, tt.want)
			}
		})
	}
}
