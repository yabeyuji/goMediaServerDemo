package grpcout

import (
	cpb "github.com/YujiYabe/goMediaServerDemo/backend/commonpb"

	"file/internal/2_adapter/service"
	"file/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("file", "infrastructure:grpcout")
}

// grpcOut ...
type grpcOut struct{}

// NewToGrpcOut ...
func NewToGrpcOut() service.ToGrpcOut {
	got := new(grpcOut)

	return got
}

// IsGetFiles ...
func (out *grpcOut) IsGetFiles() (string, error) {
	res, err := shared.ReceiveContent(shared.GRPCAddressFile, shared.TargetGetFiles)
	if err != nil {
		myErr.Logging(err)
		return "", err
	}

	return res.GetValue(), nil
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
