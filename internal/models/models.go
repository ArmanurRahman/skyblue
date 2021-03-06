package models

import (
	"time"
)

type Address struct {
	ID        int
	Country   string `validate:"required"`
	City      string `validate:"required"`
	Word      string `validate:"required"`
	Street    string
	OtherInfo string
	CreateAt  time.Time
	UpdateAt  time.Time
}

type Saler struct {
	ID        int
	Name      string `validate:"required"`
	Details   string `validate:"required"`
	Phone     string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	AddressId int
	CreateAt  time.Time
	UpdateAt  time.Time
	Address   Address
}

type User struct {
	ID        int
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Phone     string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	AddressId int
	CreateAt  time.Time
	UpdateAt  time.Time
	Address   Address
}

type Category struct {
	ID        int
	Name      string `validate:"required"`
	Details   string
	Type      string `validate:"required"`
	OtherInfo string
	CreateAt  time.Time
	UpdateAt  time.Time
}

type Product struct {
	ID          int
	Name        string `validate:"required"`
	Description string
	CategoryId  int `validate:"required"`
	CreateAt    time.Time
	UpdateAt    time.Time
	Category    *Category `validate:"omitempty"`
	Url         string    `validate:"required"`
}

type AvailableProduct struct {
	ID          int
	ProductId   int
	SalerId     int
	Price       int
	Quantity    int
	ArrivalDate time.Time
	CreateAt    time.Time
	UpdateAt    time.Time
	Product     Product
	Saler       Saler
}

type Order struct {
	ID                 int
	UserId             int
	AvailableProductId int
	OrderDate          time.Time
	CategoryId         int
	Status             string
	Quantity           int
	DeliveryDate       time.Time
	DeliveryFlowJson   string
	CreateAt           time.Time
	UpdateAt           time.Time
	User               User
	AvailableProduct   AvailableProduct
}

type Comment struct {
	ID        int
	ProductId int
	UserId    int
	Body      string
	CreateAt  time.Time
	UpdateAt  time.Time
	User      User
	Product   Product
}

type Rating struct {
	ID        int
	ProductId int
	UserId    int
	Rate      float32
	CreateAt  time.Time
	UpdateAt  time.Time
	User      User
	Product   Product
}

type WishList struct {
	ID           int
	UserId       int
	WishlistNmae string
	CreateAt     time.Time
	UpdateAt     time.Time
	User         User
}

type WishedProduct struct {
	ID         int
	WishlistId int
	ProductId  int
	CreateAt   time.Time
	UpdateAt   time.Time
	Wishlist   WishList
	Product    Product
}

type PostResponseJson struct {
	Result  string      `json:"result"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type GetResponseJson struct {
	Result  string
	Message string
}

type LoginRes struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
