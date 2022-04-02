package serverhttp

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/app"
	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/common"
	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/models"
	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/server/http/gen"
	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/server/http/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type HTTPServer struct {
	app    *app.App
	host   string
	port   string
	log    common.Logger
	server http.Server
}

func (s *HTTPServer) CreateEvent(w http.ResponseWriter, r *http.Request) {
	eventDto := gen.Event{}
	if err := render.Decode(r, &eventDto); err != nil {
		s.BadRequest(err, w, r)
		return
	}

	event, err := models.NewEvent("", eventDto.Title, eventDto.StartsAt, eventDto.EndsAt, eventDto.OwnerId)
	if err != nil {
		s.BadRequest(err, w, r)
		return
	}
	event.AlertBefore = time.Duration(eventDto.AlertBefore * 1_000_000_000)
	event.Notes = eventDto.Notes
	if err = s.app.UpdateEvent(r.Context(), event); err != nil {
		s.InternalError(err, w, r)
		return
	}
	render.Respond(w, r, &gen.Event{
		AlertBefore: int(event.AlertBefore.Seconds()),
		EndsAt:      event.EndsAt,
		Id:          event.ID,
		Notes:       event.Notes,
		OwnerId:     event.OwnerID,
		StartsAt:    event.StartsAt,
		Title:       event.Title,
	})
}

func (s *HTTPServer) ListEvents(w http.ResponseWriter, r *http.Request, params gen.ListEventsParams) {
	var events []*models.Event
	var err error
	switch params.Agenda {
	case "daily":
		events, err = s.app.GetDailyAgenda(r.Context(), params.StartFrom)
	case "weekly":
		events, err = s.app.GetWeeklyAgenda(r.Context(), params.StartFrom)
	case "monthly":
		events, err = s.app.GetMonthlyAgenda(r.Context(), params.StartFrom)
	default:
		err = fmt.Errorf("http: invalid agenda type %s", params.Agenda)
		s.BadRequest(err, w, r)
		return
	}

	if err != nil {
		s.InternalError(err, w, r)
		return
	}

	result := make([]*gen.Event, 0, len(events))
	for _, e := range events {
		result = append(result, &gen.Event{
			AlertBefore: int(e.AlertBefore.Seconds()),
			EndsAt:      e.EndsAt,
			Id:          e.ID,
			Notes:       e.Notes,
			OwnerId:     e.OwnerID,
			StartsAt:    e.StartsAt,
			Title:       e.Title,
		})
	}
	render.Respond(w, r, result)
}

func (s *HTTPServer) DeleteEvent(w http.ResponseWriter, r *http.Request, id string) {
	if e, err := s.app.GetEvent(r.Context(), id); errors.Is(err, models.ErrNotFound) {
		s.NotFound(err, w, r)
		return
	} else if err != nil {
		s.InternalError(err, w, r)
		return
	} else if err = s.app.DeleteEvent(r.Context(), e); err != nil {
		s.InternalError(err, w, r)
		return
	}
	s.NoEror(w, r)
}

func (s *HTTPServer) GetEvent(w http.ResponseWriter, r *http.Request, id string) {
	switch e, err := s.app.GetEvent(r.Context(), id); {
	case errors.Is(err, models.ErrNotFound):
		s.NotFound(err, w, r)
		return
	case err != nil:
		s.InternalError(err, w, r)
		return
	default:
		render.Respond(w, r, &gen.Event{
			AlertBefore: int(e.AlertBefore.Seconds()),
			EndsAt:      e.EndsAt,
			Id:          e.ID,
			Notes:       e.Notes,
			OwnerId:     e.OwnerID,
			StartsAt:    e.StartsAt,
			Title:       e.Title,
		})
	}
}

func (s *HTTPServer) PutEvent(w http.ResponseWriter, r *http.Request, id string) {
	eventDto := gen.Event{}
	if err := render.Decode(r, &eventDto); err != nil {
		s.BadRequest(err, w, r)
		return
	}

	event, err := models.NewEvent(id, eventDto.Title, eventDto.StartsAt, eventDto.EndsAt, eventDto.OwnerId)
	if err != nil {
		s.BadRequest(err, w, r)
		return
	}
	event.AlertBefore = time.Duration(eventDto.AlertBefore * 1_000_000_000)
	event.Notes = eventDto.Notes
	if err = s.app.UpdateEvent(r.Context(), event); err != nil {
		s.InternalError(err, w, r)
		return
	}
	render.Respond(w, r, &gen.Event{
		AlertBefore: int(event.AlertBefore.Seconds()),
		EndsAt:      event.EndsAt,
		Id:          event.ID,
		Notes:       event.Notes,
		OwnerId:     event.OwnerID,
		StartsAt:    event.StartsAt,
		Title:       event.Title,
	})
}

func NewServer(host string, port string, logger common.Logger, app *app.App) *HTTPServer {
	return &HTTPServer{host: host, port: port, app: app, log: logger}
}

func (s *HTTPServer) Start() error {
	rootRouter := s.buildRouter()
	s.server = http.Server{
		Addr:    net.JoinHostPort(s.host, s.port),
		Handler: rootRouter,
	}
	go func() {
		s.log.Error(s.server.ListenAndServe())
	}()
	return nil
}

func (s *HTTPServer) buildRouter() *chi.Mux {
	apiRouter := chi.NewRouter()
	apiRouter.Use(middleware.LoggingMiddleware(s.log))
	rootRouter := chi.NewRouter()
	rootRouter.Mount("/api", gen.HandlerFromMux(s, apiRouter))
	return rootRouter
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
