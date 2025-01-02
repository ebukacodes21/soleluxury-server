package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	db "github.com/ebukacodes21/soleluxury-server/db/sqlc"
	"github.com/ebukacodes21/soleluxury-server/servers"
	"github.com/ebukacodes21/soleluxury-server/utils"
	"github.com/ebukacodes21/soleluxury-server/worker"
	"github.com/hibiken/asynq"
	"golang.org/x/sync/errgroup"
)

var signals = []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}

	repository := db.NewSoleluxuryRepository(conn)

	ctx, stop := signal.NotifyContext(context.Background(), signals...)
	defer stop()

	group, ctx := errgroup.WithContext(ctx)

	opt := asynq.RedisClientOpt{
		Addr: config.REDISServerAddr,
	}
	td := worker.NewTaskDistributor(opt)
	tp := worker.NewTaskProcessor(opt, repository)

	servers.RunMigration(config.MigrationURL, config.DBSource)
	servers.RunGrpcServer(group, ctx, repository, config, td, tp)
	servers.RunGrpcGateway(group, ctx, repository, config, td, tp)
	servers.RunTasks(group, ctx, opt, repository)
	worker.SetupScheduler(ctx, repository)

	err = group.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
