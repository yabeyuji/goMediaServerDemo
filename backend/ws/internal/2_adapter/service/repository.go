package service

import (
	"bytes"

	"ws/pkg/shared"
)

type (
	// ToGrpcOut ...
	ToGrpcOut interface {
		IsSendContent(address string, cc *shared.CommonContent) (string, error)
		IsReceiveContent(address, funcName string) (string, error)
		IsFileUpload(address, fileName string, fileBody *bytes.Buffer) (string, error)
	}

	// ToWsOrder ...
	ToWsOrder interface {
		IsSendToAgent(agentID string, cc *shared.CommonContent)
	}
)
