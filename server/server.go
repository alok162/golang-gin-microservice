package server

import (
	"golang-gin-microservice/pkg/users"

	"github.com/gin-gonic/gin"
)

// Main routing
func Server() {
	router := gin.Default()

	router.GET("/user", users.UserDetails)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
