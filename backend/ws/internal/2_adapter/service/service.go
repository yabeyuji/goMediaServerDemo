package service

import (
	"bytes"

	"ws/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("ws", "adapter:service")
}

// Service ...
type Service struct {
	ToGrpcOut ToGrpcOut
	ToWsOrder ToWsOrder
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

// SvFileUpload ...
func (sv *Service) SvFileUpload(address, fileName string, fileBody *bytes.Buffer) (string, error) {
	msg, err := sv.ToGrpcOut.IsFileUpload(address, fileName, fileBody)
	if err != nil {
		myErr.Logging(err, address, fileName)
		return "", err
	}

	return msg, nil
}

// SvSendToAgent ...
func (sv *Service) SvSendToAgent(agentID string, cc *shared.CommonContent) {
	sv.ToWsOrder.IsSendToAgent(agentID, cc)

	return
}
