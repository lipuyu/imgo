package conf

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
)

var Orm *xorm.Engine

func loadOrm(c iris.Configuration){
	if Orm == nil {
		var err error
		Orm, err = xorm.NewEngine(
			"mysql",
			c.Other["DataSourceName"].(string),
			)
		if err != nil {
			fmt.Print(err)
		} else {
			iris.RegisterOnInterrupt(func() {
				err = Orm.Close()
			})
		}
	}
}
