package dao

import (
	. "AppCommon/db"
	. "AppCommon/object"
)

func SelectCommonUser(userName string) *CommonUser {
	commonUser := CommonUser{}
	stmt, e := DB.Prepare("select `user_name`, mobile,password,token from `common_user` where `user_name`=?")
	CheckError(e)
	row := stmt.QueryRow(userName)
	row.Scan(&commonUser.Name, &commonUser.Mobile, &commonUser.PassWord, &commonUser.Token)
	return &commonUser
}

func InsertCommonUser(account, password string) int {
	stmt, e := DB.Prepare("insert `common_user` (`user_name`,`password`) values (?,?)")
	CheckError(e)
	rs, e := stmt.Exec(account, password)
	CheckError(e)
	rs.LastInsertId()
	result, _ := rs.RowsAffected()
	return int(result)
}
