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

	servers.RunGrpcServer(group, ctx, repository)
	servers.RunMigration(config.MigrationURL, config.DBSource)

	err = group.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
