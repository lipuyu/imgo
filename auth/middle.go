package auth

import (
	"fmt"
	"github.com/kataras/iris"
	"stouch_server/auth/model"
	"stouch_server/common/er"
	"stouch_server/conf"
	"strings"
)

func Before(ctx iris.Context) {
	/*
		shareInformation := "this is a sharable information between handlers"
		requestPath := ctx.Path()
		println("Before the indexHandler or contactHandler: " + requestPath)
		ctx.Values().Set("info", shareInformation)
	*/
	path := ctx.GetCurrentRoute().ResolvePath()
	ticket := ctx.GetHeader("ticket")
	app := ctx.GetHeader("app")

	// websocket 信息进行特殊处理
	if strings.HasPrefix(path, "/websocket") {
		ticket, _ = ctx.URLParams()["ticket"]
		app, _ = ctx.URLParams()["app"]
	}
	fmt.Println(ticket, app, "end")

	// 读取user信息
	var user model.User
	if ticket == "" {
		user = model.User{}
	} else {
		token := model.Token{Ticket: ticket}
		if g, err := conf.Orm.Get(&token); err == nil && g {
			user = model.User{Id: token.UserId}
			if c, err := conf.Orm.Get(&user); err != nil || !c {
				user = model.User{}
			}
		} else {
			user = model.User{}
		}
	}
	ctx.Values().Set("user", user)

	// URL 拦截
	if (app == "stouch" && user.Id != 0) || path == "/storage/token" || path == "/storage/picture/editor" ||
		path == "/" || ctx.Method() == "OPTIONS" || strings.HasPrefix(path, "/user/sign") ||
		strings.HasPrefix(path, "/web/") {
		ctx.Next()
	} else {
		ctx.JSON(er.AppError)
	}
}
