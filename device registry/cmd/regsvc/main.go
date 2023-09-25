package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/piusalfred/igrid/logger"
	"github.com/piusalfred/registry"
	"github.com/piusalfred/registry/api"
	"github.com/piusalfred/registry/bcrypt"
	"github.com/piusalfred/registry/postgres"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)
	flag.Parse()

	var l log.Logger
	{
		l = log.NewLogfmtLogger(os.Stderr)
		l = log.With(l, "ts", log.DefaultTimestampUTC)
		l = log.With(l, "caller", log.DefaultCaller)
	}

	log, err := logger.New(os.Stderr, "info")
	if err != nil {
		panic(err)
	}

	db := connectToDB(log)
	defer db.Close()

	hasher := bcrypt.New()

	provider := registry.New()

	users := postgres.NewUserRepository(db)
	nodes := postgres.NewNodeRepository(db)
	regio := postgres.NewRegionRepository(db)

	var s registry.Service
	{
		s = registry.NewService(users, nodes, regio, hasher, log, provider)
		s = api.LoggingMiddleware(log)(s)
	}

	var h http.Handler
	{
		h = api.MakeHTTPHandler(s, l)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {

		log.Info(fmt.Sprintf("registry started on port %v", *httpAddr))
		//l.Log("transport", "HTTP", "addr", *httpAddr)
		errs <- http.ListenAndServe(*httpAddr, h)
	}()

	log.Info(fmt.Sprintf("exit due to: %v", <-errs))
}

func connectToDB(logger logger.Logger) *sql.DB {
	db, err := postgres.Connect()
	if err != nil {
		logger.Info(fmt.Sprintf("Failed to connect to postgres: %s", err))
		os.Exit(1)
	}
	return db
}
