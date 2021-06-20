package wsapp

import (
	"errors"
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/pborman/uuid"

	"ws/internal/1_infrastructure/ws/wschannel"
	"ws/pkg/shared"
)

// WebSocket ...
func (wd *WsApp) WebSocket(c echo.Context) error {
	var err error
	var upgrader = websocket.Upgrader{}

	webSocket, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		myErr.Logging(err)
		return err
	}

	id := uuid.NewUUID()
	agent := new(Agent)
	agent.Socket = webSocket
	agent.ID = id.String()

	wd.Mutex.Lock()
	wd.Agents[agent.ID] = agent
	wd.Mutex.Unlock()

	go wd.Controller.InitialInfo(agent.ID)
	wd.ReceiveFromAgent(agent.ID)

	return nil
}

// ReceiveFromAgent ...
func (wd *WsApp) ReceiveFromAgent(agentID string) {
	log.Println("------------------------------ ")
	log.Println("start web socket")
	for {
		cc := &shared.CommonContent{}
		err := wd.Agents[agentID].Socket.ReadJSON(cc)

		if err != nil {
			myErr.Logging(err, agentID)
			wd.Disconnect(agentID)
			return
		}

		switch cc.Object {
		case shared.DataObjectVlc, shared.DataObjectFile, shared.DeviceContain(cc.Object):
			// 他のappへ渡す
			wd.Controller.PassOtherApp(cc)
		default:
			err = errors.New("not found object")
		}

		if err != nil {
			myErr.Logging(err, cc.Object)
		}

	}
}

// SendToAgents ....
func (wd *WsApp) SendToAgents() {
	for {
		content := <-wschannel.Cc
		switch content.AgentID {
		case "":
			// クライアントの数だけループ
			for _, agent := range wd.Agents {
				wd.sendToAgent(agent.ID, content.Content)
			}
		default:
			wd.sendToAgent(content.AgentID, content.Content)
		}
	}
}

// Disconnect ...
func (wd *WsApp) Disconnect(agentID string) {
	wd.Mutex.Lock()
	delete(wd.Agents, agentID)
	wd.Mutex.Unlock()

	cc := &shared.CommonContent{
		Room:   shared.DataRoomCommon,
		Object: shared.DataObjectClient,
		Key:    shared.DataKeyLeave,
		Value:  agentID,
	}

	wd.sendToAgent("", cc)
}

func (wd *WsApp) sendToAgent(agentID string, cc *shared.CommonContent) {
	err := wd.Agents[agentID].Socket.WriteJSON(cc)
	if err != nil {
		myErr.Logging(err, agentID, cc)
	}
}
