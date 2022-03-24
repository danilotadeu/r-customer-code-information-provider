package main

import (
	serverInit "github.com/danilotadeu/r-customer-code-information-provider/server"
)

var (
	server serverInit.Server
)

func init() {
	server = serverInit.New()
}

func main() {
	server.Start()
}
