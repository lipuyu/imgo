package model

import (
	"fmt"
	"stouch_server/conf"
)

func init(){
	if err := conf.Orm.Sync2(new(Picture)); err != nil {
		fmt.Println(err)
	}
}
