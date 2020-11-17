package controller

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/gin-gonic/gin"
)

var muscleModel model.Muscle

// CreateMuscle Create a muscle group
func CreateMuscle(c *gin.Context) {
	prueba := muscleModel.CreateMuscle()
	fmt.Println(prueba)
}

// GetMuscle gets a muscle group
func GetMuscle(c *gin.Context) {
	prueba := muscleModel.GetMuscle()
	fmt.Println(prueba)
}

// GetMuscles get all muscle groups
func GetMuscles(c *gin.Context) {
	prueba := muscleModel.GetMuscles()
	fmt.Println(prueba)
}

// UpdateMuscle Updates a muscle group
func UpdateMuscle(c *gin.Context) {
	prueba := muscleModel.UpdateMuscle()
	fmt.Println(prueba)
}

// DeleteMuscle Deletes a muscle group
func DeleteMuscle(c *gin.Context) {
	prueba := muscleModel.DeleteMuscle()
	fmt.Println(prueba)
}
