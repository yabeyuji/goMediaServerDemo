package service

import (
	"os"

	"file/pkg/shared"
)

type (
	// ToGrpcOut ...
	ToGrpcOut interface {
		IsSendContent(address string, cc *shared.CommonContent) (string, error)
	}

	// ToFile ...
	ToFile interface {
		IsExecCommand(value []string) ([]byte, error)
		IsUploadFile(filePath string, chunks *[]byte) error
		IsWriteJSONFile(filePath string, data []byte, perm os.FileMode) error
		IsReadJSONFile(filePath string) ([]byte, error)
	}
)
