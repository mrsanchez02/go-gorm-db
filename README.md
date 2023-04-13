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

## Query one by id

  ```go
  driver := storage.MySQL
  storage.New(driver)
  
  myProduct := model.Product{}
  
  storage.DB().First(&myProduct, 1)
  fmt.Println(myProduct)
  ```

## Update

  ```go
  driver := storage.Postgres
  storage.New(driver)
  
  myProduct := model.Product{}
  myProduct.ID = 3
  
  storage.DB().Model(&myProduct).Updates(
    model.Product{Name: "Curso de Java", Price: 120},
  )
  ```
## Delete Soft

  ```go
  driver := storage.Postgres
  storage.New(driver)
  
  myProduct := model.Product{}
  myProduct.ID = 5
  
  storage.DB().Delete(&myProduct)
  ```

## Delete Permanent

  ```go
  driver := storage.Postgres
  storage.New(driver)
  
  myProduct := model.Product{}
  myProduct.ID = 5
  
  storage.DB().Unscoped().Delete(&myProduct)
  ```

## Transactions

  ```go
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
  ```
