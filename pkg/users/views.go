package users

import (
	"golang-gin-microservice/database"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// Get details of user
func UserDetails(c *gin.Context) {
	// database.DB.AutoMigrate(&User{})
	// user := User{Name: "Alok", Email: "alokchoudhary162@gmail.com"}
	// database.DB.Create(&user)
	var users []User
	c.JSON(200, gin.H{
		"message": database.DB.Find(&users),
	})
}
