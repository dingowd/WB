package app

import (
	"github.com/dingowd/WB/L0/model"
)

type App struct {
	Logg    model.Logger
	Storage model.Storage
}

func New(logger model.Logger, storage model.Storage) *App {
	return &App{
		Logg:    logger,
		Storage: storage,
	}
}
