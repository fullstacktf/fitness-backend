package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/fullstacktf/fitness-backend/model"
)

var routineSpecificExerciseModel model.RoutineSpecificExercise

// CreateRoutineSpecificExercise Creates a routine specific exercise
func CreateRoutineSpecificExercise(c *gin.Context) {
	prueba := routineSpecificExerciseModel.CreateRoutineSpecificExercise()
	fmt.Println(prueba)
}

// GetRoutineSpecificExercise Gets a routine specific exercise
func GetRoutineSpecificExercise(c *gin.Context) {
	prueba := routineSpecificExerciseModel.GetRoutineSpecificExercise()
	fmt.Println(prueba)
}

// UpdateRoutineSpecificExercise Updates a routine specific exercises
func UpdateRoutineSpecificExercise(c *gin.Context) {
	prueba := routineSpecificExerciseModel.UpdateRoutineSpecificExercise()
	fmt.Println(prueba)
}

// DeleteRoutineSpecificExercise Deletes a routine specific exercise
func DeleteRoutineSpecificExercise(c *gin.Context) {
	prueba := routineSpecificExerciseModel.DeleteRoutineSpecificExercise()
	fmt.Println(prueba)
}
