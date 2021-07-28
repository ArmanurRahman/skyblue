package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ArmanurRahman/skyblue/internal/helpers"
	"github.com/ArmanurRahman/skyblue/internal/models"
)

func (m *Repository) SearchProduct(w http.ResponseWriter, r *http.Request) {

}

type ProductCategoryJson struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	Url             string `json:"url"`
	CategoryId      string `json:"categoryId"`
	CategoryName    string `json:"categoryName"`
	CategoryDetails string `json:"categoryDetails"`
	Type            string `json:"type"`
	OtherInfo       string `json:"otherInfo"`
}

func (m *Repository) AddProduct(w http.ResponseWriter, r *http.Request) {
	req, _ := ioutil.ReadAll(r.Body)

	var postProduct ProductCategoryJson

	err := json.Unmarshal(req, &postProduct)
	if err != nil {
		log.Println(err)
		helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "error", "server error1")
		return
	}

	var categoryId int
	var category models.Category

	if postProduct.CategoryId == "" {
		category = models.Category{
			Name:      postProduct.CategoryName,
			Details:   postProduct.CategoryDetails,
			Type:      postProduct.Type,
			OtherInfo: postProduct.OtherInfo,
			CreateAt:  time.Now(),
			UpdateAt:  time.Now(),
		}

		err = m.App.Validate.Struct(category)
		if err != nil {
			log.Println(err)
			helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "error", "invalid parameter")
			return
		}

		categoryId, err = m.DB.InsetCategory(category)
		if err != nil {
			log.Println(err)
			helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "error", "can't insert category")
			return
		}
	} else {
		categoryId, _ = strconv.Atoi(postProduct.CategoryId)
	}

	product := models.Product{
		Name:        postProduct.Name,
		Description: postProduct.Description,
		CategoryId:  categoryId,
		Url:         postProduct.Url,
		CreateAt:    time.Now(),
		UpdateAt:    time.Now(),
	}

	err = m.App.Validate.Struct(product)
	if err != nil {
		log.Println(err)
		helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "error", "invalid parameter")
		return
	}

	err = m.DB.InsetProduct(product)
	if err != nil {
		log.Println(err)
		helpers.GenerateClientResponseJson(w, http.StatusInternalServerError, "error", "can't insert product")
		return
	}

	helpers.GenerateClientResponseJson(w, http.StatusOK, "success", "product registration successfully")

}
