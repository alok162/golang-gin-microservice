package server

import (
	"golang-gin-microservice/pkg/users"

	"github.com/gin-gonic/gin"
)

// Main routing
func Server() {
	router := gin.Default()

	// users.SubRoutes(router)

	v1 := router.Group("/api/v1/user")
	{
		users.SubRoutes(v1)
		// v1.GET("/user-details", users.UserDetails)
		// v1.POST("/register", users.RegisterUser)
	}

	// v1 := router.Group("/v1")

	// router.GET("/user-details", users.UserDetails)
	// router.POST("/register", users.RegisterUser)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
