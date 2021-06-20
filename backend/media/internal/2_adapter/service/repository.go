package service

import "media/pkg/shared"

type (
	// ToGrpcOut ...
	ToGrpcOut interface {
		IsSendContent(address string, cc *shared.CommonContent) (string, error)
		IsReceiveContent(address, funcName string) (string, error)
	}

	// ToMediaOrder ...
	ToMediaOrder interface {
		IsChangeStatus(status string)
		IsChangeVlcProgress(value float32)
		IsSwapPlayList(playListString string)
	}
)
