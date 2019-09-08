package tasks

import (
	"context"

	"github.com/pkg/errors"
)

func (r *Repo) InitQueue(ctx context.Context) error {
	_, err := r.ch.QueueDeclare(r.queueName, true, false, false, false, nil)
	if err != nil {
		return errors.Wrap(err, "failed to declare queue")
	}

	return nil
}
