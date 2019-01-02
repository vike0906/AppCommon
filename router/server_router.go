package router

import (
	. "AppCommon/controller"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func ServerStart() {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/", func(cxt iris.Context) {
		cxt.HTML("<h1>index</h1>")
	})
	//app.Handle("GET","/welcome", func(cxt iris.Context) {
	//	cxt.HTML("<h1>welcome</h1>")
	//}, func(cxt iris.Context) {
	//	cxt.HTML("<h1>aaa</h1>")
	//})
	//app.Get("/ping", func(cxt iris.Context) {
	//	cxt.WriteString("PONG")
	//})
	//app.Get("/hello", func(cxt iris.Context) {
	//	cxt.JSON(map[string]string{"message":"hello"})
	//})
	app.Post("/login", Login)
	app.Post("/logOut", Login)

	app.Get("/", func(cxt iris.Context) {
		cxt.Redirect("/index")
	})
	user := app.Party("/user", func(cxt iris.Context) {
		fmt.Print("are you red")

		cxt.Next()
	})
	user.Get("/hello", func(cxt iris.Context) {
		cxt.JSON(map[string]string{"aa": "11", "bb": "22"})
	})
	user.Get("/hello/{id:int}", func(cxt iris.Context) {
		cxt.WriteString(cxt.Params().Get("id"))
	})
	app.Run(iris.Addr(":8080"))
}
