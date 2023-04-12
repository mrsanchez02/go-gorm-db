package main

import (
	"github.com/mrsanchez02/go-gorm-db/model"
	"github.com/mrsanchez02/go-gorm-db/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	storage.DB().AutoMigrate(&model.Product{}, &model.InvoiceHeader{}, &model.InvoiceItem{})
}
