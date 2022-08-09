package config

import (
	natsstream "github.com/dingowd/WB/L0/subscriber/nats_stream"
	"github.com/nats-io/stan.go"
)

type Config struct {
	Subscriber natsstream.NatsStream
	Cache      Cache
	Logger     LoggerConf
	DSN        string
	HTTPSrv    string
}

type Subscriber struct {
	ClusterID   string
	ClientID    string
	URL         string
	UserCreds   string
	ShowTime    bool
	Qgroup      string
	Unsubscribe bool
	StartSeq    uint64
	StartDelta  string
	DeliverAll  bool
	NewOnly     bool
	DeliverLast bool
	Durable     string
	Subj        string
}

type Cache struct {
	Size int
}

type LoggerConf struct {
	Level   string
	LogFile string
}

func NewConfig() *Config {
	return &Config{}
}

func Default() *Config {
	return &Config{
		Subscriber: natsstream.NatsStream{
			ClusterID:   "WB_Cluster",
			ClientID:    "wbClient",
			URL:         stan.DefaultNatsURL,
			UserCreds:   "",
			ShowTime:    false,
			Qgroup:      "",
			Unsubscribe: false,
			StartSeq:    0,
			StartDelta:  "",
			DeliverAll:  true,
			NewOnly:     true,
			DeliverLast: false,
			Durable:     "",
			Subj:        "order",
		},

		Cache: Cache{
			Size: 25,
		},
		Logger: LoggerConf{
			Level:   "INFO",
			LogFile: "./log.txt",
		},
		DSN:     "user=postgres dbname=WB sslmode=disable password=masterkey",
		HTTPSrv: "127.0.0.1:3541",
	}
}
