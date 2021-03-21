package websocket

import (
	"fmt"
	"net/http"

	"github.com/gomodule/redigo/redis"

	"github.com/Thief.git/common"
)

type Hub struct {
	Clients          map[*Client]bool
	RegisterClient   chan *Client
	UnRegisterClient chan *Client
	Broadcast        chan []byte
}

var hubIns *Hub

func SetHub(hub *Hub) {
	hubIns = hub
}

func GetHub() *Hub {
	return hubIns
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.RegisterClient:
			hub.Clients[client] = true
		case client := <-hub.UnRegisterClient:
			if _, ok := hub.Clients[client]; ok {
				delete(hub.Clients, client)
				close(client.Send)
			}
		case msg := <-hub.Broadcast:
			// 1.本机集线器所有链接发送消息
			for client, _ := range hub.Clients {
				client.Send <- msg
				fmt.Println("send a message through local hub")
			}

			// 2.将消息发送到除本机外的所有集线器(分布式部署)
			ipStr, err := common.GetInterfaceAddr()
			if err != nil {
				fmt.Print(err.Error())
			}

			ipSets, err := redis.Strings(common.GetRedisIns().Do("SMEMBERS", "ipSets"))
			if err != nil {
				fmt.Print(err.Error())
			}

			for _, ip := range ipSets {
				if ip != ipStr {
					fmt.Println("send a message through other hub")
					syncMessageToOtherHub(ip)
				}
			}
		}
	}
}

// 讲消息同步到其他的中继器 GetSyncInfo
func syncMessageToOtherHub(ip string) {
	resp, err := http.Get(fmt.Sprintf("http://%s:8080/ws/sync/info", ip))
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Print(resp)
}
