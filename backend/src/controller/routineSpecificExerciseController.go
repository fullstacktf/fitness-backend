package controller

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/service"
)

// CreateSpecificExercise Creates a routine specific exercise for the user
func CreateSpecificExercise(c *gin.Context) {
	baseExercise := model.BaseExercise{}
	id, _ := strconv.Atoi(c.Param("id"))
	c.BindJSON(&assignedRoutineModel)

	err := service.CreateSpecificExerciseFromBase(baseExercise, uint64(id))

	if err != nil {
		error := service.GetGormErrorCode(err.Error())
		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// GetSpecificExercise Returns a routine specific exercise
func GetSpecificExercise(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := service.GetSpecificExercise(id)

	if result.ID == 0 {
		c.JSON(404, "Exercise not found")
	} else {
		c.JSON(200, result)
	}
}

// GetSpecificExercises Get all non deleted routine specific exercises and by using a filter
func GetSpecificExercises(c *gin.Context) {
	filter := model.RoutineSpecificExercise{}
	reqErr := c.BindJSON(&filter)

	if reqErr == nil {
		result := service.GetSpecificExercises(filter)
		c.JSON(200, result)
	}
}

// UpdateSpecificExercise Update specific category using id param in URL
func UpdateSpecificExercise(c *gin.Context) {

	updatedExercise := model.RoutineSpecificExercise{}
	c.BindJSON(&updatedExercise)

	err := service.UpdateSpecificExercise(updatedExercise)
	if err != nil {
		fmt.Println(err.Error())
		error := service.GetGormErrorCode(err.Error())

		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// DeleteSpecificExercise Delete routine specific exercise by id, logical delete
func DeleteSpecificExercise(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	service.DeleteSpecificExercise(id)

	c.String(200, c.Param("id"))

}
