package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Thief.git/middleware"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"

	"github.com/Thief.git/controller"
	"github.com/kataras/iris/mvc"

	"github.com/Thief.git/initserver"
	"github.com/kataras/iris"
)

func main() {
	// 1.新建server
	app := iris.New()
	registerMiddleWare(app)

	// 2.路由设置
	healthPart := app.Party("/")
	mvc.New(healthPart).Handle(&controller.Health{})

	moviePart := app.Party("/movie")
	mvc.New(moviePart).Handle(&controller.Movie{})

	// 3.启动server
	doneChan := make(chan bool, 1)
	go shutDownServer(app, doneChan)

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

// server优雅退出
func shutDownServer(app *iris.Application, doneChan chan bool) {
	// server延迟退出时间
	const ServerRemainTime = 5
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

// 注册服务用到的中间件
func registerMiddleWare(app *iris.Application) {
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(middleware.LoginFilter)
	app.Use(middleware.RequestInFilter)
	app.Use(middleware.RequestOutFilter)
}
