package protocol

import (
	"github.com/gorilla/websocket"
)

type WsServer struct {
	Upgrader websocket.Upgrader `json:"upgrader"`
}

type WsConnection struct {
	ID         int64          `json:"id"`
	UserName   string         `json:"user_name"`
	targetName string         `json:"target_name"`
	conn       websocket.Conn `json:"conn"`
}
