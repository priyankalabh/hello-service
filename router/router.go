package router

import (
	"hello-service/handlers"

	"github.com/gin-gonic/gin"
)

func AttachRouter(router *gin.Engine) {
	router.GET("/api/hello", handlers.HelloHandler)
}
