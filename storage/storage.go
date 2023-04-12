package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Driver of storage
type Driver string

const (
	MySQL    Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New create the connection with db
func New(d Driver) {
	switch d {
	case "MYSQL":
		newMySQLDB()
	case "POSTGRES":
		newPostgresDB()
	default:

	}
}

func newPostgresDB() {
	once.Do(func() {
		var err error
		dsn := "postgres://leandrosc:leandrosc@localhost:5432/gormdb?sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		fmt.Println("Connected to Postgres")
	})

}

func newMySQLDB() {
	once.Do(func() {
		var err error
		dsn := "leandrosc:leandrosc@tcp(localhost:3306)/gormdb?parseTime=true"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Can't open db: %v", err)
		}

		fmt.Println("Connected to Mysql")
	})
}

// DB return a unique instancie of db
func DB() *gorm.DB {
	return db
}
