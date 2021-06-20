package usecase

import (
	mock "file/internal/3_usecase/usecase/mock_repository"
	"file/internal/4_domain/domain"
	"file/pkg/shared"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUseCase_ReadFilesFromJSON(t *testing.T) {
	type requestSv struct {
		filePath string
	}

	type responseSv struct {
		raw []byte
		err error
	}

	type requestDm struct {
		raw   []byte
		files []domain.File
	}

	type responseDm struct {
		err error
	}

	type mocks struct {
		requestSv  *requestSv
		responseSv *responseSv
		requestDm  *requestDm
		responseDm *responseDm
	}

	type args struct {
		filePath string
	}
	var files []domain.File

	tests := []struct {
		name    string
		args    args
		want    *[]domain.File
		wantErr bool
		mocks   mocks
	}{
		// =======================================
		{
			name: "normal",
			args: args{
				filePath: shared.JSONPath,
			},
			want: &[]domain.File{
				{Vid: "a", Name: "a", Valid: true},
			},
			wantErr: false,

			mocks: mocks{
				requestSv: &requestSv{
					filePath: shared.JSONPath,
				},
				responseSv: &responseSv{
					raw: []byte{91, 123, 34, 86, 105, 100, 34, 58, 34, 97, 34, 44, 34, 78, 97, 109, 101, 34, 58, 34, 97, 34, 44, 34, 86, 97, 108, 105, 100, 34, 58, 116, 114, 117, 101, 125, 93},
					err: nil,
				},
				requestDm: &requestDm{
					raw:   []byte{91, 123, 34, 86, 105, 100, 34, 58, 34, 97, 34, 44, 34, 78, 97, 109, 101, 34, 58, 34, 97, 34, 44, 34, 86, 97, 108, 105, 100, 34, 58, 116, 114, 117, 101, 125, 93},
					files: files,
				},
				responseDm: &responseDm{
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

			mockService := mock.NewMockToService(ctrl)
			mockDomain := mock.NewMockToDomain(ctrl)
			mockService.EXPECT().
				SvReadJSONFile(tt.mocks.requestSv.filePath).
				Return(tt.mocks.responseSv.raw, tt.mocks.responseSv.err)

			mockDomain.EXPECT().
				JSONUnmarshal(tt.mocks.requestDm.raw, &tt.mocks.requestDm.files).
				Return(tt.mocks.responseDm.err)

			uc := &UseCase{
				ToService: mockService,
				ToDomain:  mockDomain,
			}

			_, err := uc.ReadFilesFromJSON(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.ReadFilesFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// TODO: ポインタを返却する処理
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("UseCase.ReadFilesFromJSON() = %v, want %v", got, tt.want)
			// }
		})
	}
}
