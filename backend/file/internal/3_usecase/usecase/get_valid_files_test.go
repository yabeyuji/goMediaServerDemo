package usecase

import (
	"testing"

	mock "file/internal/3_usecase/usecase/mock_repository"
	"file/internal/4_domain/domain"

	"github.com/golang/mock/gomock"
)

func TestUseCase_GetValidFiles(t *testing.T) {
	type request struct {
		files *[]domain.File
	}
	type response struct {
		fileString string
	}
	type mocks struct {
		request  *request
		response *response
	}
	type args struct {
		files *[]domain.File
	}
	tests := []struct {
		name  string
		mocks mocks
		args  args
		want  string
	}{
		// =======================================
		{
			name: "normal",
			args: args{
				files: &[]domain.File{
					{Vid: "a", Name: "a", Valid: true},
				},
			},
			want: "a",

			mocks: mocks{
				request: &request{
					files: &[]domain.File{
						{Vid: "a", Name: "a", Valid: true},
					},
				},
				response: &response{
					fileString: "a",
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
				GetValidFiles(tt.mocks.request.files).
				Return(tt.mocks.response.fileString)

			uc := &UseCase{ToDomain: mockDomain}

			if got := uc.GetValidFiles(tt.args.files); got != tt.want {
				t.Errorf("UseCase.GetValidFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
