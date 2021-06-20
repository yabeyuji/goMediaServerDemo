package usecase

import (
	"errors"
	mock "file/internal/3_usecase/usecase/mock_repository"
	"file/pkg/shared"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUseCase_SendFilesToMedia(t *testing.T) {
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
		mocks   mocks
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "normal1",
			args: args{
				value: "aaa, bbb, ccc",
			},
			mocks: mocks{
				request: &request{
					Address: shared.GRPCAddressMedia,
					Cc: &shared.CommonContent{
						Room:   shared.DataRoomCommon,
						Object: shared.DataObjectVlc,
						Key:    shared.DataKeyPlayList,
						Value:  "aaa, bbb, ccc",
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
				value: "aaa, bbb, ccc",
			},
			mocks: mocks{
				request: &request{
					Address: shared.GRPCAddressMedia,
					Cc: &shared.CommonContent{
						Room:   shared.DataRoomCommon,
						Object: shared.DataObjectVlc,
						Key:    shared.DataKeyPlayList,
						Value:  "aaa, bbb, ccc",
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

			got, err := uc.SendFilesToMedia(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.SendFilesToMedia() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UseCase.SendFilesToMedia() = %v, want %v", got, tt.want)
			}
		})
	}
}
