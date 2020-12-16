package controller

import (
	"fmt"
	"strconv"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-gonic/gin"
)

// CreatePermission Creates a permission
func CreatePermission(c *gin.Context) {
	newPermission := model.Permission{}
	c.BindJSON(&newPermission)

	err := service.CreatePermission(newPermission)

	if err != nil {
		error := service.GetGormErrorCode(err.Error())
		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// GetPermission Gets a permission
func GetPermission(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := service.GetPermission(uint8(id))

	if result.ID == 0 {
		c.JSON(404, nil)
	} else {
		c.JSON(200, result)
	}
}

// GetPermissions Gets all permissions
func GetPermissions(c *gin.Context) {
	result := service.GetPermissions()

	if len(*result) == 0 {
		c.JSON(404, nil)
	} else {
		c.JSON(200, result)
	}
}

// UpdatePermission Updates a permission
func UpdatePermission(c *gin.Context) {
	updatePermission := model.Permission{}
	c.BindJSON(&updatePermission)

	err := service.UpdatePermission(updatePermission)

	if err != nil {
		fmt.Println(err.Error())
		error := service.GetGormErrorCode(err.Error())

		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// DeletePermission Deletes a permission
func DeletePermission(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := service.DeletePermission(uint8(id))

	if err != nil {
		c.String(404, "Cannot delete")
	} else {
		c.String(200, "ok")
	}
}
