package database

import "github.com/karoline-gaia/API-GO/internal/entity"

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
