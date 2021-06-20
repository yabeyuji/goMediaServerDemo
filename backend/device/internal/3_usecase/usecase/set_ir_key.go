package usecase

import (
	"fmt"

	"device/pkg/shared"
)

// SetIrKey ...
func (uc *UseCase) SetIrKey(cc *shared.CommonContent) string {
	irKey := cc.Value
	if cc.Key != shared.DataKeyStatus {
		irKey = fmt.Sprint(cc.Key, cc.Value)
	}

	return irKey
}
