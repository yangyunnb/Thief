package initserver

import (
	"github.com/Thief.git/common"
	"github.com/gomodule/redigo/redis"
)

func InitRedis() {
	conn, err := redis.Dial("tcp", Conf.RedisConfig.Host)
	if err != nil {
		panic(err)
	}
	common.SetRedisIns(conn)
}
