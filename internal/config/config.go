package config

import "log"

type AppConfig struct {
	IsProduction bool
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
}
