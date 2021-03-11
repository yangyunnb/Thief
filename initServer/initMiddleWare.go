package initserver

import (
	"fmt"
	"runtime/debug"

	"github.com/kataras/iris"
)

func InitMiddleWare(app *iris.Application) {
	app.Use(RecoverFilter)
	app.Use(LoginFilter)
	app.Use(RequestInFilter)
	app.Done(RequestOutFilter)
}

func RecoverFilter(ctx iris.Context) {
	defer func() {
		if r := recover(); r != nil {
			var err error
			switch r := r.(type) {
			case error:
				err = r
			default:
				err = fmt.Errorf("%v", r)
			}

			fmt.Println(map[string]string{
				"tag":   "recover_panic",
				"error": err.Error(),
				"stack": string(debug.Stack()),
			})
		}
	}()
	ctx.Next()
}

func LoginFilter(ctx iris.Context) {
	ctx.Next()
}

func RequestInFilter(ctx iris.Context) {
	ctx.Next()
}

func RequestOutFilter(ctx iris.Context) {
	ctx.Next()
}
