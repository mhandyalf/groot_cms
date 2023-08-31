package database

import (
	"groot_cms/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	*gorm.DB
}

func InitDB() *gorm.DB {
	dsn := "user=postgres password=123456 dbname=postgres host=localhost port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.DataStore{})
	if err != nil {
		panic(err)
	}

	return db
}
