package usecase

import "file/pkg/shared"

// SendFilesToWs ...
func (uc *UseCase) SendFilesToWs(value string) (string, error) {
	var cc = &shared.CommonContent{
		Room:   shared.DataRoomCommon,
		Object: shared.DataObjectFile,
		Key:    shared.DataKeyFiles,
		Value:  value,
	}

	msg, err := uc.ToService.SvSendContent(shared.GRPCAddressWs, cc)
	if err != nil {
		myErr.Logging(err, shared.GRPCAddressWs, cc)
		return "", err
	}

	return msg, nil
}
