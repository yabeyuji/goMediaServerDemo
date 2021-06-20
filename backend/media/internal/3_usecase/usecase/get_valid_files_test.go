package usecase

import (
	"errors"
	mock "media/internal/3_usecase/usecase/mock_repository"
	"media/pkg/shared"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUseCase_GetValidFiles(t *testing.T) {
	type request struct {
		address  string
		funcName string
	}

	type response struct {
		valueString string
		err         error
	}

	type mocks struct {
		request  *request
		response *response
	}

	tests := []struct {
		name    string
		want    string
		wantErr bool
		mocks   mocks
	}{
		{
			name:    "normal",
			want:    "aaa,bbb",
			wantErr: false,
			mocks: mocks{
				request: &request{
					address:  shared.GRPCAddressFile,
					funcName: shared.TargetGetValidFiles,
				},
				response: &response{
					valueString: "aaa,bbb",
					err:         nil,
				},
			},
		},
		{
			name:    "error",
			want:    "",
			wantErr: true,
			mocks: mocks{
				request: &request{
					address:  shared.GRPCAddressFile,
					funcName: shared.TargetGetValidFiles,
				},
				response: &response{
					valueString: "",
					err:         errors.New("error"),
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			// uc := &UseCase{
			// 	ToDomain:  tt.fields.ToDomain,
			// 	ToService: tt.fields.ToService,
			// }

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mockService := mock.NewMockToService(ctrl)

			mockService.EXPECT().
				SvReceiveContent(tt.mocks.request.address, tt.mocks.request.funcName).
				Return(tt.mocks.response.valueString, tt.mocks.response.err)

			uc := &UseCase{ToService: mockService}

			got, err := uc.GetValidFiles()
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.GetValidFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UseCase.GetValidFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
