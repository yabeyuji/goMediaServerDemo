package service

import (
	"os"

	"file/pkg/shared"
)

var myErr *shared.MyErr

func init() {
	myErr = shared.NewMyErr("file", "adapter:service")
}

// Service ...
type Service struct {
	ToGrpcOut ToGrpcOut
	ToFile    ToFile
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

// SvExecCommand ...
func (sv *Service) SvExecCommand(value []string) ([]byte, error) {
	raw, err := sv.ToFile.IsExecCommand(value)
	if err != nil {
		myErr.Logging(err, value)
		return nil, err
	}

	return raw, nil
}

// SvUploadFile ...
func (sv *Service) SvUploadFile(filePath string, chunks *[]byte) error {
	err := sv.ToFile.IsUploadFile(filePath, chunks)
	if err != nil {
		myErr.Logging(err, filePath)
		return err
	}

	return nil
}

// SvWriteJSONFile ...
func (sv *Service) SvWriteJSONFile(filePath string, data []byte, perm os.FileMode) error {
	err := sv.ToFile.IsWriteJSONFile(filePath, data, perm)
	if err != nil {
		myErr.Logging(err, filePath, data)
		return err
	}

	return nil
}

// SvReadJSONFile ...
func (sv *Service) SvReadJSONFile(filePath string) ([]byte, error) {
	raw, err := sv.ToFile.IsReadJSONFile(filePath)
	if err != nil {
		myErr.Logging(err, filePath)
		return nil, err
	}

	return raw, nil
}
