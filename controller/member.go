package controller

import (
	"AppCommon/db"
	"AppCommon/object"
	"AppCommon/request"
	"AppCommon/utils"
	"github.com/kataras/iris"
	"strings"
)

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
	user := db.QueryUser(userName[0])
	token := utils.UUID()
	if strings.EqualFold(password[0], user.PassWord) {
		user.Token = token
	} else {
		cxt.JSON(request.Error("username or password error"))
		return
	}
	//insert token in redis

	//return login info
	user1 := object.CommonUser{Name: user.Name, Token: user.Token}
	cxt.JSON(request.Response(user1))
}
