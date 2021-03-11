package controller

import (
	"github.com/kataras/iris"
)

type User struct {
}

func (ctl *User) GetBy(ctx iris.Context, id int) {
	ctx.JSON(map[string]string{"id": ctx.Params().Get("id")})
}
