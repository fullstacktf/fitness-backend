package controller

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/gin-gonic/gin"
)

var assignedRoutineModel model.AssignedRoutine

// CreateAssignedRoutine Create a new assigned routine for a user
func CreateAssignedRoutine(c *gin.Context) {
	prueba := assignedRoutineModel.CreateAssignedRoutine()
	fmt.Println(prueba)
}

// GetAssignedRoutine Gets an assigned routine of a user
func GetAssignedRoutine(c *gin.Context) {
	prueba := assignedRoutineModel.GetAssignedRoutine()
	fmt.Println(prueba)
}

// GetAssignedRoutines Get all assigned routines of a user
func GetAssignedRoutines(c *gin.Context) {
	prueba := assignedRoutineModel.GetAssignedRoutines()
	fmt.Println(prueba)
}

// UpdateAssignedRoutine Update specific assigned routine
func UpdateAssignedRoutine(c *gin.Context) {
	prueba := assignedRoutineModel.UpdateAssignedRoutine()
	fmt.Println(prueba)
}

// DeleteAssignedRoutine Delete assigned routine by id
func DeleteAssignedRoutine(c *gin.Context) {
	prueba := assignedRoutineModel.DeleteAssignedRoutine()
	fmt.Println(prueba)
}
