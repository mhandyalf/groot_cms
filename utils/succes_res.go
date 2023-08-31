package utils

import "github.com/gin-gonic/gin"

// SuccessMessage ...
func SuccessMessage(c *gin.Context, status int, msg string) *gin.Context {
	c.JSON(status, gin.H{
		"responseCode":        "0000",
		"responseDescription": "Success",
	})
	return c
}

// SuccessWithData ...
func SuccessWithData(c *gin.Context, status int, data interface{}) *gin.Context {
	c.JSON(status, gin.H{
		"responseCode":        "0000",
		"responseDescription": "Success",
		"data":                data,
	})

	return c
}
