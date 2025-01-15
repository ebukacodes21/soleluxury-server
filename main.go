package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	db "github.com/ebukacodes21/soleluxury-server/db"
	"github.com/ebukacodes21/soleluxury-server/servers"
	"github.com/ebukacodes21/soleluxury-server/utils"
	"github.com/ebukacodes21/soleluxury-server/worker"

	"github.com/hibiken/asynq"
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

	// mongo connection and a repo
	client, err := mongo.Connect(options.Client().ApplyURI(config.MongoUrl))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	repository := db.NewRepository(client, config.MongoDBName)

	opts := asynq.RedisClientOpt{
		Addr: config.REDISServerAddr,
	}
	// dedicated workers
	td := worker.NewTaskDistributor(opts)
	tp := worker.NewTaskProcessor(opts, repository)

	// servers
	servers.RunGrpcGateway(group, ctx, repository, config, td, tp)
	servers.RunGrpcServer(group, ctx, repository, config, td, tp)
	servers.RunTasks(group, ctx, opts, repository)

	// listen for cancelled contexts
	err = group.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
