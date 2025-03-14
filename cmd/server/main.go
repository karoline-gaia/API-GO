package main

import (
	"net/http"

	"github.com/goccy/go-json"
	"github.com/karoline-gaia/API-GO/configs"
	"github.com/karoline-gaia/API-GO/internal/dto"
	"github.com/karoline-gaia/API-GO/internal/entity"
	"github.com/karoline-gaia/API-GO/internal/infra/database"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("teste.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	http.ListenAndServe(":8000", nil)
}

type ProductHandler struct {
	ProductDB database.PoductInterface
}

func NewProductHandler(db database.PoductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (ph *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = ph.ProductDB.CreateProduct(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
