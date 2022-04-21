package mem

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/VladNF/calendar/internal/models"
)

type EventList map[string]*models.Event

type MemoryStorage struct {
	sync.RWMutex
	eventFromID  EventList
	eventFromDay map[string]EventList
}

func isoDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func joinEventLists(eventLists ...EventList) []*models.Event {
	size := 0
	for _, l := range eventLists {
		size += len(l)
	}
	r := make([]*models.Event, 0, size)
	for _, l := range eventLists {
		for _, e := range l {
			r = append(r, e)
		}
	}
	sort.Slice(r, func(i, j int) bool { return r[i].StartsAt.Before(r[j].StartsAt) })
	return r
}

func (s *MemoryStorage) Get(id string) (*models.Event, error) {
	s.RLock()
	defer s.RUnlock()
	if e, ok := s.eventFromID[id]; ok {
		return e, nil
	}

	return nil, models.ErrNotFound
}

func (s *MemoryStorage) Put(e *models.Event) error {
	s.Lock()
	defer s.Unlock()
	s.eventFromID[e.ID] = e
	if _, ok := s.eventFromDay[isoDate(e.StartsAt)]; !ok {
		s.eventFromDay[isoDate(e.StartsAt)] = make(EventList)
	}
	s.eventFromDay[isoDate(e.StartsAt)][e.ID] = e
	return nil
}

func (s *MemoryStorage) Delete(e *models.Event) error {
	s.Lock()
	defer s.Unlock()
	delete(s.eventFromID, e.ID)
	delete(s.eventFromDay[isoDate(e.StartsAt)], e.ID)
	return nil
}

func (s *MemoryStorage) GetDayList(d time.Time) ([]*models.Event, error) {
	s.RLock()
	defer s.RUnlock()
	return joinEventLists(s.eventFromDay[isoDate(d)]), nil
}

func (s *MemoryStorage) GetWeekList(d time.Time) ([]*models.Event, error) {
	s.RLock()
	defer s.RUnlock()
	dayLists := make([]EventList, 0, 7)
	sunday := d.AddDate(0, 0, int(time.Sunday-d.Weekday()))
	for i := 0; i < 7; i++ {
		dayLists = append(dayLists, s.eventFromDay[isoDate(sunday.AddDate(0, 0, i))])
	}
	return joinEventLists(dayLists...), nil
}

func (s *MemoryStorage) GetMonthList(d time.Time) ([]*models.Event, error) {
	s.RLock()
	defer s.RUnlock()
	dayLists := make([]EventList, 0, 31)
	y, m, _ := d.Date()
	day := time.Date(y, m, 1, 0, 0, 0, 0, d.Location())
	for day.Month() == m {
		dayLists = append(dayLists, s.eventFromDay[isoDate(day)])
		day = day.AddDate(0, 0, 1)
	}
	return joinEventLists(dayLists...), nil
}

func (s *MemoryStorage) IsBusy(d1, d2 time.Time) (bool, error) {
	if !models.FitsOneDay(d1, d1) {
		return false, fmt.Errorf("%w: start and end must be of the same date", models.ErrValueError)
	}
	s.RLock()
	defer s.RUnlock()
	events, err := s.GetDayList(d1)
	if err != nil {
		return false, err
	}
	for _, e := range events {
		if (e.StartsAt.After(d1) && e.StartsAt.Before(d2)) || (e.EndsAt.After(d1) && e.EndsAt.Before(d2)) {
			return true, nil
		}
	}

	return false, nil
}

func NewMemoryStorage() models.EventsRepo {
	return &MemoryStorage{
		eventFromID:  make(EventList),
		eventFromDay: make(map[string]EventList),
	}
}
