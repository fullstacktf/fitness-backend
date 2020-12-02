package controller

import (
	"fmt"
	"strconv"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-gonic/gin"
)

// CreateBaseRoutine Creates a base routine with the given data
func CreateBaseRoutine(c *gin.Context) {
	newBaseRoutine := model.BaseRoutine{}
	bodyErr := c.BindJSON(&newBaseRoutine)

	err := service.CreateBaseRoutine(newBaseRoutine)

	if bodyErr == nil {
		if err != nil {
			error := service.GetGormErrorCode(err.Error())

			c.JSON(500, error)
		} else {
			c.String(200, "ok")
		}
	}
}

// GetBaseRoutine Gets a base routine by id
func GetBaseRoutine(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := service.GetBaseRoutine(id)

	if result.ID == 0 {
		c.JSON(404, "BaseRoutine not found")
	} else {
		c.JSON(200, result)
	}
}

// GetBaseRoutines Get all non deleted base routines and by using a filter
func GetBaseRoutines(c *gin.Context) {
	filter := model.BaseRoutine{}
	reqErr := c.BindJSON(&filter)

	if reqErr == nil {
		result := service.GetBaseRoutines(filter)
		c.JSON(200, result)
	}
}

// UpdateBaseRoutine Update specific base routine using id param in URL
func UpdateBaseRoutine(c *gin.Context) {
	updatedBaseRoutine := model.BaseRoutine{}
	c.BindJSON(&updatedBaseRoutine)

	err := service.UpdateBaseRoutine(updatedBaseRoutine)
	if err != nil {
		fmt.Println(err.Error())
		error := service.GetGormErrorCode(err.Error())

		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// DeleteBaseRoutine Delete base routine by id, logical delete
func DeleteBaseRoutine(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	service.DeleteBaseRoutine(id)

	c.String(200, c.Param("id"))
}
