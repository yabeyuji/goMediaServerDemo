package controller

import (
	"device/internal/2_adapter/service"
	"device/internal/3_usecase/usecase"
	"device/internal/4_domain/domain"
	"device/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("device", "adapter:controller")
}

type (
	// Controller ...
	Controller struct {
		Rooms   map[string]*domain.Room
		UseCase usecase.UseCase
	}
)

// NewController ...
func NewController(toGrpcOut service.ToGrpcOut, toDevice service.ToDevice) *Controller {
	ct := &Controller{
		UseCase: usecase.UseCase{
			ToDomain: domain.NewDomain(),
			ToService: &service.Service{
				ToGrpcOut: toGrpcOut,
				ToDevice:  toDevice,
			},
		},
	}
	ct.Rooms = make(map[string]*domain.Room)
	return ct
}

// BootStrap ...
func (ctrl *Controller) BootStrap() {
	ctrl.Rooms = ctrl.UseCase.InfoInitialize()
}

// ReceiveContent ...
func (ctrl *Controller) ReceiveContent(funcName string) string {
	var roomsString string
	var err error
	roomsString, err = ctrl.UseCase.GetRoomsString(ctrl.Rooms)
	if err != nil {
		myErr.Logging(err, ctrl.Rooms)
	}

	return roomsString
}

// SendToDevice ...
func (ctrl *Controller) SendToDevice(cc *shared.CommonContent) {
	var err error

	// 赤外線データマップ構造体のキー設定
	irKey := ctrl.UseCase.SetIrKey(cc)

	// 赤外線送信
	err = ctrl.UseCase.SendIRData(cc.Room, irKey)
	if err != nil {
		myErr.Logging(err, cc.Room, irKey)
	}

	// デバイス管理ステータスを変更
	ctrl.Rooms = ctrl.UseCase.ChangeDeviceStatus(cc, ctrl.Rooms)

	// websocketで送る必要がなければ終了
	isBroadcastStatus := ctrl.UseCase.GetBroadcastStatus(cc.Object)
	if !isBroadcastStatus {
		return
	}

	// 部屋情報をjson化
	roomsString, err := ctrl.UseCase.GetRoomsString(ctrl.Rooms)
	if err != nil {
		myErr.Logging(err, ctrl.Rooms)
	}

	// json送信
	_, err = ctrl.UseCase.SendDeviceToWs(roomsString)
	if err != nil {
		myErr.Logging(err, roomsString)
	}

	return
}
