package usecase

import (
	mock "file/internal/3_usecase/usecase/mock_repository"
	"file/internal/4_domain/domain"
	"file/pkg/shared"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUseCase_WriteFilesIntoJSON(t *testing.T) {
	type requestDm struct {
		files []domain.File
	}

	type responseDm struct {
		raw []byte
		err error
	}

	type requestSv struct {
		filePath string
		data     []byte
		perm     os.FileMode
	}

	type responseSv struct {
		err error
	}

	type mocks struct {
		requestDm  *requestDm
		responseDm *responseDm
		requestSv  *requestSv
		responseSv *responseSv
	}

	type args struct {
		files *[]domain.File
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mocks   mocks
	}{
		// =======================================
		{
			name: "normal",
			args: args{
				files: &[]domain.File{
					{Vid: "a", Name: "a", Valid: true},
				},
			},
			wantErr: false,

			mocks: mocks{
				requestDm: &requestDm{
					// raw:   []byte{91, 123, 34, 86, 105, 100, 34, 58, 34, 97, 34, 44, 34, 78, 97, 109, 101, 34, 58, 34, 97, 34, 44, 34, 86, 97, 108, 105, 100, 34, 58, 116, 114, 117, 101, 125, 93},
					files: []domain.File{
						{Vid: "a", Name: "a", Valid: true},
					},
				},
				responseDm: &responseDm{
					raw: []byte{91, 123, 34, 86, 105, 100, 34, 58, 34, 97, 34, 44, 34, 78, 97, 109, 101, 34, 58, 34, 97, 34, 44, 34, 86, 97, 108, 105, 100, 34, 58, 116, 114, 117, 101, 125, 93},
					err: nil,
				},
				requestSv: &requestSv{
					filePath: shared.JSONPath,
					data:     []byte{91, 123, 34, 86, 105, 100, 34, 58, 34, 97, 34, 44, 34, 78, 97, 109, 101, 34, 58, 34, 97, 34, 44, 34, 86, 97, 108, 105, 100, 34, 58, 116, 114, 117, 101, 125, 93},
					perm:     os.ModePerm,
				},
				responseSv: &responseSv{
					err: nil,
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
			mockService := mock.NewMockToService(ctrl)
			mockDomain.EXPECT().
				JSONMarshalIndent(&tt.mocks.requestDm.files).
				Return(tt.mocks.responseDm.raw, tt.mocks.responseDm.err)

			mockService.EXPECT().
				SvWriteJSONFile(tt.mocks.requestSv.filePath, tt.mocks.requestSv.data, tt.mocks.requestSv.perm).
				Return(tt.mocks.responseSv.err)

			uc := &UseCase{
				ToDomain:  mockDomain,
				ToService: mockService,
			}

			if err := uc.WriteFilesIntoJSON(tt.args.files); (err != nil) != tt.wantErr {
				t.Errorf("UseCase.WriteFilesIntoJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
