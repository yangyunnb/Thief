package middleware

import "github.com/kataras/iris"

func RequestOutFilter(ctx iris.Context) {
	ctx.Next()
}
