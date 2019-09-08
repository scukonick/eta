package tasks

import (
	"context"
	"encoding/json"

	"github.com/scukonick/eta/helpers"
	"github.com/scukonick/eta/logger"
	"github.com/scukonick/eta/repos/structs"
)

func (r *Repo) Subscribe(ctx context.Context) (<-chan structs.Task, error) {
	out := make(chan structs.Task)

	go func() {
		defer close(out)
		defer helpers.Close(ctx, r.ch)

		deliveries, err := r.ch.Consume(r.queueName,
			"", false, false, false, false, nil)
		if err != nil {
			logger.FromContext(ctx).WithError(err).Error("failed to consume tasks")
			return
		}

		for {
			select {
			case d := <-deliveries:
				task := &structs.Task{}
				err = json.Unmarshal(d.Body, task)
				if err != nil {
					logger.FromContext(ctx).WithError(err).Error("failed to unmarshal task")
					continue
				}
				task.Tag = d.DeliveryTag

				select {
				case out <- *task:
					continue
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return out, nil
}
