package dao

import (
	. "AppCommon/db"
	. "AppCommon/object"
)

func SelectCommonUser(userName string) *CommonUser {
	commonUser := CommonUser{}
	stmt, e := DB.Prepare("select `user_name` Name, mobile Mobile,password PassWord,token Token from `common_user` where `user_name`=?")
	CheckError(e)
	row := stmt.QueryRow(userName)
	row.Scan(&commonUser.Name, &commonUser.Mobile, &commonUser.PassWord, &commonUser.Token)
	return &commonUser
}
