package usecase

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock "device/internal/3_usecase/usecase/mock_repository"
	"device/internal/4_domain/domain"
)

func TestUseCase_GetRoomsString(t *testing.T) {
	type request struct {
		rooms map[string]*domain.Room
	}
	type response struct {
		raw []byte
		err error
	}
	type mocks struct {
		request  *request
		response *response
	}

	type args struct {
		rooms map[string]*domain.Room
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
		mocks   mocks
	}{
		// =======================================
		{
			name: "normal",
			args: args{
				map[string]*domain.Room{
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
			},
			want:    "{\"bed\":{\"Aircon\":{\"Status\":\"airconStop\",\"WarmTemperature\":23,\"CoolTemperature\":18},\"Light\":{\"Status\":\"lightPower\"}},\"living\":{\"Aircon\":{\"Status\":\"airconStop\",\"WarmTemperature\":23,\"CoolTemperature\":18},\"Light\":{\"Status\":\"lightPower\"}}}",
			wantErr: false,
			mocks: mocks{
				request: &request{
					rooms: map[string]*domain.Room{
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
				},
				response: &response{
					raw: []byte{123, 34, 98, 101, 100, 34, 58, 123, 34, 65, 105, 114, 99, 111, 110, 34, 58, 123, 34, 83, 116, 97, 116, 117, 115, 34, 58, 34, 97, 105, 114, 99, 111, 110, 83, 116, 111, 112, 34, 44, 34, 87, 97, 114, 109, 84, 101, 109, 112, 101, 114, 97, 116, 117, 114, 101, 34, 58, 50, 51, 44, 34, 67, 111, 111, 108, 84, 101, 109, 112, 101, 114, 97, 116, 117, 114, 101, 34, 58, 49, 56, 125, 44, 34, 76, 105, 103, 104, 116, 34, 58, 123, 34, 83, 116, 97, 116, 117, 115, 34, 58, 34, 108, 105, 103, 104, 116, 80, 111, 119, 101, 114, 34, 125, 125, 44, 34, 108, 105, 118, 105, 110, 103, 34, 58, 123, 34, 65, 105, 114, 99, 111, 110, 34, 58, 123, 34, 83, 116, 97, 116, 117, 115, 34, 58, 34, 97, 105, 114, 99, 111, 110, 83, 116, 111, 112, 34, 44, 34, 87, 97, 114, 109, 84, 101, 109, 112, 101, 114, 97, 116, 117, 114, 101, 34, 58, 50, 51, 44, 34, 67, 111, 111, 108, 84, 101, 109, 112, 101, 114, 97, 116, 117, 114, 101, 34, 58, 49, 56, 125, 44, 34, 76, 105, 103, 104, 116, 34, 58, 123, 34, 83, 116, 97, 116, 117, 115, 34, 58, 34, 108, 105, 103, 104, 116, 80, 111, 119, 101, 114, 34, 125, 125, 125},
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

			mockDomain.EXPECT().
				JSONMarshal(tt.mocks.request.rooms).
				Return(tt.mocks.response.raw, tt.mocks.response.err)

			uc := &UseCase{ToDomain: mockDomain}

			got, err := uc.GetRoomsString(tt.args.rooms)
			if (err != nil) != tt.wantErr {
				t.Errorf("UseCase.GetRoomsString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UseCase.GetRoomsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
