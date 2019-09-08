package calculator

import (
	"context"

	"github.com/scukonick/eta/logger"
	"github.com/scukonick/eta/repos/structs"
)

func (s *Service) ackTasks(ctx context.Context, in <-chan structs.Task) {
	for t := range in {
		err := s.tasksRepo.Ack(ctx, t)
		if err != nil {
			logger.FromContext(ctx).WithError(err).
				Error("failed to ack task")
		}
	}
}
