package main

import (
	"github.com/rahulmysore23/get-prime/http_server"
)

func main() {
	if err := http_server.Start(); err != nil {
		panic(err)
	}
}
