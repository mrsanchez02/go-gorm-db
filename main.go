package main

import (
	"fmt"

	"github.com/mrsanchez02/go-gorm-db/model"
	"github.com/mrsanchez02/go-gorm-db/storage"
)

func main() {
	driver := storage.MySQL
	storage.New(driver)

	myProduct := model.Product{}

	storage.DB().First(&myProduct, 1)
	fmt.Println(myProduct)
}
