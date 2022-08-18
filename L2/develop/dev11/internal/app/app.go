package app

import (
	"context"
	"github.com/dingowd/WB/L2/develop/dev11/internal/logger"
	"github.com/dingowd/WB/L2/develop/dev11/internal/storage"
	"github.com/dingowd/WB/L2/develop/dev11/models"
)

type App struct {
	Logg    logger.Logger
	Storage storage.Storage
}

func New(logger logger.Logger, storage storage.Storage) *App {
	return &App{
		Logg:    logger,
		Storage: storage,
	}
}

func (a *App) CreateEvent(ctx context.Context, e models.Event) error {
	var err error
	if err = a.Storage.Create(e); err != nil {
		a.Logg.Error(err.Error())
	}
	return err
}

func (a *App) UpdateEvent(ctx context.Context, id int, e models.Event) error {
	var err error
	if err = a.Storage.Update(id, e); err != nil {
		a.Logg.Error(err.Error())
	}
	return err
}

func (a *App) DeleteEvent(ctx context.Context, id int) error {
	var err error
	if err = a.Storage.Delete(id); err != nil {
		a.Logg.Error(err.Error())
	}
	return err
}

func (a *App) GetDayEvent(day string) ([]models.Event, error) {
	var err error
	var events []models.Event
	if events, err = a.Storage.GetDayEvent(day); err != nil {
		a.Logg.Error(err.Error())
	}
	return events, err
}

func (a *App) GetWeekEvent(day string) ([]models.Event, error) {
	var err error
	var events []models.Event
	if events, err = a.Storage.GetWeekEvent(day); err != nil {
		a.Logg.Error(err.Error())
	}
	return events, err
}

func (a *App) GetMonthEvent(day string) ([]models.Event, error) {
	var err error
	var events []models.Event
	if events, err = a.Storage.GetMonthEvent(day); err != nil {
		a.Logg.Error(err.Error())
	}
	return events, err
}
