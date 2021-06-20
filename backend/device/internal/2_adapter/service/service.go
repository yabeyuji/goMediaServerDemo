package service

import "device/pkg/shared"

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("device", "adapter:service")
}

// Service ...
type Service struct {
	ToGrpcOut ToGrpcOut
	ToDevice  ToDevice
}

// SvSendIRData ...
func (sv *Service) SvSendIRData(room string, irKey string) error {
	err := sv.ToDevice.IsSendIRData(room, irKey)
	if err != nil {
		myErr.Logging(err, room, irKey)
		return err
	}

	return nil
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
