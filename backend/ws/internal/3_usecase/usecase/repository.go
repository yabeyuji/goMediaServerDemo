package usecase

import (
	"bytes"

	"ws/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("ws", "usecase")
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
		SvFileUpload(address, fileName string, fileBody *bytes.Buffer) (string, error)

		// ws
		SvSendToAgent(agentID string, cc *shared.CommonContent)
	}

	// ToDomain ...
	ToDomain interface {
	}
)
