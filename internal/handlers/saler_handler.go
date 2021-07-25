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

type salerJson struct {
	Name      string `json:"name"`
	Details   string `json:"details"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Phone     string `json:"phone"`
	Country   string `json:"country"`
	City      string `json:"city"`
	Word      string `json:"word"`
	Street    string `json:"street"`
	OtherInfo string `json:"otherInfo"`
}

func (m *Repository) RegistrationSaler(w http.ResponseWriter, r *http.Request) {
	//read all data from post api
	req, _ := ioutil.ReadAll(r.Body)

	//decode request json
	var salerParam salerJson
	err := json.Unmarshal(req, &salerParam)
	if err != nil {
		log.Println(err)
		helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "error")
		return
	}

	address := models.Address{
		Country:   salerParam.Country,
		City:      salerParam.City,
		Word:      salerParam.Word,
		Street:    salerParam.Street,
		OtherInfo: salerParam.OtherInfo,
		CreateAt:  time.Now(),
		UpdateAt:  time.Now(),
	}

	//validate request
	err = m.App.Validate.Struct(address)
	if err != nil {
		log.Println(err)
		helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "invalid parameter")
		return
	}

	saler := models.Saler{
		Name:     salerParam.Name,
		Details:  salerParam.Details,
		Phone:    salerParam.Phone,
		Email:    salerParam.Email,
		Password: helpers.GenerateHashPasswors(salerParam.Password),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
		Address:  address,
	}
	//validate request
	err = m.App.Validate.Struct(saler)
	if err != nil {
		log.Println(err)
		helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "invalid parameter")
		return
	}

	id, err := m.DB.InsetAddress(address)
	if err != nil {
		log.Println(err)
		helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "invalid parameter")
		return
	}

	saler.AddressId = id
	err = m.DB.InsetSaler(saler)
	if err != nil {
		log.Println(err)
		helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "invalid parameter")
		return
	}

	helpers.GenerateClientResponseJson(w, http.StatusOK, "success")
}
