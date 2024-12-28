package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type Distributor interface {
	DistributeTaskRegisterMail(ctx context.Context, payload *RegisterMailPayload, opts ...asynq.Option) error
}

type TaskDistributor struct {
	client *asynq.Client
}

func NewTaskDistributor(options asynq.RedisClientOpt) Distributor {
	client := asynq.NewClient(options)
	return &TaskDistributor{client: client}
}
