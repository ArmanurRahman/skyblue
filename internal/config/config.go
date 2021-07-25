package config

import (
	"log"

	"github.com/ArmanurRahman/skyblue/internal/token"
	"github.com/go-playground/validator"
)

type AppConfig struct {
	IsProduction bool
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	Validate     *validator.Validate
	TokenMaker   token.Maker
}
