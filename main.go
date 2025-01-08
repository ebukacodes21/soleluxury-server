package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ebukacodes21/soleluxury-server/utils"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"golang.org/x/sync/errgroup"
)

var signals = []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	// create go routines
	ctx, stop := signal.NotifyContext(context.Background(), signals...)
	defer stop()
	group, ctx := errgroup.WithContext(ctx)

	// mongo connection
	client, err := mongo.Connect(options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// dbb.NewMongoRepository()

	err = group.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
