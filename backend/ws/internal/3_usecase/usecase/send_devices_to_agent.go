package usecase

import (
	"ws/pkg/shared"
)

// SendDevicesToAgent ...
func (uc *UseCase) SendDevicesToAgent(agentID string) error {
	value, err := uc.ToService.SvReceiveContent(shared.GRPCAddressDevice, shared.TargetGetRooms)
	if err != nil {
		myErr.Logging(err, shared.GRPCAddressDevice, shared.TargetGetRooms)
		return err
	}

	cc := &shared.CommonContent{
		Room:   shared.DataRoomCommon,
		Object: shared.DataObjectDevice,
		Key:    shared.DataKeyDevices,
		Value:  value,
	}

	uc.ToService.SvSendToAgent(agentID, cc)
	return nil
}
