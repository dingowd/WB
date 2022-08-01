package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/dingowd/WB/L0/app"
	"github.com/dingowd/WB/L0/config"
	"github.com/dingowd/WB/L0/logger"
	"github.com/dingowd/WB/L0/logger/lrus"
	"github.com/dingowd/WB/L0/model"
	"github.com/dingowd/WB/L0/storage"
	"github.com/dingowd/WB/L0/storage/postgres"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/BurntSushi/toml"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./config/config.toml", "Path to configuration file")
}

func main() {
	flag.Parse()
	// Init config
	conf := config.NewConfig()
	if _, err := toml.DecodeFile(configFile, &conf); err != nil {
		fmt.Fprintln(os.Stdout, "error decoding toml"+err.Error()+" setting default values")
		conf = config.Default()
	}
	// init logger
	logg := lrus.New()
	logg.SetLevel(conf.Logger.Level)
	file, err := os.OpenFile("./log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
	if err == nil {
		logg.SetOutput(file)
		defer file.Close()
	} else {
		logg.SetOutput(os.Stdout)
	}
	// init storage
	var stor storage.Storage
	stor = postgres.New(logg)
	if err := stor.Connect(context.Background(), conf.DSN); err != nil {
		logg.Error("failed to connect database" + err.Error())
		os.Exit(1) // nolint:gocritic
	}
	defer stor.Close()

	// init application
	service := app.New(logg, stor)

}
