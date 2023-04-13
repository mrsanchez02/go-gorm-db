package main

import (
	"github.com/mrsanchez02/go-gorm-db/model"
	"github.com/mrsanchez02/go-gorm-db/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	myProduct := model.Product{}
	myProduct.ID = 3

	storage.DB().Model(&myProduct).Updates(
		model.Product{Name: "Curso de Java", Price: 120},
	)
}
