package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/scukonick/eta/logger"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/scukonick/eta/calculator"
	"github.com/scukonick/eta/external/locator"
	"github.com/scukonick/eta/external/timer"
	"github.com/scukonick/eta/handlers"
	"github.com/scukonick/eta/repos/results"
	"github.com/scukonick/eta/repos/tasks"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func main() {
	l := logrus.New()

	amqpConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		l.WithError(err).Fatal("failed to connect to amqp queue")
	}
	amqpCh, err := amqpConn.Channel()
	if err != nil {
		l.WithError(err).Error("failed to get amqp channel")
	}

	ctx := context.Background()
	ctx = logger.ToContext(ctx, logrus.NewEntry(l))

	tasksRepo := tasks.NewRepo(amqpCh)
	err = tasksRepo.InitQueue(ctx)
	if err != nil {
		l.WithError(err).Error("failed to init tasks repo queue")
	}

	pgxConfig, err := pgx.ParseEnvLibpq()
	if err != nil {
		l.WithError(err).Fatal("failed to parse postgresql config")
	}
	sqlDB := stdlib.OpenDB(pgxConfig)
	err = sqlDB.Ping()
	if err != nil {
		l.WithError(err).Fatal("failed to ping DB")
	}

	resultsRepo := results.NewRepo(sqlx.NewDb(sqlDB, "pgx"))

	incomer := handlers.NewIncomer(tasksRepo)
	resulter := handlers.NewResulter(resultsRepo)

	mux := http.NewServeMux()
	mux.Handle("/put", incomer)
	mux.Handle("/get", resulter)

	handler := loggerMiddleware(mux, logrus.NewEntry(l))

	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}

	go func() {
		err = server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			l.WithError(err).Fatal("failed to listen and serve")
		}
	}()

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	locatorService := locator.NewService(httpClient, "dev-api.wheely.com")
	timerService := timer.NewService(httpClient, "dev-api.wheely.com")

	calc := calculator.NewService(tasksRepo, resultsRepo, locatorService, timerService)

	calcCtx, cancel := context.WithCancel(ctx)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		err := calc.RunCalculation(calcCtx, wg)
		if err != nil {
			l.WithError(err).Fatal("failed to run calculations")
		}
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	server.Shutdown(ctx)
	cancel()

	wg.Wait()
}

func loggerMiddleware(next http.Handler, l *logrus.Entry) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := logger.ToContext(r.Context(), l)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
