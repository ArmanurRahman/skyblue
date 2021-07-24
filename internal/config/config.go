package config

import (
	"log"

	"github.com/go-playground/validator"
)

type AppConfig struct {
	IsProduction bool
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	Validate     *validator.Validate
}
