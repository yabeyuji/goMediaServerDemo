package controller

import (
	"fmt"

	"media/internal/2_adapter/service"
	"media/internal/3_usecase/usecase"
	"media/internal/4_domain/domain"
	"media/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("media", "adapter:controller")
}

type (
	// Controller ...
	Controller struct {
		UseCase usecase.UseCase
	}
)

// NewController ...
func NewController(toGrpcOut service.ToGrpcOut, toMediaOrder service.ToMediaOrder) *Controller {
	ct := &Controller{
		UseCase: usecase.UseCase{
			ToDomain: domain.NewDomain(),
			ToService: &service.Service{
				ToGrpcOut:    toGrpcOut,
				ToMediaOrder: toMediaOrder,
			},
		},
	}

	return ct
}

// BootStrap ...
func (ctrl *Controller) BootStrap() {
	playListString, err := ctrl.UseCase.GetValidFiles()
	if err != nil {
		myErr.Logging(err)
	}

	ctrl.UseCase.SwapPlayList(playListString)

	return
}

// VlcOperation ...
func (ctrl *Controller) VlcOperation(cc *shared.CommonContent) error {
	err := ctrl.UseCase.ParseKey(cc.Key)
	if err != nil {
		myErr.Logging(err, cc.Key)
		return err
	}

	switch cc.Key {
	case shared.DataKeyStatus:
		ctrl.UseCase.ChangeStatus(cc.Value)
	case shared.DataKeyProgress:
		err = ctrl.UseCase.ChangeProgress(cc.Value)
	case shared.DataKeyPlayList:
		ctrl.UseCase.SwapPlayList(cc.Value)
	}
	if err != nil {
		myErr.Logging(err, cc.Key)
		return err
	}

	return nil
}

// SwapPlayList ...
func (ctrl *Controller) SwapPlayList(playListString string) {
	ctrl.UseCase.SwapPlayList(playListString)

	return
}

// SendProgress ...
func (ctrl *Controller) SendProgress(progress float32) {
	// progressのレスポンスはログが埋まってしまうので出さない
	_, err := ctrl.UseCase.SendProgressToWs(fmt.Sprint(progress))
	if err != nil {
		myErr.Logging(err, progress)
	}

	return
}
