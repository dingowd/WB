package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/dingowd/WB/weather/service/internal/app"
	"github.com/dingowd/WB/weather/service/internal/config"
	"github.com/dingowd/WB/weather/service/internal/logger/lrus"
	internalhttp "github.com/dingowd/WB/weather/service/internal/server/http"
	"github.com/dingowd/WB/weather/service/internal/storage"
	"github.com/dingowd/WB/weather/service/internal/storage/postgres"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	// init application
	weather := app.New(logg, store)

	// init http server
	server := internalhttp.NewServer(weather, conf.HTTPSrv)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-exit
		logg.Info("Weather service stopping...")
		server.Stop()
		logg.Info("Weather service stopped")
		time.Sleep(5 * time.Second)
	}()
	logg.Info("Weather service is running...")

	server.Start()
}
