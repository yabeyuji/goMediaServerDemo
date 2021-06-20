package wschannel

import (
	"ws/pkg/shared"
)

// ContentWithAgent ...
type ContentWithAgent struct {
	AgentID string
	Content *shared.CommonContent
}

// Cc ...
var Cc = make(chan ContentWithAgent)
