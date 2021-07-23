package main

import (
	"net/http"

	"github.com/ArmanurRahman/skyblue/internal/config"
	"github.com/ArmanurRahman/skyblue/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func Routes(a *config.AppConfig) http.Handler {
	mux := chi.NewRouter()
	//mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)

	return mux
}
