package calculator

import (
	"context"
	"database/sql"
	"sync"

	"github.com/scukonick/eta/logger"
	"github.com/scukonick/eta/repos/structs"
)

func (s *Service) storeResults(ctx context.Context, in <-chan structs.Task) <-chan structs.Task {
	out := make(chan structs.Task)

	go func() {
		defer close(out)

		wg := &sync.WaitGroup{}
		for t := range in {
			wg.Add(1)
			go s.storeTaskResult(ctx, wg, out, t)
		}
		wg.Wait()
	}()

	return out
}

func (s *Service) storeTaskResult(ctx context.Context, wg *sync.WaitGroup, out chan<- structs.Task, task structs.Task) {
	defer wg.Done()
	err := s.resultsRepo.UpsertResult(ctx, structs.Result{
		ID:     task.ID,
		Status: structs.StatusProcessed,
		ETA:    sql.NullInt64{Valid: true, Int64: int64(task.ETA)},
	})

	if err != nil {
		logger.FromContext(ctx).WithError(err).Error("failed to upsert result")
		err = s.tasksRepo.Requeue(ctx, task)
		if err != nil {
			logger.FromContext(ctx).WithError(err).Error("failed to nack task")
			return
		}
	}

	out <- task
}
