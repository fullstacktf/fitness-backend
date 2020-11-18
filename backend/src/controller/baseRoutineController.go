package controller

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/gin-gonic/gin"
)

var baseRoutineModel model.BaseRoutine

// CreateBaseRoutine Creates a base routine
func CreateBaseRoutine(c *gin.Context) {
	prueba := baseRoutineModel.CreateBaseRoutine()
	fmt.Println(prueba)
}

// GetBaseRoutine Gets a base routine
func GetBaseRoutine(c *gin.Context) {
	prueba := baseRoutineModel.GetBaseRoutine()
	fmt.Println(prueba)
}

// GetBaseRoutines Gets all base routine
func GetBaseRoutines(c *gin.Context) {
	prueba := baseRoutineModel.GetBaseRoutines()
	fmt.Println(prueba)
}

// UpdateBaseRoutine Updates a base routine
func UpdateBaseRoutine(c *gin.Context) {
	prueba := baseRoutineModel.UpdateBaseRoutine()
	fmt.Println(prueba)
}

// DeleteBaseRoutine Deletes a base routine
func DeleteBaseRoutine(c *gin.Context) {
	prueba := baseRoutineModel.DeleteBaseRoutine()
	fmt.Println(prueba)
}
