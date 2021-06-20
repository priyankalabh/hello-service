package main

import (
	"hello-service/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	router.AttachRouter(server)
	err := server.Run(":5050")
	if err != nil {
		log.Fatalf("Run error: %v", err)
	}

}
