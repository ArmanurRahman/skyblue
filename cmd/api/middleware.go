package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ArmanurRahman/skyblue/internal/helpers"
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

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get(authorizationHeaderKey)
		if authorizationHeader == "" {
			log.Println("authorization header not provided")
			helpers.GenerateClientResponseJson(w, http.StatusBadRequest, "error", "authorization header not provided")
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			log.Println("invalid authorization header format")
			helpers.GenerateClientResponseJson(w, http.StatusBadRequest, "error", "invalid authorization header format")
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			log.Println("unsupported authorization type ")
			helpers.GenerateClientResponseJson(w, http.StatusBadRequest, "error", fmt.Sprintf("server doesn't support %s type authorization", authorizationType))
			return
		}
		accessKey := fields[1]
		_, err := app.TokenMaker.VerifyToken(accessKey)

		if err != nil {
			log.Println(err)
			helpers.GenerateClientResponseJson(w, http.StatusBadRequest, "error", "invalid token")
			return
		}

		next.ServeHTTP(w, r)
	})
}
