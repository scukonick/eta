package tasks

import "github.com/streadway/amqp"

type Repo struct {
	queueName string
	ch        *amqp.Channel
}

func NewRepo(ch *amqp.Channel) *Repo {
	return &Repo{
		queueName: "tasks",
		ch:        ch,
	}
}
