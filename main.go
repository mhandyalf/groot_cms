package main

import (
	"groot_cms/database"
	"groot_cms/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	router := setupRouter(db)

	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)
}

func setupRouter(db *database.DB) http.Handler {
	router := gin.Default()

	authHandler := handlers.NewAuthHandler(db)

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	return router
}
