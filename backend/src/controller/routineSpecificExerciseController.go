package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/fullstacktf/fitness-backend/model"
)

var routineSpecificExerciseModel model.RoutineSpecificExercise

// CreateRoutineSpecificExercise
func CreateRoutineSpecificExercise(c *gin.Context) {
	prueba := routineSpecificExerciseModel.CreateRoutineSpecificExercise()
	fmt.Println(prueba)
}

// GetRoutineSpecificExercise
func GetRoutineSpecificExercise(c *gin.Context) {
	prueba := routineSpecificExerciseModel.GetRoutineSpecificExercise()
	fmt.Println(prueba)
}

// UpdateRoutineSpecificExercise
func UpdateRoutineSpecificExercise(c *gin.Context) {
	prueba := routineSpecificExerciseModel.UpdateRoutineSpecificExercise()
	fmt.Println(prueba)
}

// DeleteRoutineSpecificExercise
func DeleteRoutineSpecificExercise(c *gin.Context) {
	prueba := routineSpecificExerciseModel.DeleteRoutineSpecificExercise()
	fmt.Println(prueba)
}
