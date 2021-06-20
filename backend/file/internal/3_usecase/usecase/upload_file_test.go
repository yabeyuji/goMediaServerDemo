package usecase

import (
	"errors"
	mock "file/internal/3_usecase/usecase/mock_repository"
	"file/pkg/shared"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUseCase_UploadFile(t *testing.T) {
	type domainRequest struct {
		pathName string
		fileName string
	}

	type domainResponse struct {
		fullPath string
	}

	type domainMocks struct {
		domainRequest  *domainRequest
		domainResponse *domainResponse
	}

	type serviceRequest struct {
		filePath string
		chunks   *[]byte
	}

	type serviceResponse struct {
		err error
	}

	type serviceMocks struct {
		serviceRequest  *serviceRequest
		serviceResponse *serviceResponse
	}

	type args struct {
		vid    string
		chunks *[]byte
	}
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		domainMocks  domainMocks
		serviceMocks serviceMocks
	}{
		// =======================================
		{
			name: "normal",
			args: args{
				vid:    "ll",
				chunks: &[]byte{11, 22, 33},
			},
			wantErr: false,

			domainMocks: domainMocks{
				domainRequest: &domainRequest{
					pathName: shared.TempFilePath,
					fileName: "ll" + ".mp4",
				},
				domainResponse: &domainResponse{
					fullPath: filepath.Join(shared.TempFilePath, "ll"+".mp4"),
				},
			},

			serviceMocks: serviceMocks{
				serviceRequest: &serviceRequest{
					filePath: filepath.Join(shared.TempFilePath, "ll"+".mp4"),
					chunks:   &[]byte{11, 22, 33},
				},
				serviceResponse: &serviceResponse{
					err: nil,
				},
			},
		},
		// =======================================
		{
			name: "error",
			args: args{
				vid:    "ll",
				chunks: &[]byte{11, 22, 33},
			},
			wantErr: true,
			domainMocks: domainMocks{
				domainRequest: &domainRequest{
					pathName: shared.TempFilePath,
					fileName: "ll" + ".mp4",
				},
				domainResponse: &domainResponse{
					fullPath: filepath.Join(shared.TempFilePath, "ll"+".mp4"),
				},
			},
			serviceMocks: serviceMocks{
				serviceRequest: &serviceRequest{
					filePath: filepath.Join(shared.TempFilePath, "ll"+".mp4"),
					chunks:   &[]byte{11, 22, 33},
				},
				serviceResponse: &serviceResponse{
					err: errors.New("error"),
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
				FilepathJoin(tt.domainMocks.domainRequest.pathName, tt.domainMocks.domainRequest.fileName).
				Return(tt.domainMocks.domainResponse.fullPath)

			mockService.EXPECT().
				SvUploadFile(tt.serviceMocks.serviceRequest.filePath, tt.serviceMocks.serviceRequest.chunks).
				Return(tt.serviceMocks.serviceResponse.err)

			uc := &UseCase{
				ToDomain:  mockDomain,
				ToService: mockService,
			}

			if err := uc.UploadFile(tt.args.vid, tt.args.chunks); (err != nil) != tt.wantErr {
				t.Errorf("UseCase.UploadFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
