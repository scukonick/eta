package tasks

import (
	"context"

	"github.com/pkg/errors"
	"github.com/scukonick/eta/repos/structs"
)

func (r *Repo) Ack(ctx context.Context, task structs.Task) error {
	err := r.ch.Ack(task.Tag, false)
	if err != nil {
		return errors.Wrap(err, "failed to ack task")
	}

	return nil
}
