package main

import (
	"github.com/alexflint/go-arg"
	"github.com/fpawel/wasmhello/internal/server"
)

// Args аргументы процесса из переменных окружения
type Args struct {
	Port string `arg:"env:PORT" default:"8881" help:"The port this service will be run on"`
}

func main() {
	var args Args
	arg.MustParse(&args)
	server.Run(args.Port)
}
