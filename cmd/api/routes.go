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

	mux.Post("/user-registration", handlers.Repo.PostUser)
	mux.Post("/saler-registration", handlers.Repo.PostSaler)

	mux.Post("/user/login", handlers.Repo.Login)
	mux.Route("/user", func(mux chi.Router) {

	})
	return mux
}
