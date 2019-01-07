package service

import (
	"AppCommon/object"
	"AppCommon/request/dao"
)

func GetUserByName(userName string) *object.CommonUser {
	user := dao.SelectCommonUser(userName)
	return user
}

func AddUser(account, password string) int {
	return dao.InsertCommonUser(account, password)
}
