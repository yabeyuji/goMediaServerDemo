package usecase

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"

	mock "ws/internal/3_usecase/usecase/mock_repository"
	"ws/pkg/shared"
)

func TestUseCase_FileUpload(t *testing.T) {
	type request struct {
		address  string
		fileName string
		fileBody *bytes.Buffer
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
		fileName string
		fileBody *bytes.Buffer
	}

	bytesBuffer := bytes.NewBuffer([]byte("test"))

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		mocks   mocks
	}{

		// =======================================
		{
			name: "normal1",
			args: args{
				fileName: "a",
				fileBody: bytesBuffer,
			},
			want:    "ok",
			wantErr: false,
			mocks: mocks{
				request: &request{
					address:  shared.GRPCAddressFile,
					fileName: "a",
					fileBody: bytesBuffer,
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
				SvFileUpload(tt.mocks.request.address, tt.mocks.request.fileName, tt.mocks.request.fileBody).
				Return(tt.mocks.response.value, tt.mocks.response.err)

			uc := &UseCase{ToService: mockService}

			got, err := uc.FileUpload(tt.args.fileName, tt.args.fileBody)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.FileUpload() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UseCase.FileUpload() = %v, want %v", got, tt.want)
			}
		})
	}
}
