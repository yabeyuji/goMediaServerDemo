package usecase

import (
	"testing"
	mock "ws/internal/3_usecase/usecase/mock_repository"
	"ws/pkg/shared"

	"github.com/golang/mock/gomock"
)

func TestUseCase_SendContentToAgents(t *testing.T) {
	type request struct {
		agentID string
		cc      *shared.CommonContent
	}

	type mocks struct {
		request *request
	}

	type args struct {
		cc *shared.CommonContent
	}

	tests := []struct {
		name  string
		args  args
		mocks mocks
	}{
		{
			name: "normal1",
			args: args{
				cc: &shared.CommonContent{
					Room:   shared.DataRoomCommon,
					Object: shared.DataObjectVlc,
					Key:    shared.DataKeyPlayList,
					Value:  "aaa, bbb, ccc",
				},
			},
			mocks: mocks{
				request: &request{
					agentID: "",
					cc: &shared.CommonContent{
						Room:   shared.DataRoomCommon,
						Object: shared.DataObjectVlc,
						Key:    shared.DataKeyPlayList,
						Value:  "aaa, bbb, ccc",
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
			mockService := mock.NewMockToService(ctrl)

			mockService.EXPECT().
				SvSendToAgent(tt.mocks.request.agentID, tt.mocks.request.cc)

			uc := &UseCase{ToService: mockService}
			uc.SendContentToAgents(tt.args.cc)
		})
	}
}
