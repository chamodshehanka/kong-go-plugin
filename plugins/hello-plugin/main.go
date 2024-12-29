package main

import (
	"hello-plugin/handlers"

	"github.com/Kong/go-pdk/server"
)

var Version = "0.2"
var Priority = 5000

func main() {
	server.StartServer(handlers.New, Version, Priority)
}
