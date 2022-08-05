package app

import (
	"github.com/dingowd/WB/L0/cache"
	"github.com/dingowd/WB/L0/logger"
	"github.com/dingowd/WB/L0/storage"
)

type App struct {
	Log   logger.Logger
	Store storage.Storage
	Cache cache.CacheInterface
}

func New(logger logger.Logger, storage storage.Storage, cache cache.CacheInterface) *App {
	return &App{
		Log:   logger,
		Store: storage,
		Cache: cache,
	}
}
