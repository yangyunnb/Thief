package websocket

import (
	"github.com/Thief.git/common"
	"github.com/gorilla/websocket"
	"github.com/kataras/iris"
)

// 1.存储集群IP地址
func HandleClusterIPSets(ctx iris.Context, conn *websocket.Conn) error {
	interIp, err := common.GetInterfaceAddr()
	if err != nil {
		return err
	}

	// 1.存储集群的IP全集
	if _, err := common.GetRedisIns().Do("SADD", "ipSets", interIp); err != nil {
		return err
	}

	return nil
}
