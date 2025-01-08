package worker

import (
	"context"

	db "github.com/ebukacodes21/soleluxury-server/db"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	Critical = "critical"
	Default  = "default"
)

type Processor interface {
	Start() error
	Shutdown()
	ProcessSendRegisterMail(ctx context.Context, task *asynq.Task) error
}

type TaskProcessor struct {
	server     *asynq.Server
	repository db.DatabaseContract
}

func NewTaskProcessor(options asynq.RedisConnOpt, repository db.DatabaseContract) Processor {
	server := asynq.NewServer(options, asynq.Config{
		Queues: map[string]int{
			Critical: 10,
			Default:  5,
		},
		ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
			log.Error().
				Err(err).
				Str("task_type", task.Type()).
				Bytes("payload", task.Payload()).
				Msg("task failed")
		}),
		Logger: NewLogger(),
	})

	return &TaskProcessor{
		server:     server,
		repository: repository,
	}
}

func (tp *TaskProcessor) Start() error {
	mux := asynq.NewServeMux()

	mux.HandleFunc(send_register_mail, tp.ProcessSendRegisterMail)

	return tp.server.Start(mux)
}

func (tp *TaskProcessor) Shutdown() {
	tp.server.Shutdown()
}
