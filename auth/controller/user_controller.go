package controller

import (
	"github.com/kataras/iris"
	"stouch_server/auth/model"
	"stouch_server/common/er"
	"stouch_server/common/utils"
	"stouch_server/conf"
	"time"
)

type UserController struct{
	Ctx iris.Context
}

func (c *UserController) Get(ctx iris.Context) interface{}{
	/*
	user1 := model.User{}
	ctx.ReadJSON(&user1)
	fmt.Println(user1)
	*/
	user := ctx.Values().Get("user").(model.User)
	return er.Data(map[string]model.User{"user": user})
}

func (c *UserController) PostSignin() interface{} {
	jsonData := map[string]string{"username": "", "password": ""}
	c.Ctx.ReadJSON(&jsonData)
	username, _ := jsonData["username"]
	password, _ := jsonData["password"]
	user := model.User{Username: username}
	if ok, _ := conf.Orm.Get(&user); ok {
		if user.Check(password){
			token := model.Token{UserId: user.Id, Ticket: utils.GetUUID()}
			conf.Orm.Insert(token)
			return er.Data(map[string]string{"ticket": token.Ticket})
		} else {
			return er.PasswordError
		}
	} else {
		return er.UserNotExistError
	}
}

func (c *UserController) PostSignup() interface{} {
	jsonData := map[string]string{"username": "", "password": ""}
	c.Ctx.ReadJSON(&jsonData)
	username, _ := jsonData["username"]
	password, _ := jsonData["password"]
	user := &model.User{Username: username, CreatedAt: time.Now()}
	user.SetPassword(password)
	var token model.Token
	if has, _ := conf.Orm.Get(&model.User{Username: username}); has {
		return er.UserNameRepeatError
	}
	if  _, err := conf.Orm.Insert(user); err == nil {
		token = model.Token{UserId: user.Id, Ticket: utils.GetUUID()}
		conf.Orm.Insert(token)
	} else {
		conf.Logger.Error(err)
	}
	return er.Data(map[string]string{"ticket": token.Ticket})
}

func (c *UserController) GetBy(id int64) interface{}{
	user := model.User{Id: id}
	if ok, _ := conf.Orm.Get(&user); ok {
		return er.Data(map[string]model.User{"user": user})
	} else  {
		return er.UserNotExistError
	}
}

func (c *UserController) GetVerificationCode() interface{}{
	// user := c.Ctx.Values().Get("user").(model.User)
	return er.UserNotExistError
}
