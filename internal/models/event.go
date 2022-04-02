package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Event - main object
type Event struct {
	ID          string
	Title       string
	StartsAt    time.Time
	EndsAt      time.Time
	Notes       string
	OwnerID     string
	AlertBefore time.Duration
}

type EventsRepo interface {
	Get(id string) (*Event, error)
	Put(e *Event) error
	Delete(e *Event) error
	GetDayList(d time.Time) ([]*Event, error)
	GetWeekList(d time.Time) ([]*Event, error)
	GetMonthList(d time.Time) ([]*Event, error)
	IsBusy(d1, d2 time.Time) (bool, error)
}

func NewEvent(id string, title string, start time.Time, end time.Time, owner string) (*Event, error) {
	if !FitsOneDay(start, end) {
		return nil, fmt.Errorf("%w: start and end must be of the same date", ErrValueError)
	}

	if len(id) == 0 {
		id = uniqueID()
	}
	return &Event{
		ID:       id,
		Title:    title,
		StartsAt: time.Unix(start.Unix(), 0),
		EndsAt:   time.Unix(end.Unix(), 0),
		OwnerID:  owner,
	}, nil
}

func FitsOneDay(start time.Time, end time.Time) bool {
	y1, m1, d1 := start.Date()
	y2, m2, d2 := end.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}

func uniqueID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
