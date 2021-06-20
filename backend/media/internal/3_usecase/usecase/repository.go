package usecase

import (
	"media/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("media", "usecase")
}

type (
	// UseCase ...
	UseCase struct {
		ToDomain  ToDomain
		ToService ToService
	}

	// ToService ...
	ToService interface {
		// grpcOut
		SvSendContent(address string, cc *shared.CommonContent) (string, error)
		SvReceiveContent(address, funcName string) (string, error)

		// media
		SvChangeStatus(status string)
		SvChangeVlcProgress(value float32)
		SvSwapPlayList(playListString string)
	}

	// ToDomain ...
	ToDomain interface {
		ParseKey(key string) error
		StringToFloat32(valueString string) (float32, error)
	}
)
