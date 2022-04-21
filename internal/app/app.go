package app

import (
	"context"
	"time"

	"github.com/VladNF/calendar/internal/common"
	m "github.com/VladNF/calendar/internal/models"
)

type App struct {
	logger common.Logger
	repo   m.EventsRepo
}

func New(logger common.Logger, repo m.EventsRepo) *App {
	return &App{logger: logger, repo: repo}
}

func (app *App) CreateEvent(
	ctx context.Context, id, title string, start, end time.Time, owner string,
) (*m.Event, error) {
	var event *m.Event
	var err error
	if event, err = m.NewEvent(id, title, start, end, owner); err != nil {
		return nil, err
	}
	err = app.repo.Put(event)
	return event, err
}

func (app *App) UpdateEvent(ctx context.Context, event *m.Event) error {
	return app.repo.Put(event)
}

func (app *App) GetEvent(ctx context.Context, id string) (*m.Event, error) {
	return app.repo.Get(id)
}

func (app *App) DeleteEvent(ctx context.Context, event *m.Event) error {
	return app.repo.Delete(event)
}

func (app *App) GetDailyAgenda(ctx context.Context, start time.Time) ([]*m.Event, error) {
	return app.repo.GetDayList(start)
}

func (app *App) GetWeeklyAgenda(ctx context.Context, start time.Time) ([]*m.Event, error) {
	return app.repo.GetWeekList(start)
}

func (app *App) GetMonthlyAgenda(ctx context.Context, start time.Time) ([]*m.Event, error) {
	return app.repo.GetMonthList(start)
}
