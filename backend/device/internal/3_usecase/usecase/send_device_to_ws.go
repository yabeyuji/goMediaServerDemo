package usecase

import "device/pkg/shared"

// SendDeviceToWs ...
func (uc *UseCase) SendDeviceToWs(value string) (string, error) {
	var cc = &shared.CommonContent{
		Room:   shared.DataRoomCommon,
		Object: shared.DataObjectDevice,
		Key:    shared.DataKeyDevices,
		Value:  value,
	}

	msg, err := uc.ToService.SvSendContent(shared.GRPCAddressWs, cc)
	if err != nil {
		myErr.Logging(err, shared.GRPCAddressWs, cc)
		return "", err
	}

	return msg, nil
}
