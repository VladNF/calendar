package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/app"
	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/common"
	m "github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/models"
	q "github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/queue"
	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/storage"
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

	eventsRepo, err := storage.NewStorage(config.Storage.Kind)
	if err != nil {
		log.Fatalf("storage was not created: %v", err)
	}
	scheduler := &Scheduler{
		app:    app.New(log, eventsRepo),
		alerts: make(chan string),
		config: config,
		log:    log,
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	log.Info("scheduler is running...")
	go scheduler.startServer(ctx, cancel)

	<-ctx.Done()
	scheduler.stopServer()
}

type Scheduler struct {
	app      *app.App
	alerts   chan string
	config   Config
	log      common.Logger
	mqServer common.StartStopper
}

func (s *Scheduler) startServer(ctx context.Context, cancel context.CancelFunc) {
	s.mqServer = q.NewProducer(s.config.MQ, s.log, s.alerts)
	if err := s.mqServer.Start(); err != nil {
		s.log.Errorf("failed to start %T: %s", s, err.Error())
		cancel()
		os.Exit(1)
	}

	go s.sendAlerts(ctx)
}

func (s *Scheduler) stopServer() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	if err := s.mqServer.Stop(ctx); err != nil {
		s.log.Errorf("failed to stop %T: %s", s, err.Error())
	}
}

func (s *Scheduler) sendAlerts(ctx context.Context) {
	period := time.Duration(s.config.SchedMins) * time.Minute //nolint:durationcheck // there's no other way

	for {
		select {
		case <-ctx.Done():
			close(s.alerts)
			return
		case <-time.After(period):
			s.log.Infof("attempting to make alerts...")
			s.makeAlerts(ctx)
		}
	}
}

func (s *Scheduler) makeAlerts(ctx context.Context) {
	events, err := s.app.GetDailyAgenda(ctx, time.Now().AddDate(0, 0, -s.config.NoticeDays))
	if err != nil {
		s.log.Errorf("make alerts: %v", err)
	} else {
		s.log.Infof("queried %v event to alert of...", len(events))
	}

	for _, e := range events {
		if alert, err := json.Marshal(m.NewAlert(e)); err != nil {
			s.log.Errorf("make alerts: %v", err)
		} else {
			s.alerts <- string(alert)
		}
	}
}
