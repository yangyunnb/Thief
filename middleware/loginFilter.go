package middleware

import "github.com/kataras/iris"

func LoginFilter(ctx iris.Context) {
	ctx.Next()
}
