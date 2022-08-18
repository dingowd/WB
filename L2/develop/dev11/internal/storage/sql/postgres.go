package sqlstorage

import (
	"context"
	"fmt"
	"github.com/dingowd/WB/L2/develop/dev11/internal/logger"
	"github.com/dingowd/WB/L2/develop/dev11/models"
	"time"

	"github.com/dingowd/WB/L2/develop/dev11/internal/storage"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/stdlib"
)

type Storage struct {
	DB  *sqlx.DB
	Log logger.Logger
}

func New(l logger.Logger) *Storage {
	return &Storage{
		Log: l,
	}
}

func (s *Storage) Connect(ctx context.Context, dsn string) error {
	var err error
	s.DB, err = sqlx.Open("pgx", dsn)
	if err == nil {
		s.Log.Info("DB connected")
	}
	return err
}

func (s *Storage) Close() error {
	s.Log.Info("DB disconnected")
	return s.DB.Close()
}

func (s *Storage) IsEventExist(e models.Event) (bool, error) {
	rows, err := s.DB.Query("select * from events where start_date = $1 "+
		"and end_date = $2 and owner = $3 and title = $4 and descr = $5", e.StartDate, e.EndDate, e.Owner, e.Title, e.Descr)
	if err != nil {
		return false, err
	}
	if rows.Err() != nil {
		return false, rows.Err()
	}
	defer rows.Close()
	events := []models.Event{}
	for rows.Next() {
		e := models.Event{}
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

func (s *Storage) Create(e models.Event) error {
	exist, _ := s.IsEventExist(e)
	if exist {
		return storage.ErrorDateBusy
	}
	_, err := s.DB.Exec("insert into events(owner, title, descr, start_date, start_time, end_date, end_time)"+
		"values($1, $2, $3, $4, $5, $6, $7)",
		e.Owner, e.Title, e.Descr, e.StartDate, e.StartTime, e.EndDate, e.EndTime)
	return err
}

func (s *Storage) Update(e models.DBEvent) error {
	_, err := s.DB.NamedExec("update events set owner = :owner, title = :title, descr = :descr, "+
		"start_date = :sd, start_time = :st, end_date = :ed, end_time = :et where id = :id",
		map[string]interface{}{"owner": e.Owner, "title": e.Title, "descr": e.Descr, "sd": e.StartDate, "st": e.StartTime, "ed": e.EndDate, "et": e.EndTime, "id": e.ID})
	return err
}

func (s *Storage) Delete(id int) error {
	_, err := s.DB.Exec("delete from events where id = $1", id)
	return err
}

func (s *Storage) GetEventByInterval(id int, day string, n int) ([]models.DBEvent, error) {
	date, err := time.Parse("2006-01-02", day)
	if err != nil {
		return nil, storage.DateFormatError
	}
	timeOut := date.Add(time.Duration(n) * time.Hour)
	day2 := timeOut.Format("2006-01-02")
	events := []models.DBEvent{}
	rows, _ := s.DB.Query("select * from events where (start_date between $1 and $2) and owner = $3", day, day2, id)
	if rows.Err() != nil {
		return nil, rows.Err()
	}
	defer rows.Close()
	for rows.Next() {
		//var sd, st, ed, et time.Time
		e := models.DBEvent{}
		err := rows.Scan(&e.ID, &e.Owner, &e.Title, &e.Descr, &e.StartDate, &e.StartTime, &e.EndDate, &e.EndTime)
		sd, _ := time.Parse(time.RFC3339, e.StartDate)
		e.StartDate = sd.Format("2006-01-02")
		ed, _ := time.Parse(time.RFC3339, e.EndDate)
		e.EndDate = ed.Format("2006-01-02")
		if err != nil {
			fmt.Println(err)
			continue
		}
		events = append(events, e)
	}
	return events, nil
}

func (s *Storage) GetDayEvent(id int, day string) ([]models.DBEvent, error) {
	return s.GetEventByInterval(id, day, 0)
}

func (s *Storage) GetWeekEvent(id int, day string) ([]models.DBEvent, error) {
	return s.GetEventByInterval(id, day, 168)
}

func (s *Storage) GetMonthEvent(id int, day string) ([]models.DBEvent, error) {
	return s.GetEventByInterval(id, day, 720)
}
