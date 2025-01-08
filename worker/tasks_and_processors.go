package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	send_register_mail = "task:send_register_mail"
)

type RegisterMailPayload struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (td *TaskDistributor) DistributeTaskRegisterMail(ctx context.Context, payload *RegisterMailPayload, opts ...asynq.Option) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload %w", err)
	}
	task := asynq.NewTask(send_register_mail, []byte(data), opts...)

	info, err := td.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("failed to queue task")
	}

	log.Info().
		Str("type", task.Type()).
		Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retries", info.MaxRetry).
		Msg("message enqueued")
	return nil
}

func (tp *TaskProcessor) ProcessSendRegisterMail(ctx context.Context, task *asynq.Task) error {
	var payload RegisterMailPayload

	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("unable to unmarshal JSON: %v", err)
	}

	// user, err := tp.repository.GetUser(ctx, payload.Email)
	// if err != nil {
	// 	return fmt.Errorf("failed to get user %v ", err)
	// }

	// log.Print(user)
	// send email
	return nil
}
