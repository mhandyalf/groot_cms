package main

import (
	"groot_cms/database"
	"groot_cms/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()

	router := gin.Default()

	authHandler := handlers.NewAuthHandler(db)

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)
}
