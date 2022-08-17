package sqlstorage

import (
	"context"
	"fmt"
	"time"

	"github.com/dingowd/otus/hw12_13_14_15_calendar/internal/storage"
	"gopkg.in/metakeule/fmtdate.v1"

	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

type Storage struct {
	DB *sqlx.DB
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, dsn string) error {
	var err error
	s.DB, err = sqlx.Open("pgx", dsn)
	return err
}

func (s *Storage) Close() error {
	return s.DB.Close()
}

func (s *Storage) IsEventExist(e storage.Event) (bool, error) {
	rows, err := s.DB.Query("select * from events where Start_date = $1 and End_date = $2", e.StartDate, e.EndDate)
	if err != nil {
		return false, err
	}
	if rows.Err() != nil {
		return false, rows.Err()
	}
	defer rows.Close()
	events := []storage.Event{}
	for rows.Next() {
		e := storage.Event{}
		var id int
		err := rows.Scan(&id, &e.Owner, &e.Title, &e.Descr, &e.StartDate, &e.StartTime, &e.EndDate, &e.EndTime)
		if err != nil {
			fmt.Println(err)
			continue
		}
		events = append(events, e)
	}
	return len(events) != 0, nil
}

func (s *Storage) Create(e storage.Event) error {
	exist, _ := s.IsEventExist(e)
	if exist {
		return storage.ErrorDateBusy
	}
	_, err := s.DB.Exec("insert into events(owner, title, descr, StartDate, StartTime, EndDate, EndTime)"+
		"values($1, $2, $3, $4, $5, $6, $7)",
		e.Owner, e.Title, e.Descr, e.StartDate, e.StartTime, e.EndDate, e.EndTime)
	return err
}

func (s *Storage) Update(id int, e storage.Event) error {
	_, err := s.DB.Exec("update events set owner = $1, title = $2, descr = $3,"+
		"StartDate = $4, StartTime = $5, EndDate = $6, EndTime = &7 where id = $8",
		e.Owner, e.Title, e.Descr, e.StartDate, e.StartTime, e.EndDate, e.EndTime, id)
	return err
}

func (s *Storage) Delete(id int) error {
	_, err := s.DB.Exec("delete from events where id = $1", id)
	return err
}

func (s *Storage) GetIntervalEvent(day string, n int) ([]storage.Event, error) {
	date, err := fmtdate.Parse(fmtdate.DefaultDateFormat, day)
	if err != nil {
		return nil, storage.DateFormatError
	}
	timeOut := date.Add(time.Duration(n) * time.Hour)
	day2 := fmtdate.Format(fmtdate.DefaultDateFormat, timeOut)
	events := []storage.Event{}
	rows, _ := s.DB.Query("select * from events where StartDate between $1 and $2", day, day2)
	if rows.Err() != nil {
		return events, rows.Err()
	}
	defer rows.Close()
	for rows.Next() {
		var sd, st, ed, et time.Time
		e := storage.Event{}
		var id int
		err := rows.Scan(&id, &e.Owner, &e.Title, &e.Descr, &sd, &st, &ed, &et)
		e.StartDate = fmtdate.Format(fmtdate.DefaultDateFormat, sd)
		e.StartTime = fmtdate.Format(fmtdate.DefaultTimeFormat, st)
		e.EndDate = fmtdate.Format(fmtdate.DefaultDateFormat, ed)
		e.EndTime = fmtdate.Format(fmtdate.DefaultTimeFormat, et)
		if err != nil {
			fmt.Println(err)
			continue
		}
		events = append(events, e)
	}
	return events, nil
}

func (s *Storage) GetDayEvent(day string) ([]storage.Event, error) {
	return s.GetIntervalEvent(day, 0)
}

func (s *Storage) GetWeekEvent(day string) ([]storage.Event, error) {
	return s.GetIntervalEvent(day, 168)
}

func (s *Storage) GetMonthEvent(day string) ([]storage.Event, error) {
	return s.GetIntervalEvent(day, 720)
}
