package tasks

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/scukonick/eta/repos/structs"
	"github.com/streadway/amqp"
)

func (r *Repo) Store(ctx context.Context, task structs.Task) error {
	data, err := json.Marshal(&task)
	if err != nil {
		return errors.Wrap(err, "failed to marshal message")
	}

	err = r.ch.Publish(
		"",          // exchange
		r.queueName, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        data,
		})

	if err != nil {
		return errors.Wrap(err, "failed to store task in queue")
	}

	return nil
}
