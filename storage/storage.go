package storage

import (
	"fmt"
	"log"

	"github.com/SKilliu/gogql/graph/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type QInterface interface {
	DB() *gorm.DB

	UsersQ() UsersQ
}

type DB struct {
	db *gorm.DB
}

func (d DB) DB() *gorm.DB {
	return d.db
}

func InitConnection() (QInterface, error) {
	dsn := "host=localhost user=postgres password=1234567 dbname=service port=5430 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect to DB")
		return nil, err
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to DB on localhost:5432")

	return &DB{db: db}, err
}
