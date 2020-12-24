package controller

import (
	"github.com/kataras/iris"
	"stouch_server/common/base"
)

type AppConfController struct{
	Ctx iris.Context
}

func (c *AppConfController) GetCdn() interface{} {
	return re.NewByData(iris.Map{"cdn": "http://airport.xiaorere.com"})
}
