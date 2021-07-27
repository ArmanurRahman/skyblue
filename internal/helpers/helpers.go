package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/ArmanurRahman/skyblue/internal/config"
	"github.com/ArmanurRahman/skyblue/internal/models"
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

func GenerateClientResponseJson(w http.ResponseWriter, status int, result string, message string) {
	returnJson := models.GetResponseJson{
		Result:  result,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(returnJson)
}

func GenerateClientResponseWithPayloadJson(w http.ResponseWriter, status int, message string, payload interface{}, result string) {
	returnJson := models.PostResponseJson{
		Result:  result,
		Message: message,
		Data:    payload,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(returnJson)

}

func CheckPassword(hashedPassword, userPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, errors.New("incorrect password")
	} else if err != nil {
		return false, err
	}

	return true, nil
}
