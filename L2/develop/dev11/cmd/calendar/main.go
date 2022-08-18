package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/dingowd/WB/L2/develop/dev11/internal/app"
	"github.com/dingowd/WB/L2/develop/dev11/internal/logger/lrus"
	internalhttp "github.com/dingowd/WB/L2/develop/dev11/internal/server/http"
	storage "github.com/dingowd/WB/L2/develop/dev11/internal/storage"
	sqlstorage "github.com/dingowd/WB/L2/develop/dev11/internal/storage/sql"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}
	// init config
	config := NewConfig()
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		fmt.Fprintln(os.Stdout, "error decoding toml"+err.Error()+" setting default values")
		config = Default()
	}

	// init logger
	logg := lrus.New()
	logg.SetLevel(config.Logger.Level)
	file, err := os.OpenFile(config.Logger.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err == nil {
		logg.SetOutput(file)
		defer file.Close()
	} else {
		logg.SetOutput(os.Stdout)
	}

	// init storage
	var storage storage.Storage
	storage = sqlstorage.New(logg)
	if err := storage.Connect(context.Background(), config.DSN); err != nil {
		logg.Error("failed to connect database" + err.Error())
		os.Exit(1) // nolint:gocritic
	}
	defer storage.Close()

	// init application
	calendar := app.New(logg, storage)

	// init http server
	server := internalhttp.NewServer(calendar, config.HTTPSrv)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-exit
		logg.Info("Calendar stopping...")
		server.Stop()
		logg.Info("Calendar stopped")
		time.Sleep(5 * time.Second)
	}()
	logg.Info("calendar is running...")

	server.Start()
}
