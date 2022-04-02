package storage

import (
	"fmt"

	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/models"
	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/storage/mem"
	"github.com/VladNF/otus-golang/hw12_13_14_15_calendar/internal/storage/pgsql"
)

func NewStorage(storageType string) (models.EventsRepo, error) {
	switch storageType {
	case "in-memory":
		return mem.NewMemoryStorage(), nil
	case "pgsql":
		db, _ := pgsql.NewPgSQLConnection()
		return pgsql.NewPgSQLStorage(db), nil
	default:
		return nil, fmt.Errorf("unsupported storage type %v", storageType)
	}
}
