package usecase

import (
	"testing"
	mock "ws/internal/3_usecase/usecase/mock_repository"
	"ws/pkg/shared"

	"github.com/golang/mock/gomock"
)

func TestUseCase_PassOtherApp(t *testing.T) {
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
		address string
		cc      *shared.CommonContent
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
				address: shared.GRPCAddressMedia,
				cc: &shared.CommonContent{
					Room:   shared.DataRoomCommon,
					Object: shared.DataObjectVlc,
					Key:    shared.DataKeyPlayList,
					Value:  "aaaaa",
				},
			},
			want:    "ok",
			wantErr: false,

			mocks: mocks{
				request: &request{
					address: shared.GRPCAddressMedia,
					cc: &shared.CommonContent{
						Room:   shared.DataRoomCommon,
						Object: shared.DataObjectVlc,
						Key:    shared.DataKeyPlayList,
						Value:  "aaaaa",
					},
				},
				response: &response{
					value: "ok",
					err:   nil,
				},
			},
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

			got, err := uc.PassOtherApp(tt.args.address, tt.args.cc)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.PassOtherApp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UseCase.PassOtherApp() = %v, want %v", got, tt.want)
			}
		})
	}
}
