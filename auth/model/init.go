package model

import (
	"fmt"
	"stouch_server/conf"
)

func init(){
	if err := conf.Orm.Sync2(new(User), new(Token), new(VerificationCode)); err != nil {
		fmt.Println(err)
	}
}
