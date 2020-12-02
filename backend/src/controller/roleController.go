package controller

import (
	"fmt"
	"strconv"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-gonic/gin"
)

// CreateRole Creates a role
func CreateRole(c *gin.Context) {
	newRole := model.Role{}
	c.BindJSON(&newRole)

	err := service.CreateRole(newRole)

	if err != nil {
		error := service.GetGormErrorCode(err.Error())
		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// GetRole Gets a role
func GetRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := service.GetRole(uint8(id))

	if result.ID == 0 {
		c.JSON(404, nil)
	} else {
		c.JSON(200, result)
	}
}

// GetRoles Gets all roles
func GetRoles(c *gin.Context) {
	result := service.GetRoles()

	if len(*result) == 0 {
		c.JSON(404, nil)
	} else {
		c.JSON(200, result)
	}
}

// UpdateRole Updates a role
func UpdateRole(c *gin.Context) {
	updatedRole := model.Role{}
	c.BindJSON(&updatedRole)

	err := service.UpdateRole(updatedRole)

	if err != nil {
		fmt.Println(err.Error())
		error := service.GetGormErrorCode(err.Error())

		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// DeleteRole Deletes a role
func DeleteRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := service.DeleteRole(uint8(id))

	if err != nil {
		c.String(404, "Cannot delete")
	} else {
		c.String(200, "ok")
	}
}
