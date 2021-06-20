package grpcout

import (
	"bytes"

	cpb "github.com/YujiYabe/goMediaServerDemo/backend/commonpb"

	"ws/internal/2_adapter/service"
	"ws/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("ws", "infrastructure:grpcout")
}

type grpcOut struct{}

// NewToGrpcOut ...
func NewToGrpcOut() service.ToGrpcOut {
	got := new(grpcOut)
	return got
}

// IsSendContent ...
func (out *grpcOut) IsSendContent(address string, cc *shared.CommonContent) (string, error) {
	var req = &cpb.SendContentRequest{
		Room:   cc.Room,
		Object: cc.Object,
		Key:    cc.Key,
		Value:  cc.Value,
	}

	res, err := shared.SendContent(address, req)
	if err != nil {
		myErr.Logging(err, address, req)
		return "", err
	}

	return res.String(), nil
}

// IsReceiveContent ...
func (out *grpcOut) IsReceiveContent(address, funcName string) (string, error) {
	res, err := shared.ReceiveContent(address, funcName)
	if err != nil {
		myErr.Logging(err, address, funcName)
		return "", err
	}

	return res.GetValue(), nil
}

// IsFileUpload ...
func (out *grpcOut) IsFileUpload(address, fileName string, fileBody *bytes.Buffer) (string, error) {
	res, err := shared.PostFile(address, fileName, fileBody)
	if err != nil {
		myErr.Logging(err, fileName)
		return "", err
	}

	return res.String(), nil
}
