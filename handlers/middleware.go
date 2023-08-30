package handlers

import (
	"groot_cms/utils"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		utils.LogRequest(c.Request.Method, c.Request.URL.Path, start)
	}
}
