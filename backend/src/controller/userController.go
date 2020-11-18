package controller

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/gin-gonic/gin"
)

var userModel model.User

// CreateUser Create a new user
func CreateUser(c *gin.Context) {
	prueba := userModel.CreateUser()
	fmt.Println(prueba)
}

// GetUser Get user by id
func GetUser(c *gin.Context) {
	prueba := userModel.GetUser()
	fmt.Println(prueba)
}

// GetUsers Get all users
func GetUsers(c *gin.Context) {
	prueba := userModel.GetUsers()
	fmt.Println(prueba)
}

// UpdateUser Update specific user
func UpdateUser(c *gin.Context) {
	prueba := userModel.UpdateUser()
	fmt.Println(prueba)
}

// DeleteUser Delete user by id
func DeleteUser(c *gin.Context) {
	prueba := userModel.DeleteUser()
	fmt.Println(prueba)
}
