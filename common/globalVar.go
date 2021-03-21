package common

import (
	"github.com/Thief.git/protocol"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
)

var dbIns *gorm.DB
var wsServerIns *protocol.WsServer
var redisIns redis.Conn

func GetDB() *gorm.DB {
	return dbIns
}

func SetDB(db *gorm.DB) {
	dbIns = db
}

func SetWsServer(wsServer *protocol.WsServer) {
	wsServerIns = wsServer
}

func GetWsServer() *protocol.WsServer {
	return wsServerIns
}

func SetRedisIns(conn redis.Conn) {
	redisIns = conn
}

func GetRedisIns() redis.Conn {
	return redisIns
}
