package controller

import (
	"fmt"
	"strconv"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-gonic/gin"
)

var userModel model.User

// CreateUser Create a new user
func CreateUser(c *gin.Context) {
	newUser := model.User{}
	c.BindJSON(&newUser)

	err := service.CreateUser(newUser)
	if err != nil {
		fmt.Println(err.Error())
		error := service.GetGormErrorCode(err.Error())

		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// GetUser Get user by id
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := service.GetUser(id)
	fmt.Println(result)

	if result.ID == 0 {
		c.JSON(404, nil)
	} else {
		c.JSON(200, result)
	}
}

// GetUser Get user by Email
func GetUserByDni(c *gin.Context) {
	dni := c.Param("dni")
	result := service.GetUserByDni(dni)

	if result.ID == 0 {
		c.JSON(404, nil)
	} else {
		c.JSON(200, result)
	}
}

// GetUser Get user by Email
func GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	result := service.GetUserByEmail(email)

	if result.ID == 0 {
		c.JSON(404, nil)
	} else {
		c.JSON(200, result)
	}
}

// GetUsers Get all users
func GetUsers(c *gin.Context) {
	result := service.GetUsers()

	if len(*result) == 0 {
		c.JSON(404, nil)
	} else {
		c.JSON(200, result)
	}
}

// UpdateUser Update specific user
func UpdateUser(c *gin.Context) {

	updatedUser := model.User{}
	c.BindJSON(&updatedUser)

	err := service.UpdateUser(updatedUser)
	if err != nil {
		fmt.Println(err.Error())
		error := service.GetGormErrorCode(err.Error())

		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// DeleteUser Delete user by id
func DeleteUser(c *gin.Context) {
	prueba := service.DeleteUser()
	fmt.Println(prueba)
}
