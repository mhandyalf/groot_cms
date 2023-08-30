package models

type DataStore struct {
	StoreID    int    `json:"store_id"`
	StoreEmail string `json:"store_email"`
	Password   string `json:"password"`
	StoreName  string `json:"store_name"`
	StoreType  string `json:"store_type"`
}
