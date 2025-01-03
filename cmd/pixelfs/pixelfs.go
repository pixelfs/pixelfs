package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pixelfs/pixelfs/cmd/pixelfs/cli"
)

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		os.Exit(0)
	}()

	cli.Execute()
}
