package usecase

import (
	"errors"
	mock "media/internal/3_usecase/usecase/mock_repository"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUseCase_ChangeProgress(t *testing.T) {
	type requestDomain struct {
		valueString string
	}
	type responseDomain struct {
		valueFloat32 float32
		err          error
	}

	type requestService struct {
		valueFloat32 float32
	}

	type mocks struct {
		requestDomain  *requestDomain
		responseDomain *responseDomain
		requestService *requestService
	}

	type args struct {
		value string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		mocks   mocks
	}{
		{
			name: "normal",
			args: args{
				value: "0.5",
			},
			wantErr: false,
			mocks: mocks{
				requestDomain: &requestDomain{
					valueString: "0.5",
				},
				responseDomain: &responseDomain{
					valueFloat32: 0.5,
					err:          nil,
				},
				requestService: &requestService{
					valueFloat32: 0.5,
				},
			},
		},
		{
			name: "error",
			args: args{
				value: "a",
			},
			wantErr: true,
			mocks: mocks{
				requestDomain: &requestDomain{
					valueString: "a",
				},
				responseDomain: &responseDomain{
					valueFloat32: 0,
					err:          errors.New("error"),
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
				StringToFloat32(tt.mocks.requestDomain.valueString).
				Return(tt.mocks.responseDomain.valueFloat32, tt.mocks.responseDomain.err)

			if tt.mocks.requestService != nil {
				mockService.EXPECT().
					SvChangeVlcProgress(tt.mocks.requestService.valueFloat32)
			}

			uc := &UseCase{
				ToDomain:  mockDomain,
				ToService: mockService,
			}

			if err := uc.ChangeProgress(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("UseCase.ChangeProgress() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
