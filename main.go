package main

import "github.com/mrsanchez02/go-gorm-db/storage"

func main() {
	driver := storage.MySQL
	storage.New(driver)
}
