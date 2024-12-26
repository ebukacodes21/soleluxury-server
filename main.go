package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ebukacodes21/soleluxury-server/servers"
	"golang.org/x/sync/errgroup"
)

var signals = []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), signals...)
	defer stop()

	group, ctx := errgroup.WithContext(ctx)

	servers.RunGrpcServer(group, ctx)

	err := group.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
