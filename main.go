package main

import (
	_ "AppCommon/db"
	"AppCommon/router"
)

func main() {
	router.ServerStart()
}
