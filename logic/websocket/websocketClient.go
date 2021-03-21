package websocket

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Hub  *Hub
	Conn *websocket.Conn
	Send chan []byte
}

func (c *Client) ReadPump() {
	defer func() {
		c.Hub.UnRegisterClient <- c
		if err := c.Conn.Close(); err != nil {
			fmt.Println("err:", err.Error())
		}
	}()

	for {
		messType, mess, err := c.Conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(messType, mess)
		c.Hub.Broadcast <- mess
	}
}

func (c *Client) WritePump() {
	ticker := time.NewTicker(1 * time.Second)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			if !ok {
				c.Conn.WriteMessage(websocket.CloseMessage, []byte("conn close"))
				break
			}
			returnMess := string(msg) + "from server"
			c.Conn.WriteMessage(websocket.TextMessage, []byte(returnMess))
		}
	}
}
