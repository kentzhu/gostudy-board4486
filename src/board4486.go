package main

import (
	"board4486/httpserver"
	"board4486/router"
	"log"
)

func main() {
	log.Println("Board4486 start ...")
	server := httpserver.CreatServer()
	server.BindRouteRule(router.Binder)
	server.Run()
}
