package tasks

import (
	"context"

	"github.com/pkg/errors"
	"github.com/scukonick/eta/repos/structs"
)

func (r *Repo) Requeue(ctx context.Context, task structs.Task) error {
	err := r.ch.Nack(task.Tag, false, true)
	if err != nil {
		return errors.Wrap(err, "failed to nack task")
	}

	return nil
}
