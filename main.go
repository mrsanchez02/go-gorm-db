package main

import (
	"github.com/mrsanchez02/go-gorm-db/model"
	"github.com/mrsanchez02/go-gorm-db/storage"
)

func main() {
	driver := storage.Postgres
	storage.New(driver)

	invoice := model.InvoiceHeader{
		Client: "Eddy Abreu",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 1},
			{ProductID: 2},
		},
	}

	storage.DB().Create(&invoice)
}
