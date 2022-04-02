package storage

import (
	"testing"
	"time"

	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/models"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	eventsRepo, _ := NewStorage("in-memory")

	t.Run("basic test", func(t *testing.T) {
		testBasicOperations(t, eventsRepo)
	})

	t.Run("day view test", func(t *testing.T) {
		testDayViewQuery(t, eventsRepo)
	})

	t.Run("week view test", func(t *testing.T) {
		testWeekViewQuery(t, eventsRepo)
	})

	t.Run("month view test", func(t *testing.T) {
		testMonthViewQuery(t, eventsRepo)
	})
}

func testMonthViewQuery(t *testing.T, eventsRepo models.EventsRepo) {
	ny2021 := time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC)

	startAt, endAt := ny2021, ny2021.Add(time.Hour)
	eventNY1, _ := models.NewEvent("", "title 1", startAt, endAt, "1")

	startAt, endAt = startAt.AddDate(0, 0, 1), endAt.AddDate(0, 0, 1)
	eventNY2, _ := models.NewEvent("", "title 2", startAt, endAt, "1")

	startAt, endAt = startAt.AddDate(0, 0, 1), endAt.AddDate(0, 0, 1)
	eventNY3, _ := models.NewEvent("", "title 3", startAt, endAt, "1")

	startAt, endAt = startAt.AddDate(0, 0, 1), endAt.AddDate(0, 0, 1)
	eventNextWeek, _ := models.NewEvent("", "title 4", startAt, endAt, "1")
	require.NoError(t, eventsRepo.Put(eventNY3))
	require.NoError(t, eventsRepo.Put(eventNY2))
	require.NoError(t, eventsRepo.Put(eventNY1))
	require.NoError(t, eventsRepo.Put(eventNextWeek))

	list, err := eventsRepo.GetMonthList(time.Now())
	require.NoError(t, err)
	require.Len(t, list, 0)

	list, err = eventsRepo.GetMonthList(ny2021)
	require.NoError(t, err)
	require.Len(t, list, 1)
	require.Equal(t, *eventNY1, *list[0])

	list, err = eventsRepo.GetMonthList(startAt)
	require.NoError(t, err)
	require.Len(t, list, 3)
	require.Equal(t, *eventNY2, *list[0])
	require.Equal(t, *eventNY3, *list[1])
	require.Equal(t, *eventNextWeek, *list[2])

	require.NoError(t, eventsRepo.Delete(eventNY1))
	list, err = eventsRepo.GetMonthList(ny2021)
	require.NoError(t, err)
	require.Len(t, list, 0)
	require.NoError(t, eventsRepo.Delete(eventNY2))
	require.NoError(t, eventsRepo.Delete(eventNY3))
	require.NoError(t, eventsRepo.Delete(eventNextWeek))
}

func testWeekViewQuery(t *testing.T, eventsRepo models.EventsRepo) {
	ny2021 := time.Date(2020, 12, 31, 0, 0, 0, 0, time.UTC)

	startAt, endAt := ny2021, ny2021.Add(time.Hour)
	eventNY1, _ := models.NewEvent("", "title 1", startAt, endAt, "1")

	startAt, endAt = startAt.AddDate(0, 0, 1), endAt.AddDate(0, 0, 1)
	eventNY2, _ := models.NewEvent("", "title 2", startAt, endAt, "1")

	startAt, endAt = startAt.AddDate(0, 0, 1), endAt.AddDate(0, 0, 1)
	eventNY3, _ := models.NewEvent("", "title 3", startAt, endAt, "1")

	startAt, endAt = startAt.AddDate(0, 0, 1), endAt.AddDate(0, 0, 1)
	eventNextWeek, _ := models.NewEvent("", "title 4", startAt, endAt, "1")
	require.NoError(t, eventsRepo.Put(eventNY3))
	require.NoError(t, eventsRepo.Put(eventNY2))
	require.NoError(t, eventsRepo.Put(eventNY1))
	require.NoError(t, eventsRepo.Put(eventNextWeek))

	list, err := eventsRepo.GetWeekList(time.Now())
	require.NoError(t, err)
	require.Len(t, list, 0)

	list, err = eventsRepo.GetWeekList(ny2021)
	require.NoError(t, err)
	require.Len(t, list, 3)
	require.Equal(t, *eventNY1, *list[0])
	require.Equal(t, *eventNY2, *list[1])
	require.Equal(t, *eventNY3, *list[2])

	require.NoError(t, eventsRepo.Delete(eventNY1))
	list, err = eventsRepo.GetWeekList(ny2021)
	require.NoError(t, err)
	require.Len(t, list, 2)
	require.NoError(t, eventsRepo.Delete(eventNY2))
	require.NoError(t, eventsRepo.Delete(eventNY3))
	require.NoError(t, eventsRepo.Delete(eventNextWeek))
}

func testDayViewQuery(t *testing.T, eventsRepo models.EventsRepo) {
	ny2021 := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	eventNY1, _ := models.NewEvent("", "title1", ny2021, ny2021.Add(time.Hour), "1")
	eventNY2, _ := models.NewEvent("", "title2", ny2021.Add(time.Hour), ny2021.Add(2*time.Hour), "1")
	eventNY3, _ := models.NewEvent("", "title3", ny2021.Add(2*time.Hour), ny2021.Add(3*time.Hour), "1")
	nextDay := ny2021.AddDate(0, 0, 1)
	eventNextDay, _ := models.NewEvent("", "title4", nextDay, nextDay.Add(time.Hour), "1")
	require.NoError(t, eventsRepo.Put(eventNY3))
	require.NoError(t, eventsRepo.Put(eventNY2))
	require.NoError(t, eventsRepo.Put(eventNY1))
	require.NoError(t, eventsRepo.Put(eventNextDay))

	list, err := eventsRepo.GetDayList(time.Now())
	require.NoError(t, err)
	require.Len(t, list, 0)

	list, err = eventsRepo.GetDayList(ny2021)
	require.NoError(t, err)
	require.Len(t, list, 3)
	require.Equal(t, *eventNY1, *list[0])
	require.Equal(t, *eventNY2, *list[1])
	require.Equal(t, *eventNY3, *list[2])

	require.NoError(t, eventsRepo.Delete(eventNY1))
	list, err = eventsRepo.GetDayList(ny2021)
	require.NoError(t, err)
	require.Len(t, list, 2)
	require.NoError(t, eventsRepo.Delete(eventNY2))
	require.NoError(t, eventsRepo.Delete(eventNY3))
	require.NoError(t, eventsRepo.Delete(eventNextDay))
}

func testBasicOperations(t *testing.T, eventsRepo models.EventsRepo) {
	start := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	event, err := models.NewEvent("", "title", start, start.Add(time.Hour), "1")
	require.NoError(t, err)
	require.NoError(t, eventsRepo.Put(event))

	queried, err := eventsRepo.Get(event.ID)
	require.NoError(t, err)
	require.Equal(t, *event, *queried)

	list, err := eventsRepo.GetDayList(start)
	require.NoError(t, err)
	require.Len(t, list, 1)
	require.Equal(t, *event, *list[0])

	list, err = eventsRepo.GetWeekList(start)
	require.NoError(t, err)
	require.Len(t, list, 1)
	require.Equal(t, *event, *list[0])

	list, err = eventsRepo.GetMonthList(start)
	require.NoError(t, err)
	require.Len(t, list, 1)
	require.Equal(t, *event, *list[0])

	require.NoError(t, eventsRepo.Delete(event))
	list, err = eventsRepo.GetDayList(start)
	require.NoError(t, err)
	require.Len(t, list, 0)
}
