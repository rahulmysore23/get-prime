package main

import (
	"os"

	"github.com/rahulmysore23/get-prime/http_server"
)

func main() {
	app := http_server.GetCliApp()
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
