package main

import (
	"github.com/mrsanchez02/go-gorm-db/model"
	"github.com/mrsanchez02/go-gorm-db/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	// Delete Soft
	// myProduct := model.Product{}
	// myProduct.ID = 5

	// storage.DB().Delete(&myProduct)

	// Delete Permanent
	myProduct := model.Product{}
	myProduct.ID = 5

	storage.DB().Unscoped().Delete(&myProduct)

}
