package service

import (
	"AppCommon/basic"
	"AppCommon/request/dao"
)

func GetUserByName(userName string) *basic.CommonUser {
	user := dao.SelectCommonUser(userName)
	return user
}

func AddUser(account, password string) int {
	return dao.InsertCommonUser(account, password)
}
