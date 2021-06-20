package service

import "device/pkg/shared"

type (
	// ToGrpcOut ...
	ToGrpcOut interface {
		IsSendContent(address string, cc *shared.CommonContent) (string, error)
	}

	// ToDevice ...
	ToDevice interface {
		IsSendIRData(room string, irKey string) error
	}
)
