package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock "media/internal/3_usecase/usecase/mock_repository"
)

func TestUseCase_ChangeStatus(t *testing.T) {
	type request struct {
		status string
	}

	type mocks struct {
		request *request
	}

	type args struct {
		status string
	}
	tests := []struct {
		name  string
		args  args
		mocks mocks
	}{
		{
			name: "normal1",
			args: args{
				status: "stop",
			},
			mocks: mocks{
				request: &request{
					status: "stop",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock.NewMockToService(ctrl)

			mockService.EXPECT().
				SvChangeStatus(tt.mocks.request.status)

			uc := &UseCase{ToService: mockService}

			uc.ChangeStatus(tt.args.status)
		})
	}
}
