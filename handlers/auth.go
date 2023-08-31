package handlers

import (
	"fmt"
	"groot_cms/models"
	"groot_cms/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db *gorm.DB
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (ah *AuthHandler) Register(c *gin.Context) {
	var user models.DataStore
	if err := c.ShouldBindJSON(&user); err != nil {
		contractErrorResponse(c, http.StatusInternalServerError, "Terjadi kesalahan internal server.", err)
		utils.ErrorMessage(c, &utils.ErrBindingJSON)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = string(hashedPassword)

	if err := ah.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.SuccessMessage(c, http.StatusOK, "Success register")
}

func (ah *AuthHandler) Login(c *gin.Context) {
	var loginRequest struct {
		Email    string `json:"store_email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		contractErrorResponse(c, http.StatusInternalServerError, "Terjadi kesalahan internal server.", err)
		utils.ErrorMessage(c, &utils.ErrBindingJSON)
		fmt.Printf("[PaymentControlle.GetData] error "+"when get data from db : %v\n", err)
		return
	}

	var user models.DataStore
	if err := ah.db.Where("store_email = ?", loginRequest.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorMessage(c, &utils.ErrorGetData)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		utils.ErrorMessage(c, &utils.ErrorGetData)
		return
	}

	token, err := generateJWTToken(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
	utils.SuccessWithData(c, http.StatusOK, user)
}

func generateJWTToken(user *models.DataStore) (string, error) {
	claims := jwt.MapClaims{
		"email": user.StoreEmail,
		"type":  user.StoreType,
		// Add more claims as needed
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := []byte("your-secret-key") // Replace with your actual secret key
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func contractErrorResponse(c *gin.Context, status int, message string, err error) {
	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"error":   err.Error(),
	})
}
