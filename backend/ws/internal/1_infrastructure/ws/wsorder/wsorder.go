package wsorder

import (
	"ws/internal/1_infrastructure/ws/wschannel"
	"ws/internal/2_adapter/service"
	"ws/pkg/shared"
)

type wsOrder struct{}

// NewToWsOrder ...
func NewToWsOrder() service.ToWsOrder {
	wo := new(wsOrder)

	return wo
}

// IsSendToAgent ...
func (order *wsOrder) IsSendToAgent(agentID string, cc *shared.CommonContent) {
	contentWithAgent := &wschannel.ContentWithAgent{
		AgentID: agentID,
		Content: cc,
	}
	wschannel.Cc <- *contentWithAgent

	return
}
