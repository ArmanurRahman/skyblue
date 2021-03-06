package repository

import "github.com/ArmanurRahman/skyblue/internal/models"

type DatabaseRepo interface {
	InsetAddress(addr models.Address) (int, error)
	InsetUser(user models.User) error
	InsetSaler(saler models.Saler) error
	Login(email string) (models.User, error)
	InsetProduct(product models.Product) error
	InsetCategory(catetory models.Category) (int, error)
}
