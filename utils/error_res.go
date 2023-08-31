package utils

import "github.com/gin-gonic/gin"

type APIErrors struct {
	StatusCode          int
	ResponseCode        string `json:"responseCode"`
	ResponseDescription string `json:"responseDescription"`
}

func ErrorMessage(c *gin.Context, apiError *APIErrors) *gin.Context {
	c.Abort()
	c.JSON(apiError.StatusCode, gin.H{
		"responseCode":        apiError.ResponseCode,
		"responseDescription": apiError.ResponseDescription,
	})
	return c
}
