package usecase

import "file/pkg/shared"

// SendFilesToMedia ...
func (uc *UseCase) SendFilesToMedia(value string) (string, error) {
	var cc = &shared.CommonContent{
		Room:   shared.DataRoomCommon,
		Object: shared.DataObjectVlc,
		Key:    shared.DataKeyPlayList,
		Value:  value,
	}

	msg, err := uc.ToService.SvSendContent(shared.GRPCAddressMedia, cc)
	if err != nil {
		myErr.Logging(err, shared.GRPCAddressMedia, cc)
		return "", err
	}

	return msg, nil
}
