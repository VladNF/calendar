// +build e2e

package e2e

import (
	"context"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/VladNF/calendar/internal/server/http/gen"
	"github.com/stretchr/testify/require"
)

func TestE2EServer(t *testing.T) {
	ctx := context.Background()
	require.True(t, waitForPort("localhost:4200"))
	require.True(t, waitForPort("localhost:5432"))
	tc, _ := gen.NewClientWithResponses("http://localhost:4200" + "/api")

	startDate := time.Date(2021, 1, 1, 8, 0, 0, 0, time.UTC)
	startTime := startDate.Add(8 * time.Hour)
	t.Run("Get Event", func(t *testing.T) {
		event := createEvent(t, tc, startTime)

		expected := gen.Event{
			AlertBefore: event.AlertBefore,
			EndsAt:      event.EndsAt,
			Id:          event.Id,
			Notes:       event.Notes,
			OwnerId:     event.OwnerId,
			StartsAt:    event.StartsAt,
			Title:       event.Title,
		}
		r, err := tc.GetEventWithResponse(ctx, event.Id)
		require.NoError(t, err)
		require.Equal(t, r.StatusCode(), http.StatusOK)
		eventsIdentical(t, *r.JSON200, expected)
		deleteEvent(t, tc, event.Id)

		r, err = tc.GetEventWithResponse(ctx, "42")
		require.NoError(t, err)
		require.Equal(t, r.StatusCode(), http.StatusNotFound)
	})

	t.Run("Delete Event", func(t *testing.T) {
		event := createEvent(t, tc, startTime)

		rDel, err := tc.DeleteEventWithResponse(ctx, event.Id)
		require.NoError(t, err)
		require.Equal(t, rDel.StatusCode(), http.StatusOK)

		rGet, err := tc.GetEventWithResponse(ctx, event.Id)
		require.NoError(t, err)
		require.Equal(t, rGet.StatusCode(), http.StatusNotFound)
	})

	t.Run("Update Event", func(t *testing.T) {
		event := createEvent(t, tc, startTime)

		expected := gen.Event{
			AlertBefore: event.AlertBefore,
			EndsAt:      event.EndsAt,
			Id:          event.Id,
			Notes:       "42",
			OwnerId:     event.OwnerId,
			StartsAt:    event.StartsAt,
			Title:       event.Title,
		}
		r, err := tc.PutEventWithResponse(ctx, event.Id, gen.PutEventJSONRequestBody(expected))
		require.NoError(t, err)
		require.Equal(t, r.StatusCode(), http.StatusOK)
		eventsIdentical(t, *r.JSON200, expected)

		rGet, err := tc.GetEventWithResponse(ctx, event.Id)
		require.NoError(t, err)
		require.Equal(t, rGet.StatusCode(), http.StatusOK)
		eventsIdentical(t, *rGet.JSON200, expected)

		deleteEvent(t, tc, event.Id)
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

		deleteEvent(t, tc, event.Id)
	})

	t.Run("List Events", func(t *testing.T) {
		startTime := startTime.AddDate(0, 3, 0)
		event := createEvent(t, tc, startTime)

		expected := gen.Event{
			AlertBefore: event.AlertBefore,
			EndsAt:      event.EndsAt,
			Id:          event.Id,
			OwnerId:     event.OwnerId,
			StartsAt:    event.StartsAt,
			Title:       event.Title,
		}

		params := &gen.ListEventsParams{
			Agenda:    "daily",
			StartFrom: startDate.AddDate(0, 3, 0),
		}
		r, err := tc.ListEventsWithResponse(ctx, params)
		require.NoError(t, err)
		require.Equal(t, r.StatusCode(), http.StatusOK)
		events := *r.JSON200
		require.Len(t, events, 1)
		eventsIdentical(t, events[0], expected)

		deleteEvent(t, tc, event.Id)
	})
}

func createEvent(t *testing.T, tc *gen.ClientWithResponses, startTime time.Time) gen.Event {
	r, err := tc.CreateEventWithResponse(context.Background(), gen.CreateEventJSONRequestBody{
		AlertBefore: 0,
		EndsAt:      startTime.Add(time.Hour),
		Id:          "",
		Notes:       "",
		OwnerId:     "test",
		StartsAt:    startTime,
		Title:       "today event",
	})
	require.NoError(t, err)
	require.Equal(t, r.StatusCode(), http.StatusOK)
	return *r.JSON200
}

func deleteEvent(t *testing.T, tc *gen.ClientWithResponses, id string) {
	r, err := tc.DeleteEventWithResponse(context.Background(), id)
	require.NoError(t, err)
	require.Equal(t, r.StatusCode(), http.StatusOK)
}

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

func waitForPort(address string) bool {
	waitChan := make(chan struct{})

	go func() {
		for {
			conn, err := net.DialTimeout("tcp", address, time.Second)
			if err != nil {
				time.Sleep(time.Second)
				continue
			}

			if conn != nil {
				waitChan <- struct{}{}
				return
			}
		}
	}()

	timeout := time.After(10 * time.Second)
	select {
	case <-waitChan:
		return true
	case <-timeout:
		return false
	}
}
