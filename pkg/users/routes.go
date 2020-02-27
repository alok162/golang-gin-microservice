package users

import "github.com/gin-gonic/gin"

func SubRoutes(v1 *gin.RouterGroup) {
	v1.GET("/user-details", UserDetails)
	v1.POST("/register", RegisterUser)
	v1.DELETE("/delete-user", DeleteUser)
	v1.PATCH("/update-user-details", UpdateUser)
}
