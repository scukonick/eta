package repos

import (
	"context"

	"github.com/scukonick/eta/repos/structs"
)

type Tasks interface {
	Store(context.Context, structs.Task) error
	Subscribe(context.Context) (<-chan structs.Task, error)
	Requeue(context.Context, structs.Task) error
	Ack(ctx context.Context, task structs.Task) error
}

type Results interface {
	UpsertResult(context.Context, structs.Result) error
	Get(ctx context.Context, id string) (*structs.Result, error)
}
