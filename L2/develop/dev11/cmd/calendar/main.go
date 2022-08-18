package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/dingowd/WB/L2/develop/dev11/internal/logger/lrus"
	storage2 "github.com/dingowd/WB/L2/develop/dev11/internal/storage"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/dingowd/WB/L2/develop/dev11/internal/app"
	internalhttp "github.com/dingowd/WB/L2/develop/dev11/internal/server/http"
	memorystorage "github.com/dingowd/WB/L2/develop/dev11/internal/storage/memory"
	sqlstorage "github.com/dingowd/WB/L2/develop/dev11/internal/storage/sql"
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
	var storage storage2.Storage
	switch config.Storagetype {
	case "memory":
		storage = memorystorage.New()
	case "sql":
		storage = sqlstorage.New()
		if err := storage.Connect(context.Background(), config.DSN); err != nil {
			logg.Error("failed to connect database" + err.Error())
			os.Exit(1) // nolint:gocritic
		}
	default:
		storage = memorystorage.New()
	}
	defer storage.Close()

	// init application
	calendar := app.New(logg, storage)

	// init http server
	server := internalhttp.NewServer(calendar, config.HTTPSrv)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1)
	}
}
