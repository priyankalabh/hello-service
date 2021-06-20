package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HelloHandler handles the /hello API call
func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
