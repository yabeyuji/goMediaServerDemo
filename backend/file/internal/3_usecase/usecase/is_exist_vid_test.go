package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock "file/internal/3_usecase/usecase/mock_repository"
	"file/internal/4_domain/domain"
)

func TestUseCase_IsExistVid(t *testing.T) {
	type request struct {
		files *[]domain.File
		vid   string
	}
	type response struct {
		boolean bool
	}
	type mocks struct {
		request  *request
		response *response
	}

	type args struct {
		files *[]domain.File
		vid   string
	}
	tests := []struct {
		name  string
		mocks mocks
		args  args
		want  bool
	}{
		// =======================================
		{
			name: "normal1",
			args: args{
				vid: "a",
				files: &[]domain.File{
					{Vid: "a", Name: "a", Valid: true},
				},
			},
			want: true,
			mocks: mocks{
				request: &request{
					vid: "a",
					files: &[]domain.File{
						{Vid: "a", Name: "a", Valid: true},
					},
				},
				response: &response{
					boolean: true,
				},
			},
		},
		// =======================================
		{
			name: "normal2",
			args: args{
				vid: "b",
				files: &[]domain.File{
					{Vid: "a", Name: "a", Valid: true},
				},
			},
			want: false,
			mocks: mocks{
				request: &request{
					vid: "b",
					files: &[]domain.File{
						{Vid: "a", Name: "a", Valid: true},
					},
				},
				response: &response{
					boolean: false,
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
				IsExistVid(tt.mocks.request.files, tt.mocks.request.vid).
				Return(tt.mocks.response.boolean)

			uc := &UseCase{ToDomain: mockDomain}

			if got := uc.IsExistVid(tt.args.files, tt.args.vid); got != tt.want {
				t.Errorf("UseCase.IsExistVid() = %v, want %v", got, tt.want)
			}
		})
	}
}
