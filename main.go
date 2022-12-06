package main

import (
	"log"
	"party/config"
	"party/person/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Migration()

	r := gin.New()

	routes.Routes(r.Group("/api"))

	err := r.Run("localhost:3000")

	if err!= nil {
        log.Fatal("Failed To Running", err)
    }
}