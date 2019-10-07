package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
	appconfCt "stouch_server/appconf/controller"
	"stouch_server/auth"
	authCt "stouch_server/auth/controller"
	"stouch_server/conf"
	contentCt "stouch_server/content/controller"
	storageCt "stouch_server/storage/controller"
	"stouch_server/websock"
	bookCt "stouch_server/websock/controller"
)

func newApp() *iris.Application {
	app := iris.New()
	app.Use(recover.New())
	app.Use(logger.New())
	app.Get("/", func(ctx iris.Context) { ctx.Redirect("/web/home.html") })
	app.StaticWeb("/web", "./static")
	mvc.New(app.Party("/appconf")).Handle(new(appconfCt.AppConfController))
	mvc.New(app.Party("/user")).Handle(new(authCt.UserController))
	mvc.New(app.Party("/content")).Handle(new(contentCt.ContentController))
	mvc.New(app.Party("/storage/token")).Handle(new(storageCt.StorageTokenController))
	mvc.New(app.Party("/storage/picture")).Handle(new(storageCt.PictureController))
	mvc.New(app.Party("/book")).Handle(new(bookCt.BookController))
	return app
}

func main() {
	app := newApp()
	conf.LoadLog(app)
	websock.SetupWebsocket(app) // websocket 服务
	app.UseGlobal(auth.Before)
	go conf.Run()
	app.Run(iris.Addr("0.0.0.0:8080"), iris.WithConfiguration(conf.Config))
}
