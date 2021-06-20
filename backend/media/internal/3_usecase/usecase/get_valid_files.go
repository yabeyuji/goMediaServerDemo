package usecase

import "media/pkg/shared"

// GetValidFiles ...
func (uc *UseCase) GetValidFiles() (string, error) {
	value, err := uc.ToService.SvReceiveContent(shared.GRPCAddressFile, shared.TargetGetValidFiles)
	if err != nil {
		myErr.Logging(err, shared.GRPCAddressDevice, shared.TargetGetRooms)
		return "", err
	}

	return value, nil
}
