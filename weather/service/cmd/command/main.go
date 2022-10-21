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
	flag.StringVar(&configFile, "config", "./config/config.toml", "Path to configuration file")
}

func main() {
	// graceful shutdown
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	// init config
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
	defer store.Close()
	// getting weather from API with interval in seconds
	go func() {
		for {
			select {
			case <-ctx.Done():
				break
			default:
				logg.Info("Getting weather from API")
				store.GetWeather()
				logg.Info("Got weather from API")
				time.Sleep(time.Second * time.Duration(conf.Interval))
			}
		}
	}()

	// init application
	weather := app.New(logg, store)

	// init http server
	server := internalhttp.NewServer(weather, conf.HTTPSrv)

	// graceful shutdown
	go func() {
		<-ctx.Done()
		logg.Info("Weather service stopping...")
		server.Stop()
		logg.Info("Weather service stopped")
		time.Sleep(5 * time.Second)
	}()
	logg.Info("Weather service is running...")

	// start http server
	server.Start()
}
