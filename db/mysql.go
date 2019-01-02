package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	driverName = "mysql"
	source     = "aaa:123456@tcp(127.0.0.1:3306)/test?charset=utf8"
)

var DB *sql.DB

func init() {
	var e error
	DB, e = sql.Open(driverName, source)
	checkError(e)
}

func checkError(e error) {
	if e != nil {
		log.Print("db connect error...")
		panic(e)
	}
}

func CheckError(e error) {
	if e != nil {
		log.Print("sql error...")
		panic(e)
	}
}
