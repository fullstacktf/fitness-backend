package controller

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/gin-gonic/gin"
)

var baseExerciseModel model.BaseExercise

// CreateBaseExercise Create base exercise
func CreateBaseExercise(c *gin.Context) {
	prueba := baseExerciseModel.CreateBaseExercise()
	fmt.Println(prueba)
}

// GetBaseExercise Gets a base exercise
func GetBaseExercise(c *gin.Context) {
	prueba := baseExerciseModel.GetBaseExercise()
	fmt.Println(prueba)
}

// GetBaseExercises Gets all base exercise
func GetBaseExercises(c *gin.Context) {
	prueba := baseExerciseModel.GetBaseExercises()
	fmt.Println(prueba)
}

// UpdateBaseExercise Updates a base exercise
func UpdateBaseExercise(c *gin.Context) {
	prueba := baseExerciseModel.UpdateBaseExercise()
	fmt.Println(prueba)
}

// DeleteBaseExercise Deletes a base exercise
func DeleteBaseExercise(c *gin.Context) {
	prueba := baseExerciseModel.DeleteBaseExercise()
	fmt.Println(prueba)
}
