package entity

import (
	"errors"
	"time"

	"github.com/karoline-gaia/API-GO/pkg/entity"
)

var (
	ErrIDIsRequired    = errors.New("id is required")
	ErrInbvalidID      = errors.New("id is invalid")
	ErrNameIsRequired  = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrInvalidPrice    = errors.New("price is invalid")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price int) (*Product, error) {
	product := &Product{
		ID:        entity.NewID(),
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}
	if err := product.Validate(); err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInbvalidID
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	if p.Price < 0 {
		return ErrInvalidPrice
	}
	return nil
}
