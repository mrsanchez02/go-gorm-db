# Readme

## Migrations

  ```go
  driver := storage.MySQL
  storage.New(driver)
  
  storage.DB().AutoMigrate(&model.Product{}, &model.InvoiceHeader{}, &model.InvoiceItem{})
  ```

## Create

  ```go
  driver := storage.postgres
  storage.New(driver)
  
  product1 := model.Product{Name: "Curso de Go 2023", Price: 120}
  obs := "Testing with Go"
  product2 := model.Product{Name: "Curso de Testing con Go", Price: 150, Observations: &obs}
  product3 := model.Product{Name: "Curso de Python", Price: 200}
  
  storage.DB().Create(&product1)
  storage.DB().Create(&product2)
  storage.DB().Create(&product3)
  ```

## Query All

  ```go
  driver := storage.Postgres
  storage.New(driver)

  products := make([]model.Product, 0)
  storage.DB().Find(&products)

  for _, product := range products {
    fmt.Printf("%d - %s\n", product.ID, product.Name)
  }
  ```
