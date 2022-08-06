package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/dingowd/WB/L0/app"
	с "github.com/dingowd/WB/L0/cache"
	"github.com/dingowd/WB/L0/config"
	"github.com/dingowd/WB/L0/logger/lrus"
	"github.com/dingowd/WB/L0/server"
	"github.com/dingowd/WB/L0/storage"
	"github.com/dingowd/WB/L0/storage/postgres"
	"github.com/dingowd/WB/L0/subscriber"
	natsstream "github.com/dingowd/WB/L0/subscriber/nats_stream"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BurntSushi/toml"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./L0/config/config.toml", "Path to configuration file")
}

func main() {
	// Init config
	conf := config.NewConfig()
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		fmt.Fprintln(os.Stdout, "error decoding toml "+err.Error()+" setting default values")
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
	store = postgres.New(logg)
	if err := store.Connect(context.Background(), conf.DSN); err != nil {
		logg.Error("failed to connect database" + err.Error())
		os.Exit(1) // nolint:gocritic
	}
	store.Connect(context.Background(), conf.DSN)
	defer store.Close()

	// Init cache
	var cache с.CacheInterface
	cache = с.NewCache(logg, store, conf.Cache.Size)
	cache.Init()

	// Init app
	app := app.New(logg, store, cache)
	conf.Subscriber.App = app

	// init http server
	server := server.NewServer(*app, conf.HTTPSrv)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	stopChan := make(chan struct{})

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
		stopChan <- struct{}{}
	}()
	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1)
	}

	//init subscriber
	var sub subscriber.Subscriber
	sub = natsstream.NewSub(conf.Subscriber)
	sub.Start(stopChan)
}
