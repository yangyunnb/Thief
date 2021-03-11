package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/Thief.git/controller"
	"github.com/kataras/iris/mvc"

	"github.com/kataras/iris/macro"

	"github.com/Thief.git/initserver"

	"github.com/kataras/iris"
)

const ServerRemainTime = 5

func main() {
	app := iris.New()
	doneChan := make(chan bool, 1)
	go shutDownServer(app, doneChan)
	initserver.InitMiddleWare(app)

	healthPart := app.Party("/")
	mvc.New(healthPart).Handle(&controller.Health{})

	macro.Int.RegisterFunc("min", func(minValue int) func(string) bool {
		return func(paramValue string) bool {
			n, err := strconv.Atoi(paramValue)
			if err != nil {
				return false
			}

			if n <= minValue {
				return false
			}
			return true
		}
	})

	userPart := app.Party("/user")
	userPart.Handle("GET", "/getUser/{id:int min(10) else 504}", func(ctx iris.Context) {
		ctx.JSON(ctx.Params().Get("id"))
	})
	app.Use()

	if err := app.Run(iris.Addr(fmt.Sprintf(":%d", initserver.Conf.Server.Port)), iris.WithoutInterruptHandler, iris.WithConfiguration(iris.Configuration{
		DisablePathCorrection: initserver.Conf.Server.DisablePathCorrection,
		EnablePathEscape:      initserver.Conf.Server.EnablePathEscape,
		FireMethodNotAllowed:  initserver.Conf.Server.FireMethodNotAllowed,
		Charset:               initserver.Conf.Server.Charset,
	})); err != nil {
		panic(err)
	}
	<-doneChan
}

func shutDownServer(app *iris.Application, doneChan chan bool) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	<-signalChan
	ctx, cancel := context.WithTimeout(context.Background(), ServerRemainTime*time.Second)
	defer cancel()

	<-ctx.Done()
	if err := app.Shutdown(ctx); err != nil {
		app.Logger().Error(err.Error())
	}
	close(doneChan)
}
