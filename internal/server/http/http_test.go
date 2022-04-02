package serverhttp

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/app"
	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/common"
	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/server/http/gen"
	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/storage"
	"github.com/stretchr/testify/require"
)

func eventsIdentical(t *testing.T, l, r gen.Event) {
	require.True(t, eventsEqual(t, l, r) && l.Id == r.Id)
}

func eventsEqual(t *testing.T, l, r gen.Event) bool {
	equal := l.StartsAt.Equal(r.StartsAt) &&
		l.EndsAt.Equal(r.EndsAt) &&
		l.OwnerId == r.OwnerId &&
		l.Notes == r.Notes && l.Title == r.Title
	require.True(t, equal)
	return equal
}

func TestHttpServer(t *testing.T) {
	ctx := context.Background()
	s := makeServer()
	ts := httptest.NewServer(s.buildRouter())
	tc, _ := gen.NewClientWithResponses(ts.URL + "/api")

	startDate := time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC)
	startTime := startDate.Add(8 * time.Hour)
	t.Run("Get Event", func(t *testing.T) {
		event, err := s.app.CreateEvent(ctx, "", "today event", startTime, startTime.Add(time.Hour), "test")
		require.NoError(t, err)

		expected := gen.Event{
			AlertBefore: int(event.AlertBefore.Seconds()),
			EndsAt:      event.EndsAt,
			Id:          event.ID,
			Notes:       event.Notes,
			OwnerId:     event.OwnerID,
			StartsAt:    event.StartsAt,
			Title:       event.Title,
		}
		r, err := tc.GetEventWithResponse(ctx, event.ID)
		require.NoError(t, err)
		require.Equal(t, r.StatusCode(), http.StatusOK)
		eventsIdentical(t, *r.JSON200, expected)

		r, err = tc.GetEventWithResponse(ctx, "42")
		require.NoError(t, err)
		require.Equal(t, r.StatusCode(), http.StatusNotFound)
	})

	t.Run("Delete Event", func(t *testing.T) {
		event, err := s.app.CreateEvent(ctx, "", "today event", startTime, startTime.Add(time.Hour), "test")
		require.NoError(t, err)

		rDel, err := tc.DeleteEventWithResponse(ctx, event.ID)
		require.NoError(t, err)
		require.Equal(t, rDel.StatusCode(), http.StatusOK)

		rGet, err := tc.GetEventWithResponse(ctx, event.ID)
		require.NoError(t, err)
		require.Equal(t, rGet.StatusCode(), http.StatusNotFound)
	})

	t.Run("Update Event", func(t *testing.T) {
		event, err := s.app.CreateEvent(ctx, "", "today event", startTime, startTime.Add(time.Hour), "test")
		require.NoError(t, err)

		expected := gen.Event{
			AlertBefore: int(event.AlertBefore.Seconds()),
			EndsAt:      event.EndsAt,
			Id:          event.ID,
			Notes:       "42",
			OwnerId:     event.OwnerID,
			StartsAt:    event.StartsAt,
			Title:       event.Title,
		}
		r, err := tc.PutEventWithResponse(ctx, event.ID, gen.PutEventJSONRequestBody(expected))
		require.NoError(t, err)
		require.Equal(t, r.StatusCode(), http.StatusOK)
		eventsIdentical(t, *r.JSON200, expected)

		rGet, err := tc.GetEventWithResponse(ctx, event.ID)
		require.NoError(t, err)
		require.Equal(t, rGet.StatusCode(), http.StatusOK)
		eventsIdentical(t, *rGet.JSON200, expected)
	})

	t.Run("Create Event", func(t *testing.T) {
		expected := gen.Event{
			AlertBefore: 0,
			Title:       "today event",
			Notes:       "42",
			OwnerId:     "test",
			StartsAt:    startTime,
			EndsAt:      startTime.Add(time.Hour),
		}
		r, err := tc.CreateEventWithResponse(ctx, gen.CreateEventJSONRequestBody(expected))
		require.NoError(t, err)
		require.Equal(t, r.StatusCode(), http.StatusOK)
		event := *r.JSON200
		eventsEqual(t, event, expected)

		rGet, err := tc.GetEventWithResponse(ctx, event.Id)
		require.NoError(t, err)
		require.Equal(t, rGet.StatusCode(), http.StatusOK)
		eventsIdentical(t, *rGet.JSON200, event)
	})

	t.Run("List Events", func(t *testing.T) {
		startTime := startTime.AddDate(0, 1, 0)
		event, err := s.app.CreateEvent(
			ctx, "",
			"today event",
			startTime,
			startTime.Add(time.Hour),
			"test",
		)
		require.NoError(t, err)

		expected := gen.Event{
			AlertBefore: int(event.AlertBefore.Seconds()),
			EndsAt:      event.EndsAt,
			Id:          event.ID,
			OwnerId:     event.OwnerID,
			StartsAt:    event.StartsAt,
			Title:       event.Title,
		}

		params := &gen.ListEventsParams{
			Agenda:    "daily",
			StartFrom: startDate.AddDate(0, 1, 0),
		}
		r, err := tc.ListEventsWithResponse(ctx, params)
		require.NoError(t, err)
		require.Equal(t, r.StatusCode(), http.StatusOK)
		events := *r.JSON200
		require.Len(t, events, 1)
		eventsIdentical(t, events[0], expected)
	})
}

func makeServer() *HTTPServer {
	log := common.NewLogger("debug", "")

	eventsRepo, err := storage.NewStorage("in-memory")
	if err != nil {
		log.Fatalf("storage was not created: %v", err)
	}
	calendar := app.New(log, eventsRepo)
	return NewServer("", "", log, calendar)
}
