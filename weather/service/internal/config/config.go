package config

type Config struct {
	Logger   LoggerConf
	DSN      string
	HTTPSrv  string
	AppId    string
	Interval int
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
		Logger: LoggerConf{
			Level:   "INFO",
			LogFile: "./log.txt",
		},
		DSN:      "postgres://postgres:masterkey@postgresdb:5432/weather?sslmode=disable",
		HTTPSrv:  ":3541",
		AppId:    "effe2454b05e87a8f799b78681cbd4e1",
		Interval: 60,
	}
}
