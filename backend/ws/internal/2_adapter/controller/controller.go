package controller

import (
	"bytes"
	"sync"

	"ws/internal/2_adapter/service"
	"ws/internal/3_usecase/usecase"
	"ws/internal/4_domain/domain"
	"ws/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("ws", "adapter:controller")
}

type (
	// Controller ...
	Controller struct {
		UseCase usecase.UseCase
	}
)

// NewController ...
func NewController(toGrpcOut service.ToGrpcOut, toWsOrder service.ToWsOrder) *Controller {
	ct := &Controller{
		UseCase: usecase.UseCase{
			ToDomain: domain.NewDomain(),
			ToService: &service.Service{
				ToGrpcOut: toGrpcOut,
				ToWsOrder: toWsOrder,
			},
		},
	}

	return ct
}

// InitialInfo ...
// websocket接続時の初期情報
func (ctrl *Controller) InitialInfo(agentID string) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := ctrl.UseCase.SendFilesToAgent(agentID)
		if err != nil {
			myErr.Logging(err, agentID)
		}
	}()
	go func() {
		defer wg.Done()
		err := ctrl.UseCase.SendDevicesToAgent(agentID)
		if err != nil {
			myErr.Logging(err, agentID)
		}
	}()
	wg.Wait()

	return
}

// SendContentToAgents ...
func (ctrl *Controller) SendContentToAgents(cc *shared.CommonContent) {
	// websocketのagentへの送信はエラーが返却されない
	ctrl.UseCase.SendContentToAgents(cc)

	return
}

// PassOtherApp ...
func (ctrl *Controller) PassOtherApp(cc *shared.CommonContent) {
	var address string

	switch cc.Object {
	case shared.DataObjectVlc:
		address = shared.GRPCAddressMedia
	case shared.DataObjectFile:
		address = shared.GRPCAddressFile
	case shared.DeviceContain(cc.Object):
		address = shared.GRPCAddressDevice
	}

	_, err := ctrl.UseCase.PassOtherApp(address, cc)
	if err != nil {
		myErr.Logging(err, address, cc)
	}

	return
}

// FileUpload ...
func (ctrl *Controller) FileUpload(fileName string, fileBody *bytes.Buffer) {
	_, err := ctrl.UseCase.FileUpload(fileName, fileBody)
	if err != nil {
		// バイナリはログに出力しない
		myErr.Logging(err, fileName)
	}
}
