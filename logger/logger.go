package logger

import (
	"context"

	"github.com/sirupsen/logrus"
)

type key string

var loggerKey key

func init() {
	loggerKey = "logger"
}

func FromContext(ctx context.Context) *logrus.Entry {
	if entry, ok := ctx.Value(loggerKey).(*logrus.Entry); ok {
		return entry
	}

	return logrus.NewEntry(logrus.New())
}

func ToContext(ctx context.Context, entry *logrus.Entry) context.Context {
	ctx = context.WithValue(ctx, loggerKey, entry)
	return ctx
}
