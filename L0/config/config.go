package config

type Config struct {
	Provider MessageProvider
	Cache    Cache
	Logger   LoggerConf
	DSN      string
	HTTPSrv  string
}

type MessageProvider struct {
	ProviderType string
	ProviderName string
}

type Cache struct {
	Size int
}

type LoggerConf struct {
	Level   string
	LogFile string
}

func NewConfig() Config {
	return Config{}
}

func Default() Config {
	return Config{
		Provider: MessageProvider{
			ProviderType: "HTTP",
			ProviderName: "",
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
