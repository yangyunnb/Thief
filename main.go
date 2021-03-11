package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	if err := app.Run(iris.Addr(":8080"), iris.WithoutInterruptHandler); err != nil {
		panic(err)
	}
	<-doneChan
}
