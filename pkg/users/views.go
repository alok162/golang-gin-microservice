package users

import (
	"golang-gin-microservice/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get details of user
func UserDetails(c *gin.Context) {
	var users []User
	id := c.Query("id")
	c.JSON(200, gin.H{
		"message": database.DB.Find(&users, id),
	})
}

func RegisterUser(c *gin.Context) {
	reqBody := User{}
	c.Bind(&reqBody)
	database.DB.Create(&reqBody)
	c.JSON(200, gin.H{
		"message": "successfully registered",
	})
}

func UpdateUser(c *gin.Context) {
	// Get model if exist
	var users User
	// Validate input
	var userInput User
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&users).Updates(userInput)

	c.JSON(http.StatusOK, gin.H{"data": userInput})
}

func DeleteUser(c *gin.Context) {
	id := c.Query("id")
	var users []User

	database.DB.Delete(&users, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted record",
	})
}
