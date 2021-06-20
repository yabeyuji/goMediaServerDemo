package usecase

import (
	mock "device/internal/3_usecase/usecase/mock_repository"
	"device/internal/4_domain/domain"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestUseCase_InfoInitialize(t *testing.T) {
	type responseBed struct {
		room *domain.Room
	}

	type responseLiving struct {
		room *domain.Room
	}

	type mocks struct {
		responseBed    *responseBed
		responseLiving *responseLiving
	}

	tests := []struct {
		name  string
		mocks mocks
		want  map[string]*domain.Room
	}{
		{
			name: "normal",
			want: map[string]*domain.Room{
				"bed": {
					Aircon: &domain.Aircon{
						Status:          "airconStop",
						WarmTemperature: 23,
						CoolTemperature: 18,
					},
					Light: &domain.Light{
						Status: "lightPower",
					},
				},
				"living": {
					Aircon: &domain.Aircon{
						Status:          "airconStop",
						WarmTemperature: 23,
						CoolTemperature: 18,
					},
					Light: &domain.Light{
						Status: "lightPower",
					},
				},
			},

			mocks: mocks{
				responseBed: &responseBed{
					room: &domain.Room{
						Aircon: &domain.Aircon{
							Status:          "airconStop",
							WarmTemperature: 23,
							CoolTemperature: 18,
						},
						Light: &domain.Light{
							Status: "lightPower",
						},
					},
				},
				responseLiving: &responseLiving{
					room: &domain.Room{
						Aircon: &domain.Aircon{
							Status:          "airconStop",
							WarmTemperature: 23,
							CoolTemperature: 18,
						},
						Light: &domain.Light{
							Status: "lightPower",
						},
					},
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
				GetDefaltValue().
				Return(tt.mocks.responseBed.room)

			mockDomain.EXPECT().
				GetDefaltValue().
				Return(tt.mocks.responseLiving.room)

			uc := &UseCase{ToDomain: mockDomain}

			if got := uc.InfoInitialize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UseCase.InfoInitialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
