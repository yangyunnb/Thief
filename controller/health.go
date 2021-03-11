package controller

import "github.com/kataras/iris"

type Health struct {
}

func (ctl *Health) GetPing(ctx iris.Context) {
	ctx.JSON(map[string]string{"message:": "pong"})
}
