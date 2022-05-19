package storage

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitConnection() error {
	dsn := "host=localhost user=postgres password=1234567 dbname=service port=5430 sslmode=disable TimeZone=Asia/Shanghai"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect to DB")
		return err
	}

	fmt.Println("Successfully connected to DB on localhost:5432")

	return err
}
