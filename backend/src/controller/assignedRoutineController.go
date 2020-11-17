package controller

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/gin-gonic/gin"
)

var assignedRoutineModel model.AssignedRoutine

// CreateAssignedRoutine Create a new user
func CreateAssignedRoutine(c *gin.Context) {
	prueba := assignedRoutineModel.CreateAssignedRoutine()
	fmt.Println(prueba)
}

// GetAssignedRoutine Get user by id
func GetAssignedRoutine(c *gin.Context) {
	prueba := assignedRoutineModel.GetAssignedRoutine()
	fmt.Println(prueba)
}

// GetAssignedRoutines Get all users
func GetAssignedRoutines(c *gin.Context) {
	prueba := assignedRoutineModel.GetAssignedRoutines()
	fmt.Println(prueba)
}

// UpdateAssignedRoutine Update specific user
func UpdateAssignedRoutine(c *gin.Context) {
	prueba := assignedRoutineModel.UpdateAssignedRoutine()
	fmt.Println(prueba)
}

// DeleteAssignedRoutine Delete user by id
func DeleteAssignedRoutine(c *gin.Context) {
	prueba := assignedRoutineModel.DeleteAssignedRoutine()
	fmt.Println(prueba)
}
