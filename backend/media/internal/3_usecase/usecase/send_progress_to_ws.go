package usecase

import "media/pkg/shared"

// SendProgressToWs ...
func (uc *UseCase) SendProgressToWs(value string) (string, error) {
	var cc = &shared.CommonContent{
		Room:   shared.DataRoomCommon,
		Object: shared.DataObjectVlc,
		Key:    shared.DataKeyProgress,
		Value:  value,
	}

	value, err := uc.ToService.SvSendContent(shared.GRPCAddressWs, cc)
	if err != nil {
		myErr.Logging(err, shared.GRPCAddressWs, cc)
		return "", err
	}

	return value, nil
}
