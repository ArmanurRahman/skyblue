package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/ArmanurRahman/skyblue/internal/config"
	"golang.org/x/crypto/bcrypt"
)

var app *config.AppConfig

func NewHelper(a *config.AppConfig) {
	app = a
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func GenerateHashPasswors(pass string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pass), 12)
	return string(hashedPassword)
}

func GenerateClientResponseJson(w http.ResponseWriter, status int, message string) http.ResponseWriter {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, message)))
	return w
}
