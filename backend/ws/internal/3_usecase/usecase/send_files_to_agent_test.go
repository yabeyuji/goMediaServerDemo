package usecase

import (
	"testing"
	mock "ws/internal/3_usecase/usecase/mock_repository"
	"ws/pkg/shared"

	"github.com/golang/mock/gomock"
)

func TestUseCase_SendFilesToAgent(t *testing.T) {
	type requestSvReceiveContent struct {
		address  string
		funcName string
	}
	type responseSvReceiveContent struct {
		value string
		err   error
	}
	type requestSvSendToAgent struct {
		agentID string
		cc      *shared.CommonContent
	}

	type mocks struct {
		requestSvReceiveContent  *requestSvReceiveContent
		responseSvReceiveContent *responseSvReceiveContent
		requestSvSendToAgent     *requestSvSendToAgent
	}
	type args struct {
		agentID string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mocks   mocks
	}{
		// =======================================
		{
			name: "normal1",
			args: args{
				agentID: "aaa",
			},
			wantErr: false,

			mocks: mocks{
				requestSvReceiveContent: &requestSvReceiveContent{
					address:  shared.GRPCAddressFile,
					funcName: shared.TargetGetFiles,
				},
				responseSvReceiveContent: &responseSvReceiveContent{
					value: "resultSvReceiveContent",
					err:   nil,
				},
				requestSvSendToAgent: &requestSvSendToAgent{
					agentID: "aaa",
					cc: &shared.CommonContent{
						Room:   shared.DataRoomCommon,
						Object: shared.DataObjectFile,
						Key:    shared.DataKeyFiles,
						Value:  "resultSvReceiveContent",
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
				SvReceiveContent(tt.mocks.requestSvReceiveContent.address, tt.mocks.requestSvReceiveContent.funcName).
				Return(tt.mocks.responseSvReceiveContent.value, tt.mocks.responseSvReceiveContent.err)

			mockService.EXPECT().
				SvSendToAgent(tt.mocks.requestSvSendToAgent.agentID, tt.mocks.requestSvSendToAgent.cc)

			uc := &UseCase{ToService: mockService}
			if err := uc.SendFilesToAgent(tt.args.agentID); (err != nil) != tt.wantErr {
				t.Errorf("UseCase.SendFilesToAgent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
