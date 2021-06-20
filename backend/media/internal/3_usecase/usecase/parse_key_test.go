package usecase

import (
	"errors"
	"testing"

	mock "media/internal/3_usecase/usecase/mock_repository"
	"media/pkg/shared"

	"github.com/golang/mock/gomock"
)

func TestUseCase_ParseKey(t *testing.T) {
	type request struct {
		key string
	}

	type response struct {
		err error
	}

	type mocks struct {
		request  *request
		response *response
	}

	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mocks   mocks
	}{

		{
			name: "normal",
			args: args{
				key: shared.DataKeyStatus,
			},
			wantErr: false,
			mocks: mocks{
				request: &request{
					key: shared.DataKeyStatus,
				},
				response: &response{
					err: nil,
				},
			},
		},
		{
			name: "normal",
			args: args{
				key: shared.DataKeyProgress,
			},
			wantErr: false,
			mocks: mocks{
				request: &request{
					key: shared.DataKeyProgress,
				},
				response: &response{
					err: nil,
				},
			},
		},
		{
			name: "normal",
			args: args{
				key: shared.DataKeyPlayList,
			},
			wantErr: false,
			mocks: mocks{
				request: &request{
					key: shared.DataKeyPlayList,
				},
				response: &response{
					err: nil,
				},
			},
		},

		{
			name: "error",
			args: args{
				key: "dummy",
			},
			wantErr: true,
			mocks: mocks{
				request: &request{
					key: "dummy",
				},
				response: &response{
					err: errors.New("not found key"),
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			MockToDomain := mock.NewMockToDomain(ctrl)

			MockToDomain.EXPECT().
				ParseKey(tt.mocks.request.key).
				Return(tt.mocks.response.err)

			uc := &UseCase{ToDomain: MockToDomain}

			if err := uc.ParseKey(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("UseCase.ParseKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
