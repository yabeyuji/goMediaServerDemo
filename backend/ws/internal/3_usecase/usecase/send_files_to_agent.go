package usecase

import (
	"ws/pkg/shared"
)

// SendFilesToAgent ...
func (uc *UseCase) SendFilesToAgent(agentID string) error {
	value, err := uc.ToService.SvReceiveContent(shared.GRPCAddressFile, shared.TargetGetFiles)
	if err != nil {
		myErr.Logging(err, shared.GRPCAddressFile, shared.TargetGetFiles)
		return err
	}

	cc := &shared.CommonContent{
		Room:   shared.DataRoomCommon,
		Object: shared.DataObjectFile,
		Key:    shared.DataKeyFiles,
		Value:  value,
	}

	uc.ToService.SvSendToAgent(agentID, cc)
	return nil
}
