package models

import "time"

// Alert - a temporary object, not stored in DB, just put it in a queue:
type Alert struct {
	ID        string
	EventID   string
	Title     string
	StartAt   time.Time
	Addressee string
}

func NewAlert(e *Event) *Alert {
	return &Alert{
		ID:        uniqueID(),
		EventID:   e.ID,
		Title:     e.Title,
		StartAt:   e.StartsAt,
		Addressee: e.OwnerID,
	}
}
