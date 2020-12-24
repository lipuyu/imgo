package conf

import (
	"github.com/kataras/iris"
	"os"
)

var Config = iris.Configuration{}

func init(){
	env := "test"
	if len(os.Args) >= 2 {
		env = os.Args[1]
	}
	Config = iris.YAML("./conf/source/" + env + ".yml")
	loadAll(Config)
}

func loadAll(c iris.Configuration) {
	loadRedis(c)
	loadCache()
	loadOrm(c)
	loadClient(c)
}
