package usecase

import "ws/pkg/shared"

// SendContentToAgents ...
func (uc *UseCase) SendContentToAgents(cc *shared.CommonContent) {
	uc.ToService.SvSendToAgent("", cc)
	return
}
