package main

import (
	"flag"
	"webServer/server"
)

func main() {
	var port string
	//sprcify port number
	flag.StringVar(&port, "p", "8000", "the port you want to listen tos")
	flag.Parse()
	server.Serve(port)
}
