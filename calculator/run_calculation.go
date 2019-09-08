package calculator

import (
	"context"
	"sync"

	"github.com/scukonick/eta/logger"
)

// RunCalculation does all the work - reads new tasks from queue,
// stores empty results in the DB,
// calculates ETAs,
// updates results in the DB,
// acknowledges tasks
func (s *Service) RunCalculation(ctx context.Context, wg *sync.WaitGroup) error {
	defer wg.Done()

	tasks, err := s.tasksRepo.Subscribe(ctx)
	if err != nil {
		logger.FromContext(ctx).WithError(err).
			Error("failed to run calculations")
		return err
	}

	storedTasks := s.storeEmptyResults(ctx, tasks)
	calculatedTasks := s.calculateTasks(ctx, storedTasks)
	storedCalculatedTasks := s.storeResults(ctx, calculatedTasks)
	s.ackTasks(ctx, storedCalculatedTasks)
	return nil
}
