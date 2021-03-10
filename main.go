package main

import (
	"github.com/Thief.git/controller"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()

	routerPart := app.Party("/")
	mvc.New(routerPart).Handle(&controller.Health{})

	if err := app.Run(iris.Addr(":8080")); err != nil {
		panic(err)
	}
}
