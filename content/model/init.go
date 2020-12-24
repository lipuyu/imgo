package model

import (
	"fmt"
	"stouch_server/conf"
)

func init(){
	if err := conf.Orm.Sync2(new(Topic)); err != nil {
		fmt.Println(err)
	}
}
