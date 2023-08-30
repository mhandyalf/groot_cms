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

func (db *DB) GetStoreByEmail(email string) (*models.DataStore, error) {
	store := &models.DataStore{}

	err := db.Where("store_email = ?", email).First(store).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err // Store not found
		}
		return nil, err
	}

	return store, nil
}

func (db *DB) CreateStore(store *models.DataStore) error {
	err := db.Create(store).Error
	if err != nil {
		return err
	}

	return nil
}
