package usecase

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"

	mock "file/internal/3_usecase/usecase/mock_repository"
	"file/internal/4_domain/domain"
)

func TestUseCase_GetFilesString(t *testing.T) {
	type request struct {
		files *[]domain.File
	}
	type response struct {
		fileString []byte
		err        error
	}
	type mocks struct {
		request  *request
		response *response
	}

	type args struct {
		files *[]domain.File
	}
	tests := []struct {
		name    string
		mocks   mocks
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				files: &[]domain.File{
					{Vid: "a", Name: "a", Valid: true},
				},
			},
			want:    "[{\"Vid\":\"a\",\"Name\":\"a\",\"Valid\":true}]",
			wantErr: false,

			mocks: mocks{
				request: &request{
					files: &[]domain.File{
						{Vid: "a", Name: "a", Valid: true},
					},
				},
				response: &response{
					fileString: []byte{91, 123, 34, 86, 105, 100, 34, 58, 34, 97, 34, 44, 34, 78, 97, 109, 101, 34, 58, 34, 97, 34, 44, 34, 86, 97, 108, 105, 100, 34, 58, 116, 114, 117, 101, 125, 93},
					err:        nil,
				},
			},
		},
		{
			name: "error",
			args: args{
				files: &[]domain.File{
					{Vid: "a", Name: "a", Valid: true},
				},
			},
			want:    "",
			wantErr: true,

			mocks: mocks{
				request: &request{
					files: &[]domain.File{
						{Vid: "a", Name: "a", Valid: true},
					},
				},
				response: &response{
					fileString: nil,
					err:        errors.New("text string"),
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
				JSONMarshal(tt.mocks.request.files).
				Return(tt.mocks.response.fileString, tt.mocks.response.err)

			uc := &UseCase{ToDomain: mockDomain}

			got, err := uc.GetFilesString(tt.args.files)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.GetFilesString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UseCase.GetFilesString() = %v, want %v", got, tt.want)
			}
		})
	}
}
