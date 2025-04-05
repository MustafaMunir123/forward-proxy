package main

import (
	"forward_proxy/proxy"
	"log"
)

func main() {
	// proxy.LoadEnv() TODO: add support for connection over TLS

	var port = 8080
	log.Println("start listening on", port)
	proxy.Listen(port)
}
