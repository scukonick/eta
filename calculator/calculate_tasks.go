package calculator

import (
	"context"
	"math"
	"sync"

	"github.com/scukonick/eta/logger"

	"github.com/scukonick/eta/repos/structs"
)

func (s *Service) calculateTasks(ctx context.Context, in <-chan structs.Task) <-chan structs.Task {
	out := make(chan structs.Task)

	go func() {
		defer close(out)
		wg := &sync.WaitGroup{}

		for t := range in {
			wg.Add(1)
			go s.calculateTask(ctx, out, wg, t)
		}

		wg.Wait()
	}()

	return out
}

func (s *Service) calculateTask(ctx context.Context, out chan<- structs.Task, wg *sync.WaitGroup, task structs.Task) {
	defer wg.Done()

	cars, err := s.locator.GetCarsNearPoint(ctx, task.Lat, task.Lng)
	if err != nil {
		logger.FromContext(ctx).WithError(err).
			Error("failed to get cars")
		err = s.tasksRepo.Requeue(ctx, task)
		if err != nil {
			logger.FromContext(ctx).WithError(err).
				Error("failed to requeue task")
		}
		return
	}

	etas, err := s.timer.GetETAs(ctx, cars)
	if err != nil {
		logger.FromContext(ctx).WithError(err).
			Error("failed to get etas")
		err = s.tasksRepo.Requeue(ctx, task)
		if err != nil {
			logger.FromContext(ctx).WithError(err).
				Error("failed to requeue task")
		}
		return
	}

	minETA := math.MaxInt64
	for _, eta := range etas {
		if eta <= minETA {
			minETA = eta
		}
	}

	task.ETA = minETA

	out <- task
}
