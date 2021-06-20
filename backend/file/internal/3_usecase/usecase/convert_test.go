package usecase

import (
	mock "file/internal/3_usecase/usecase/mock_repository"
	"file/pkg/shared"
	"path/filepath"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUseCase_Convert(t *testing.T) {
	type domainRequest struct {
		pathName string
		fileName string
	}

	type domainResponse struct {
		fullPath string
	}

	type domain1Mocks struct {
		domainRequest  *domainRequest
		domainResponse *domainResponse
	}

	type domain2Mocks struct {
		domainRequest  *domainRequest
		domainResponse *domainResponse
	}

	type serviceRequest struct {
		Command []string
	}

	type serviceResponse struct {
		Value []byte
		Err   error
	}

	type serviceMocks struct {
		serviceRequest  *serviceRequest
		serviceResponse *serviceResponse
	}

	type args struct {
		vid string
	}
	tests := []struct {
		name         string
		domain1Mocks domain1Mocks
		domain2Mocks domain2Mocks
		serviceMocks serviceMocks
		args         args
		wantErr      bool
	}{
		{
			name: "normal1",
			args: args{
				vid: "ll",
			},
			domain1Mocks: domain1Mocks{
				domainRequest: &domainRequest{
					pathName: shared.TempFilePath,
					fileName: "ll" + ".mp4",
				},
				domainResponse: &domainResponse{
					fullPath: filepath.Join(shared.TempFilePath, "ll"+".mp4"),
				},
			},

			domain2Mocks: domain2Mocks{
				domainRequest: &domainRequest{
					pathName: shared.TempFilePath,
					fileName: "ll" + ".gif",
				},
				domainResponse: &domainResponse{
					fullPath: filepath.Join(shared.TempFilePath, "ll"+".gif"),
				},
			},

			serviceMocks: serviceMocks{
				serviceRequest: &serviceRequest{
					Command: []string{
						"ffmpeg",
						"-i", filepath.Join(shared.TempFilePath, "ll"+".mp4"),
						"-t", "20",
						"-vf", "scale=320:-1",
						"-r", "10",
						"-y", filepath.Join(shared.TempFilePath, "ll"+".gif"),
					},
				},
				serviceResponse: &serviceResponse{
					Value: []byte{},
					Err:   nil,
				},
			},
			wantErr: false,
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
				FilepathJoin(tt.domain1Mocks.domainRequest.pathName, tt.domain1Mocks.domainRequest.fileName).
				Return(tt.domain1Mocks.domainResponse.fullPath)

			mockDomain.EXPECT().
				FilepathJoin(tt.domain2Mocks.domainRequest.pathName, tt.domain2Mocks.domainRequest.fileName).
				Return(tt.domain2Mocks.domainResponse.fullPath)

			mockService.EXPECT().
				SvExecCommand(tt.serviceMocks.serviceRequest.Command).
				Return(tt.serviceMocks.serviceResponse.Value, tt.serviceMocks.serviceResponse.Err)

			uc := &UseCase{
				ToDomain:  mockDomain,
				ToService: mockService,
			}

			if err := uc.Convert(tt.args.vid); (err != nil) != tt.wantErr {
				t.Errorf("UseCase.Convert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
