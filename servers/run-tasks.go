package servers

import (
	"context"
	"log"

	db "github.com/ebukacodes21/soleluxury-server/db"
	"github.com/ebukacodes21/soleluxury-server/worker"
	"github.com/hibiken/asynq"
	"golang.org/x/sync/errgroup"
)

func RunTasks(group *errgroup.Group, ctx context.Context, options asynq.RedisConnOpt, repository db.DatabaseContract) {
	processor := worker.NewTaskProcessor(options, repository)
	err := processor.Start()
	if err != nil {
		log.Fatal("unable to start processor")
	}

	group.Go(func() error {
		<-ctx.Done()
		log.Print("gracefully shutting down task processor...")

		processor.Shutdown()
		log.Print("task processor shutdown.. goodbye")

		return nil
	})
}
