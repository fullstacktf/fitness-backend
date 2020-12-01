package controller

import (
	"fmt"
	"strconv"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-gonic/gin"
)

var baseExerciseModel model.BaseExercise

// CreateBaseExercise Creates a BaseExercise with the given data
func CreateBaseExercise(c *gin.Context) {
	newBaseExercise := model.BaseExercise{}
	bodyErr := c.BindJSON(&newBaseExercise)

	err := service.CreateBaseExercise(newBaseExercise)

	if bodyErr == nil {
		if err != nil {
			error := service.GetGormErrorCode(err.Error())

			c.JSON(500, error)
		} else {
			c.String(200, "ok")
		}
	}
}

// GetBaseExercise Get BaseExercise by id
func GetBaseExercise(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := service.GetBaseExercise(id)

	if result.ID == 0 {
		c.JSON(404, "BaseExercise not found")
	} else {
		c.JSON(200, result)
	}
}

// GetBaseExercises Get all non deleted BaseExercises and by using a filter
func GetBaseExercises(c *gin.Context) {
	filter := model.BaseExercise{}
	reqErr := c.BindJSON(&filter)

	if reqErr == nil {
		result := service.GetBaseExercises(filter)
		c.JSON(200, result)
	}
}

// UpdateBaseExercise Update specific BaseExercise using id param in URL
func UpdateBaseExercise(c *gin.Context) {
	updatedBaseExercise := model.BaseExercise{}
	c.BindJSON(&updatedBaseExercise)

	err := service.UpdateBaseExercise(updatedBaseExercise)
	if err != nil {
		fmt.Println(err.Error())
		error := service.GetGormErrorCode(err.Error())

		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// DeleteBaseExercise Delete BaseExercise by id, logical delete
func DeleteBaseExercise(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	service.DeleteBaseExercise(id)

	c.String(200, c.Param("id"))
}
