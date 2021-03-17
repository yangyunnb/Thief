package middleware

import "github.com/kataras/iris"

func RequestInFilter(ctx iris.Context) {
	ctx.Next()
}
