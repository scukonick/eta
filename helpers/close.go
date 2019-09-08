package helpers

import (
	"context"
	"io"
	"runtime/debug"

	"github.com/scukonick/eta/logger"
)

func Close(ctx context.Context, closer io.Closer) {
	err := closer.Close()
	if err != nil {
		logger.FromContext(ctx).WithError(err).
			WithField("stack", string(debug.Stack())).
			Error("failed to close")
	}
}
