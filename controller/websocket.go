package controller

import (
	"fmt"

	"github.com/Thief.git/common"

	"github.com/Thief.git/logic/websocket"
	"github.com/kataras/iris"
)

type WebsocketController struct {
}

func (ctl WebsocketController) GetUpgrade(ctx iris.Context) {
	// 1.升级链接到ws
	conn, err := common.GetWsServer().Upgrader.Upgrade(ctx.ResponseWriter(), ctx.Request(), ctx.ResponseWriter().Header())
	if err != nil {
		ctx.JSON(map[string]interface{}{"err": err.Error()})
	}

	// 2.更新集群IP集合
	if err = websocket.HandleClusterIPSets(ctx, conn); err != nil {
		ctx.JSON(map[string]interface{}{"err": err.Error()})
	}

	// 3.业务逻辑
	client := &websocket.Client{
		Hub:  websocket.GetHub(),
		Send: make(chan []byte),
		Conn: conn,
	}
	websocket.GetHub().RegisterClient <- client
	go client.ReadPump()
	go client.WritePump()
}

func (ctl *WebsocketController) GetSyncInfo(ctx iris.Context) {
	fmt.Println("before GetSyncInfo:", "I am other hub message")
	websocket.GetHub().Broadcast <- []byte("I am other hub message")
	fmt.Println("after GetSyncInfo:", "I am other hub message")

	ctx.JSON(map[string]string{"message:": "GetSyncInfo success"})
}
