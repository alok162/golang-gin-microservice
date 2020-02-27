package users

import (
	"golang-gin-microservice/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

// var db *gorm.DB

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
	// fmt.Println("@@@@@@@@@@@@@@@@@@@ ", c.Param("id"))
	// if err := database.DB.Where("id = ?", c.Param("id")).First(&users).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	// 	return
	// }

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
	// database.DB.Delete(&User{id: id})

	database.DB.Delete(&users, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "successfully deleted record",
	})

	// // Get model if exist
	// var users User
	// if err := database.DB.Where("id = ?", id).First(&users).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	// 	return
	// }

	// database.DB.Delete(&users)

	// c.JSON(http.StatusOK, gin.H{"data": true})
}
