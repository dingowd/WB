package main

// При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.
type Config struct {
	Logger      LoggerConf
	DSN         string
	HTTPSrv     string
	Storagetype string
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
		Logger: LoggerConf{
			Level:   "INFO",
			LogFile: "./log.txt",
		},
		DSN:         "user=postgres dbname=calendar sslmode=disable password=masterkeyr",
		HTTPSrv:     "127.0.0.1:3541",
		Storagetype: "memory",
	}
}
