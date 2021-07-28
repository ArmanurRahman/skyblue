package handlers

import (
	"fmt"
	"net/http"

	"github.com/ArmanurRahman/skyblue/internal/config"
	"github.com/ArmanurRahman/skyblue/internal/drivers"
	"github.com/ArmanurRahman/skyblue/repository"
	"github.com/ArmanurRahman/skyblue/repository/dbrepo"
)

var Repo *Repository

//Repository is the repository type
type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig, db *drivers.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewPostgressDBRepo(db.SQL, a),
	}
}

//NewHandlers sets the repository for the handlers
func NewHandler(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func (m *Repository) PostUser(w http.ResponseWriter, r *http.Request) {
	m.RegistrationUser(w, r)
}

func (m *Repository) PostSaler(w http.ResponseWriter, r *http.Request) {
	m.RegistrationSaler(w, r)
}

func (m *Repository) Login(w http.ResponseWriter, r *http.Request) {
	m.LoginUser(w, r)
}

func (m *Repository) GetCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func (m *Repository) GetAllProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func (m *Repository) ProductDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}
