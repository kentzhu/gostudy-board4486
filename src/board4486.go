package main

import (
	"board4486/httpserver"
	"board4486/router"
	"board4486/service/config"
	"log"
)

func main() {
	log.Println("Board4486 start ...")
	server := httpserver.CreatServer()
	server.BindRouteRule(router.Binder)
	server.Run(config.Config().Http.ListenAddress)
}
