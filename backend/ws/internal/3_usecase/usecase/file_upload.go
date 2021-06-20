package usecase

import (
	"bytes"
	"ws/pkg/shared"
)

// FileUpload ...
func (uc *UseCase) FileUpload(fileName string, fileBody *bytes.Buffer) (string, error) {
	msg, err := uc.ToService.SvFileUpload(shared.GRPCAddressFile, fileName, fileBody)
	if err != nil {
		myErr.Logging(err, shared.GRPCAddressFile, fileName)
		return "", err
	}

	return msg, nil
}
