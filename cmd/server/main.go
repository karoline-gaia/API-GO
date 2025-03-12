package main

import (
	"net/http"

	"github.com/karoline-gaia/API-GO/configs"
	"github.com/karoline-gaia/API-GO/internal/entity"
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
