package app

import (
	"github.com/dingowd/WB/L0/logger"
	"github.com/dingowd/WB/L0/storage"
)

type App struct {
	Log   logger.Logger
	Store storage.Storage
}

func New(logger logger.Logger, storage storage.Storage) *App {
	return &App{
		Log:   logger,
		Store: storage,
	}
}
