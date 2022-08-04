package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/dingowd/WB/L0/app"
	"github.com/dingowd/WB/L0/config"
	"github.com/dingowd/WB/L0/logger/lrus"
	"github.com/dingowd/WB/L0/storage"
	"github.com/dingowd/WB/L0/storage/postgres"
	"github.com/dingowd/WB/L0/subscriber"
	natsstream "github.com/dingowd/WB/L0/subscriber/nats_stream"
	"os"

	"github.com/BurntSushi/toml"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "C:/Users/dingowd/go/WB/L0/config/config.toml", "Path to configuration file")
}

/*type app struct {
	Log logger.Logger
	Stor storage.Storage
	Prov provider.MessageProvider
}*/

func main() {
	/*	flag.Parse()
		var a app*/
	// Init config
	conf := config.NewConfig()
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		fmt.Fprintln(os.Stdout, "error decoding toml "+err.Error()+" setting default values")
		conf = config.Default()
	}
	_ = fmt.Sprint(123)
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

	// Init app
	app := app.New(logg, store)
	conf.Subscriber.App = app

	//init subscriber
	var sub subscriber.Subscriber
	sub = natsstream.NewSub(conf.Subscriber)
	sub.Start()

}
