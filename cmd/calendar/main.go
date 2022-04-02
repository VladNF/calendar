package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/app"
	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/common"
	servergrpc "github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/server/grpc"
	serverhttp "github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/server/http"
	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/storage"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/calendar/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := NewConfig(configFile)
	log := common.NewLogger(config.Logger.Level, config.Logger.File)

	eventsRepo, err := storage.NewStorage(config.Storage.Kind)
	if err != nil {
		log.Fatalf("storage was not created: %v", err)
	}
	calendar := app.New(log, eventsRepo)
	grpcServer := servergrpc.NewServer(config.GRPC.Host, config.GRPC.Port, log, calendar)
	httpServer := serverhttp.NewServer(config.HTTP.Host, config.HTTP.Port, log, calendar)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	log.Info("calendar is running...")
	go startServer(grpcServer, log, cancel)
	go startServer(httpServer, log, cancel)

	<-ctx.Done()
	stopServer(grpcServer, log)
	stopServer(httpServer, log)
}

func startServer(s common.StartStopper, log common.Logger, cancel context.CancelFunc) {
	if err := s.Start(); err != nil {
		log.Errorf("failed to start %T: %s", s, err.Error())
		cancel()
		os.Exit(1)
	}
}

func stopServer(s common.StartStopper, log common.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := s.Stop(ctx); err != nil {
		log.Errorf("failed to stop %T: %s", s, err.Error())
	}
}
