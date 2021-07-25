package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

//NoSurf adds CSRF protection to every post request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.IsProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
