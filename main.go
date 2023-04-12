package main

import (
	"github.com/mrsanchez02/go-gorm-db/model"
	"github.com/mrsanchez02/go-gorm-db/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	product1 := model.Product{Name: "Curso de Go 2023", Price: 120}
	obs := "Testing with Go"
	product2 := model.Product{Name: "Curso de Testing con Go", Price: 150, Observations: &obs}
	product3 := model.Product{Name: "Curso de Python", Price: 200}

	storage.DB().Create(&product1)
	storage.DB().Create(&product2)
	storage.DB().Create(&product3)
}
