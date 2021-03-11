package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Thief.git/initserver"

	"github.com/Thief.git/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

const ServerRemainTime = 5

func main() {
	app := iris.New()
	doneChan := make(chan bool, 1)

	go func() {
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
	}()

	healthPart := app.Party("/")

	mvc.New(healthPart).Handle(&controller.Health{})
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
