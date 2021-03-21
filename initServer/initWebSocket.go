package initserver

import (
	"net/http"
	"time"

	websocketLog "github.com/Thief.git/logic/websocket"

	"github.com/Thief.git/common"

	"github.com/Thief.git/protocol"

	"github.com/gorilla/websocket"
)

func InitWebSocket() {
	// 1.websocket server
	upgrader := websocket.Upgrader{
		HandshakeTimeout: time.Duration(Conf.WebSocketConfig.HandshakeTimeout) * time.Second,
		ReadBufferSize:   Conf.WebSocketConfig.ReadBufferSize,
		WriteBufferSize:  Conf.WebSocketConfig.WriteBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	wsServer := &protocol.WsServer{
		Upgrader: upgrader,
	}
	common.SetWsServer(wsServer)

	// 2.通信集线器
	websocketLog.SetHub(&websocketLog.Hub{
		Broadcast:        make(chan []byte, 2),
		Clients:          make(map[*websocketLog.Client]bool),
		RegisterClient:   make(chan *websocketLog.Client),
		UnRegisterClient: make(chan *websocketLog.Client),
	})
}
