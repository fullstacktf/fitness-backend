package controller

import (
	"fmt"
	"strconv"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-gonic/gin"
)

var userModel model.User

// CreateUser Create a new user with specified data in body
func CreateUser(c *gin.Context) {
	newUser := model.User{}
	bodyErr := c.BindJSON(&newUser)

	err := service.CreateUser(newUser)

	if bodyErr == nil {
		if err != nil {
			error := service.GetGormErrorCode(err.Error())

			c.JSON(500, error)
		} else {
			c.String(200, "ok")
		}
	}
}

// GetUser Get user by id
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := service.GetUser(id)

	if result.ID == 0 {
		c.JSON(404, "User not found")
	} else {
		c.JSON(200, result)
	}
}

// GetUsers Get all non deleted users and by using a filter
func GetUsers(c *gin.Context) {
	filter := model.User{}
	reqErr := c.BindJSON(&filter)

	if reqErr == nil {
		result := service.GetUsers(filter)
		c.JSON(200, result)
	}
}

// UpdateUser Update specific user using id param in URL
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	updatedUser := model.User{}
	c.BindJSON(&updatedUser)

	if uint(id) != updatedUser.ID {
		c.String(400, "Bad request")
		return
	}

	err := service.UpdateUser(updatedUser)
	if err != nil {
		fmt.Println(err.Error())
		error := service.GetGormErrorCode(err.Error())

		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// DeleteUser Delete user by id, logical delete
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	service.DeleteUser(id)

	c.String(200, c.Param("id"))

}

// RegisterUser Creates a new user
func RegisterUser(c *gin.Context) {
	newUser := model.User{}
	bodyErr := c.BindJSON(&newUser)

	newUser.UserRole = 3

	err := service.CreateUser(newUser)

	if bodyErr == nil {
		if err != nil {
			error := service.GetGormErrorCode(err.Error())

			c.JSON(500, error)
		} else {
			c.String(200, "ok")
		}
	}
}
