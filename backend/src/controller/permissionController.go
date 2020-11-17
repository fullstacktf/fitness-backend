package controller

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/gin-gonic/gin"
)

var permissionModel model.Permission

// CreatePermission Creates a permission
func CreatePermission(c *gin.Context) {
	prueba := permissionModel.CreatePermission()
	fmt.Println(prueba)
}

// GetPermission Gets a permission
func GetPermission(c *gin.Context) {
	prueba := permissionModel.GetPermission()
	fmt.Println(prueba)
}

// GetPermissions Gets all permissions
func GetPermissions(c *gin.Context) {
	prueba := permissionModel.GetPermissions()
	fmt.Println(prueba)
}

// UpdatePermission Updates a permission
func UpdatePermission(c *gin.Context) {
	prueba := permissionModel.UpdatePermission()
	fmt.Println(prueba)
}

// DeletePermission Deletes a permission
func DeletePermission(c *gin.Context) {
	prueba := permissionModel.DeletePermission()
	fmt.Println(prueba)
}
