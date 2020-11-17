package controller

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/gin-gonic/gin"
)

var roleModel model.Role

// CreateRole Creates a role
func CreateRole(c *gin.Context) {
	prueba := roleModel.CreateRole()
	fmt.Println(prueba)
}

// GetRole Gets a role
func GetRole(c *gin.Context) {
	prueba := roleModel.GetRole()
	fmt.Println(prueba)
}

// GetRoles Gets all roles
func GetRoles(c *gin.Context) {
	prueba := roleModel.GetRoles()
	fmt.Println(prueba)
}

// UpdateRole Updates a role
func UpdateRole(c *gin.Context) {
	prueba := roleModel.UpdateRole()
	fmt.Println(prueba)
}

// DeleteRole Deletes a role
func DeleteRole(c *gin.Context) {
	prueba := roleModel.DeleteRole()
	fmt.Println(prueba)
}
