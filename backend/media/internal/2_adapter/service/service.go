package service

import "media/pkg/shared"

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("device", "adapter:service")
}

// Service ...
type Service struct {
	ToGrpcOut    ToGrpcOut
	ToMediaOrder ToMediaOrder
}

// SvSendContent ...
func (sv *Service) SvSendContent(address string, cc *shared.CommonContent) (string, error) {
	msg, err := sv.ToGrpcOut.IsSendContent(address, cc)
	if err != nil {
		myErr.Logging(err, address, cc)
		return "", err
	}

	return msg, nil
}

// SvReceiveContent ...
func (sv *Service) SvReceiveContent(address, funcName string) (string, error) {
	msg, err := sv.ToGrpcOut.IsReceiveContent(address, funcName)
	if err != nil {
		myErr.Logging(err, address, funcName)
		return "", err
	}

	return msg, nil
}

// SvChangeStatus ...
func (sv *Service) SvChangeStatus(status string) {
	sv.ToMediaOrder.IsChangeStatus(status)

	return
}

// SvChangeVlcProgress ...
func (sv *Service) SvChangeVlcProgress(value float32) {
	sv.ToMediaOrder.IsChangeVlcProgress(value)

	return
}

// SvSwapPlayList ...
func (sv *Service) SvSwapPlayList(playListString string) {
	sv.ToMediaOrder.IsSwapPlayList(playListString)

	return
}
