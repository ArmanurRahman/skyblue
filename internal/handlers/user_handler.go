package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/ArmanurRahman/skyblue/internal/helpers"
	"github.com/ArmanurRahman/skyblue/internal/models"
)

type userJson struct {
	FirstName string `json:"firstName"`
	LstName   string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Word      string `json:"word"`
	Street    string `json:"street"`
	OtherInfo string `json:"otherInfo"`
}

func (m *Repository) RegistrationUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		helpers.ServerError(w, err)
		helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "error")
		return
	}

	req, _ := ioutil.ReadAll(r.Body)

	var postUser userJson
	err = json.Unmarshal(req, &postUser)
	if err != nil {
		log.Println(err)
		helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "error")
		return
	}
	address := models.Address{
		Country:   postUser.Country,
		City:      postUser.City,
		Word:      postUser.Word,
		Street:    postUser.Street,
		OtherInfo: postUser.OtherInfo,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
	}
	err = m.App.Validate.Struct(address)
	if err != nil {
		log.Println(err)
		helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "invalid parameter")
		return
	}

	user := models.User{
		FirstName: postUser.FirstName,
		LastName:  postUser.LstName,
		Phone:     postUser.Phone,
		Email:     postUser.Email,
		Password:  helpers.GenerateHashPasswors(postUser.Password),
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
		Address:   address,
	}

	err = m.App.Validate.Struct(user)
	if err != nil {
		log.Println(err)
		helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "invalid parameter")
		return
	}

	id, err := m.DB.InsetAddress(address)

	if err != nil {
		log.Println(err)
		helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "error")
		return
	}

	user.AddressId = id
	err = m.DB.InsetUser(user)
	if err != nil {
		log.Println(err)
		return
	}
	helpers.GenerateClientResponseJson(w, http.StatusOK, "success")
}

type loginJson struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (m *Repository) LoginUser(w http.ResponseWriter, r *http.Request) {
	req, _ := ioutil.ReadAll(r.Body)

	var loginUser loginJson
	err := json.Unmarshal(req, &loginUser)
	if err != nil {
		log.Println(err)
		helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "error")
		return
	}
	user, err := m.DB.Login(loginUser.Email)
	if err != nil {
		log.Println(err)
		helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "error")
		return
	}

	isValid, err := helpers.CheckPassword(user.Password, loginUser.Password)
	if err != nil {
		log.Println(err)
	}

	if !isValid {
		helpers.GenerateClientResponseJson(w, http.StatusNotFound, "invalid credintial")
		return
	}
	helpers.GenerateClientResponseJson(w, http.StatusOK, "success")
}
