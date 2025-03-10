package database

import "github.com/karoline-gaia/API-GO/internal/entity"

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type PoductInterface interface {
	CreateProduct(product *entity.Product) error
	FindAll(page, limt int, sort string) ([]entity.Product, error)
	FindById(id string) (*entity.Product, error)
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id string) error
}
