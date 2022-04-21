package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/VladNF/calendar/internal/common"
	q "github.com/VladNF/calendar/internal/queue"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "/etc/scheduler/config.yaml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := NewConfigFromFile(configFile)
	log := common.NewLogger(config.Logger.Level, config.Logger.File)

	sender := &Sender{
		config: config,
		log:    log,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	log.Info("sender is running...")
	go sender.startServer(ctx, cancel)

	<-ctx.Done()
	sender.stopServer()
}

type Sender struct {
	config   Config
	log      common.Logger
	mqServer common.StartStopper
}

func (s *Sender) startServer(_ context.Context, cancel context.CancelFunc) {
	s.mqServer = q.NewConsumer(
		s.config.MQ, s.log, func(msg string) {
			s.log.Errorf("!!! ATTENTION !!! YOU HAVE AN EVENT ALERT:\n%v", msg)
		},
	)
	if err := s.mqServer.Start(); err != nil {
		s.log.Errorf("failed to start %T: %s", s, err.Error())
		cancel()
		os.Exit(1)
	}
}

func (s *Sender) stopServer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := s.mqServer.Stop(ctx); err != nil {
		s.log.Errorf("failed to stop %T: %s", s, err.Error())
	}
}
