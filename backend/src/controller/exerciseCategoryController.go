package controller

import (
	"fmt"
	"strconv"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-gonic/gin"
)

var exerciseCategoryModel model.ExerciseCategory

// CreateExerciseCategory creates a exercise category with given data
func CreateExerciseCategory(c *gin.Context) {
	newCategory := model.ExerciseCategory{}
	bodyErr := c.BindJSON(&newCategory)

	err := service.CreateExerciseCategory(newCategory)

	if bodyErr == nil {
		if err != nil {
			error := service.GetGormErrorCode(err.Error())

			c.JSON(500, error)
		} else {
			c.String(200, "ok")
		}
	}
}

// GetExerciseCategory Get category by id
func GetExerciseCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := service.GetExerciseCategory(id)

	if result.ID == 0 {
		c.JSON(404, "Exercise category not found")
	} else {
		c.JSON(200, result)
	}
}

// GetExerciseCategories Get all non deleted categories and by using a filter
func GetExerciseCategories(c *gin.Context) {
	filter := model.ExerciseCategory{}
	reqErr := c.BindJSON(&filter)

	if reqErr == nil {
		result := service.GetExerciseCategories(filter)
		c.JSON(200, result)
	}
}

// UpdateExerciseCategory Update specific category using id param in URL
func UpdateExerciseCategory(c *gin.Context) {

	updatedCategory := model.ExerciseCategory{}
	c.BindJSON(&updatedCategory)

	err := service.UpdateExerciseCategory(updatedCategory)
	if err != nil {
		fmt.Println(err.Error())
		error := service.GetGormErrorCode(err.Error())

		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// DeleteExerciseCategory Delete category by id, logical delete
func DeleteExerciseCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	service.DeleteExerciseCategory(id)

	c.String(200, c.Param("id"))

}
