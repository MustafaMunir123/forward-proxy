package main

import (
	"forward_proxy/proxy"
	"log"
)

func main() {
	var port = 8080
	log.Println("start listening on", port)
	proxy.Listen(port)
}
