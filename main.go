package main

import (
	_ "./db"
	"./router"
)

func main() {
	router.ServerStart()
}
