package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/dingowd/WB/weather/service/internal/config"
	"github.com/dingowd/WB/weather/service/internal/logger/lrus"
	"github.com/dingowd/WB/weather/service/internal/storage"
	"github.com/dingowd/WB/weather/service/internal/storage/postgres"

	"os"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./weather/service/config/config.toml", "Path to configuration file")
}

func main() {
	// Init config
	conf := config.NewConfig()
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		fmt.Fprintln(os.Stdout, "ошибка чтения toml файла "+err.Error()+", установка параметров по умолчанию")
		conf = config.Default()
	}
	// init logger
	logg := lrus.New()
	logg.SetLevel(conf.Logger.Level)
	file, err := os.OpenFile(conf.Logger.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err == nil {
		logg.SetOutput(file)
		defer file.Close()
	} else {
		logg.SetOutput(os.Stdout)
	}
	// init storage
	var store storage.Storage
	store = postgres.New(logg, conf.AppId)
	if err := store.Connect(context.Background(), conf.DSN); err != nil {
		logg.Error("failed to connect database" + err.Error())
		os.Exit(1) // nolint:gocritic
	}
	defer func() {
		logg.Info("Закрытие соединения с БД")
		store.Close()
	}()
	store.GetWeather()
}
