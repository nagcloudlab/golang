package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"go-http-microservices/accounts"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const dbsource = "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"

func main() {
	var httpAddr = flag.String("http", ":8181", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")
	var db *sql.DB
	{
		var err error
		db, err = sql.Open("postgres", dbsource)
		if err != nil {
			err := level.Error(logger).Log("exit", err)
			if err != nil {
				return
			}
			os.Exit(-1)
		}
	}

	flag.Parse()
	ctx := context.Background()
	var srv accounts.Service
	{
		repository := accounts.NewRepo(db, logger)
		srv = accounts.NewService(repository, logger)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := accounts.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := accounts.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()
	level.Error(logger).Log("exit", <-errs)

}
