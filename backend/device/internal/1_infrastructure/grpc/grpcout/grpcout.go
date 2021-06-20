package grpcout

import (
	cpb "github.com/YujiYabe/goMediaServerDemo/backend/commonpb"

	"device/internal/2_adapter/service"
	"device/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("file", "infrastructure:grpcout")
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
