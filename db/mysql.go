package db

import (
	. "AppCommon/object"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	driverName = "mysql"
	source     = "aaa:123456@tcp(127.0.0.1:3306)/test?charset=utf8"
)

var db *sql.DB

func init() {
	var e error
	db, e = sql.Open(driverName, source)
	checkError(e)
}

func QueryUser(userName string) CommonUser {
	commonUser := CommonUser{}
	stmt, e := db.Prepare("select `user_name` Name, mobile Mobile,password PassWord,token Token from `common_user` where `user_name`=?")
	checkError(e)
	row := stmt.QueryRow(userName)
	row.Scan(&commonUser.Name, &commonUser.Mobile, &commonUser.PassWord, &commonUser.Token)
	return commonUser
}

func checkError(e error) {
	if e != nil {
		log.Print("db connect error...")
		panic(e)
	}
}
