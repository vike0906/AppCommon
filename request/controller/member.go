package controller

import (
	"AppCommon/db"
	"AppCommon/object"
	"AppCommon/request"
	"AppCommon/request/service"
	"AppCommon/utils"
	"github.com/kataras/iris"
	"strings"
)

//login
func Login(cxt iris.Context) {
	r := cxt.Request()
	r.ParseForm()
	userName := r.Form["userName"]
	password := r.Form["password"]
	if len(userName) == 0 {
		cxt.JSON(request.Error("please input your login name"))
		return
	} else if len(password) == 0 {
		cxt.JSON(request.Error("please input your password"))
		return
	}
	//query user from mysql db by userName
	user := service.GetUserByName(userName[0])

	token := utils.UUID()
	if strings.EqualFold(password[0], user.PassWord) {
		user.Token = token
	} else {
		cxt.JSON(request.Error("username or password error"))
		return
	}

	//insert token in redis
	db.SetJSONEX(token, user, 60)

	//return login info
	user1 := object.CommonUser{Name: user.Name, Token: user.Token}
	cxt.JSON(request.Response(user1))
}

//register
func Register(cxt iris.Context) {
	r := cxt.Request()
	r.ParseForm()
	account := r.Form["account"]
	password := r.Form["password"]
	//valication_code

	//add account
	re := service.AddUser(account[0], password[0])
	if re != 1 {
		cxt.JSON(request.Error("register error"))
		return
	}
	//return register info
}
