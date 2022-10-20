package app

import (
	"github.com/dingowd/WB/weather/service/internal/logger"
	"github.com/dingowd/WB/weather/service/internal/storage"
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
