package main

import (
	"github.com/fpawel/wasmhello/internal/server"
	"os"
)

func main() {
	port := defaultPort
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	server.Run(port)
}

const defaultPort = "8001"
