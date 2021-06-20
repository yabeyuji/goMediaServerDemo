package usecase

import (
	"ws/pkg/shared"
)

// PassOtherApp ...
func (uc *UseCase) PassOtherApp(address string, cc *shared.CommonContent) (string, error) {
	// 対象appのgrpcへ送信
	msg, err := uc.ToService.SvSendContent(address, cc)
	if err != nil {
		myErr.Logging(err, address, cc)
		return "", err
	}

	return msg, nil
}
