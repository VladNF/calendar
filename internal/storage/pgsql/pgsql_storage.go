package pgsql

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/VladNF/calendar/internal/models"

	// init driver.
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

type sqlEvent struct {
	ID          string    `db:"id"`
	Title       string    `db:"title"`
	StartsAt    time.Time `db:"start_at"`
	EndsAt      time.Time `db:"end_at"`
	Notes       string    `db:"notes"`
	OwnerID     string    `db:"owner"`
	AlertBefore int64     `db:"alert_before"`
}

func (e *sqlEvent) asModel() (*models.Event, error) {
	event, err := models.NewEvent(e.ID, e.Title, e.StartsAt, e.EndsAt, e.OwnerID)
	if err != nil {
		return nil, fmt.Errorf("%w: unexpected error %v", models.ErrDataError, err)
	}
	event.AlertBefore = time.Duration(e.AlertBefore)
	event.Notes = e.Notes
	return event, nil
}

type PgStorage struct {
	db *sqlx.DB
}

func (s *PgStorage) Get(id string) (*models.Event, error) {
	dbEvent := sqlEvent{}
	query := "SELECT * FROM events WHERE id = $1"
	if err := s.db.Get(&dbEvent, query, id); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, models.ErrNotFound
		default:
			return nil, fmt.Errorf("%w: unexpected error %v", models.ErrNotFound, err)
		}
	} else {
		return dbEvent.asModel()
	}
}

func (s *PgStorage) Put(e *models.Event) error {
	dbEvent := sqlEvent{
		ID:          e.ID,
		Title:       e.Title,
		StartsAt:    e.StartsAt,
		EndsAt:      e.EndsAt,
		Notes:       e.Notes,
		OwnerID:     e.OwnerID,
		AlertBefore: e.AlertBefore.Nanoseconds(),
	}
	query := `INSERT INTO events 
				(id, owner, title, notes, start_at, end_at, alert_before)
			VALUES
				(:id, :owner, :title, :notes, :start_at, :end_at, :alert_before)
			ON CONFLICT (id) DO UPDATE
			SET
				owner = EXCLUDED.owner, 
				title = EXCLUDED.title, 
				notes  = EXCLUDED.notes, 
				start_at  = EXCLUDED.start_at, 
				end_at  = EXCLUDED.end_at, 
				alert_before  = EXCLUDED.alert_before`
	if _, err := s.db.NamedExec(query, dbEvent); err != nil {
		return fmt.Errorf("%w: put failed -> %v", models.ErrDataError, err)
	}
	return nil
}

func (s *PgStorage) Delete(e *models.Event) error {
	query := "DELETE FROM events WHERE id = $1"
	if r, err := s.db.Exec(query, e.ID); err != nil {
		return fmt.Errorf("%w: unexpected error -> %v", models.ErrDataError, err)
	} else if rows, _ := r.RowsAffected(); rows == 0 {
		return models.ErrNotFound
	}

	return nil
}

func (s *PgStorage) getEventList(lBound time.Time, uBound time.Time) ([]*models.Event, error) {
	query := `SELECT * from events AS e 
				WHERE e.start_at > to_timestamp($1) AND e.start_at < to_timestamp($2)
				ORDER BY e.start_at`
	if rows, err := s.db.Queryx(query, lBound.Unix(), uBound.Unix()); err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, models.ErrNotFound
		default:
			return nil, fmt.Errorf("%w: unexpected error -> %v", models.ErrNotFound, err)
		}
	} else {
		var results []*models.Event
		defer rows.Close()
		for rows.Next() {
			dbEvent := sqlEvent{}
			if err := rows.StructScan(&dbEvent); err != nil {
				return nil, fmt.Errorf("%w: unexpected error -> %v", models.ErrDataError, err)
			}
			if m, err := dbEvent.asModel(); err == nil {
				results = append(results, m)
			}
		}
		return results, nil
	}
}

func (s *PgStorage) GetDayList(d time.Time) ([]*models.Event, error) {
	yy, mm, dd := d.Date()
	lBound := time.Date(yy, mm, dd, 0, 0, 0, -1, d.Location())
	uBound := time.Date(yy, mm, dd+1, 0, 0, 0, 0, d.Location())
	return s.getEventList(lBound, uBound)
}

func (s *PgStorage) GetWeekList(d time.Time) ([]*models.Event, error) {
	yy, mm, dd := d.Date()
	sunday := time.Date(yy, mm, dd+int(time.Sunday-d.Weekday()), 0, 0, 0, 0, d.Location())
	saturday := sunday.AddDate(0, 0, 7)
	return s.getEventList(sunday, saturday)
}

func (s *PgStorage) GetMonthList(d time.Time) ([]*models.Event, error) {
	yy, mm, _ := d.Date()
	lBound := time.Date(yy, mm, 1, 0, 0, 0, -1, d.Location())
	uBound := time.Date(yy, mm+1, 1, 0, 0, 0, -1, d.Location())
	return s.getEventList(lBound, uBound)
}

func (s *PgStorage) IsBusy(d1, d2 time.Time) (bool, error) {
	var overlapCount int
	query := "SELECT COUNT(*) FROM events AS e WHERE (e.start_at, e.end_at) OVERLAPS ($1, $2)"
	if err := s.db.Get(&overlapCount, query, d1.Unix(), d2.Unix()); err != nil {
		return false, fmt.Errorf("%w: unexpected error -> %v", models.ErrDataError, err)
	}

	return overlapCount > 0, nil
}

func NewPgSQLStorage(db *sqlx.DB) models.EventsRepo {
	return &PgStorage{db}
}

func NewPgSQLConnection() (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
	)
	return sqlx.Open("pgx", dsn)
}
