package repository

import "github.com/ArmanurRahman/skyblue/internal/models"

type DatabaseRepo interface {
	InsetAddress(addr models.Address) (int, error)
	InsetUser(user models.User) error
	InsetSaler(saler models.Saler) error
}
