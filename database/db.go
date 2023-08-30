package database

import (
	"database/sql"
	"groot_cms/models"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func InitDB() *DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/avenger")
	if err != nil {
		panic(err)
	}

	return &DB{db}
}

func (db *DB) GetStoreByEmail(email string) (*models.DataStore, error) {
	query := "SELECT store_id, store_email, password, store_name, store_type FROM DataStore WHERE store_email = ?"
	store := &models.DataStore{}

	err := db.QueryRow(query, email).Scan(&store.StoreID, &store.StoreEmail, &store.Password, &store.StoreName, &store.StoreType)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err // Store not found
		}
		return nil, err
	}

	return store, nil
}

func (db *DB) CreateStore(store *models.DataStore) error {
	query := "INSERT INTO DataStore (store_email, password, store_name, store_type) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, store.StoreEmail, store.Password, store.StoreName, store.StoreType)
	if err != nil {
		return err
	}

	return nil
}
