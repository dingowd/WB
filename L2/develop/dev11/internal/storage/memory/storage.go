package memorystorage

import (
	"context"
	"github.com/dingowd/WB/L2/develop/dev11/models"
	"sync"
	"time"

	"github.com/dingowd/WB/L2/develop/dev11/internal/storage"
	"gitlab.com/metakeule/fmtdate"
)

type Storage struct {
	id     int
	Events map[int]models.Event
	mu     sync.RWMutex
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Connect(ctx context.Context, dsn string) error {
	s.id = 0
	s.Events = make(map[int]models.Event)
	return nil
}

func (s *Storage) Close() error {
	return nil
}

func (s *Storage) IsEventExist(e models.Event) (bool, error) {
	for _, v := range s.Events {
		if v.StartDate == e.StartDate && v.EndDate == e.EndDate {
			return true, nil
		}
	}
	return false, nil
}

func (s *Storage) Create(e models.Event) error {
	s.mu.Lock()
	if ok, _ := s.IsEventExist(e); ok {
		return storage.ErrorDateBusy
	}
	s.Events[s.id] = e
	s.id++
	s.mu.Unlock()
	return nil
}

func (s *Storage) Update(id int, e models.Event) error {
	if _, InMap := s.Events[id]; InMap {
		s.mu.Lock()
		s.Events[id] = e
		s.mu.Unlock()
		return nil
	}
	return storage.ErrorEventNotExist
}

func (s *Storage) Delete(id int) error {
	if _, ok := s.Events[id]; ok {
		s.mu.Lock()
		delete(s.Events, id)
		s.mu.Unlock()
		return nil
	}
	return storage.ErrorEventNotExist
}

func (s *Storage) GetIntervalEvent(day string, n int) ([]models.Event, error) {
	out := make([]models.Event, 0)
	d1, err := fmtdate.Parse(fmtdate.DefaultDateFormat, day)
	if err != nil {
		return nil, storage.DateFormatError
	}
	d2 := d1.Add(time.Duration(n) * time.Hour)
	for _, v := range s.Events {
		i1, _ := fmtdate.Parse(fmtdate.DefaultDateFormat, v.StartDate)
		i2, _ := fmtdate.Parse(fmtdate.DefaultDateFormat, v.EndDate)
		if i1.After(d1) && i1.Before(d2) || i2.After(d1) && i2.Before(d2) {
			out = append(out, v)
		}
	}
	return out, nil
}

func (s *Storage) GetDayEvent(day string) ([]models.Event, error) {
	return s.GetIntervalEvent(day, 0)
}

func (s *Storage) GetWeekEvent(day string) ([]models.Event, error) {
	return s.GetIntervalEvent(day, 168)
}

func (s *Storage) GetMonthEvent(day string) ([]models.Event, error) {
	return s.GetIntervalEvent(day, 720)
}
