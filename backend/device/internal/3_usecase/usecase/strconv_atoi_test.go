package usecase

import (
	mock "device/internal/3_usecase/usecase/mock_repository"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUseCase_StrConvAtoi(t *testing.T) {
	type request struct {
		valueString string
	}
	type response struct {
		valueInt int
		err      error
	}
	type mocks struct {
		request  *request
		response *response
	}

	type args struct {
		valueString string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
		mocks   mocks
	}{
		{
			name: "normal",
			args: args{
				valueString: "12",
			},
			want:    12,
			wantErr: false,
			mocks: mocks{
				request: &request{
					valueString: "12",
				},
				response: &response{
					valueInt: 12,
					err:      nil,
				},
			},
		},
		{
			name: "error",
			args: args{
				valueString: "a",
			},
			want:    0,
			wantErr: true,
			mocks: mocks{
				request: &request{
					valueString: "a",
				},
				response: &response{
					valueInt: 0,
					err:      errors.New("error"),
				},
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockDomain := mock.NewMockToDomain(ctrl)

			mockDomain.EXPECT().
				StrConvAtoi(tt.mocks.request.valueString).
				Return(tt.mocks.response.valueInt, tt.mocks.response.err)

			uc := &UseCase{
				ToDomain: mockDomain,
			}
			got, err := uc.StrConvAtoi(tt.args.valueString)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.StrConvAtoi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UseCase.StrConvAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
