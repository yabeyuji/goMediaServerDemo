package usecase

import (
	mock "media/internal/3_usecase/usecase/mock_repository"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUseCase_SwapPlayList(t *testing.T) {
	type request struct {
		playlist string
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
		playlist string
	}
	tests := []struct {
		name  string
		mocks mocks
		args  args
	}{
		{
			name: "normal1",
			args: args{
				playlist: "aaa,bbb",
			},
			mocks: mocks{
				request: &request{
					playlist: "aaa,bbb",
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
				SvSwapPlayList(tt.mocks.request.playlist)

			uc := &UseCase{ToService: mockService}
			uc.SwapPlayList(tt.args.playlist)
		})
	}
}
