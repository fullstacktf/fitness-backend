package controller

import (
	"fmt"
	"strconv"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-gonic/gin"
)

var muscleModel model.Muscle

// CreateMuscle Create a new Muscle with specified data in body
func CreateMuscle(c *gin.Context) {
	newMuscle := model.Muscle{}
	bodyErr := c.BindJSON(&newMuscle)

	err := service.CreateMuscle(newMuscle)

	if bodyErr == nil {
		if err != nil {
			error := service.GetGormErrorCode(err.Error())

			c.JSON(500, error)
		} else {
			c.String(200, "ok")
		}
	}
}

// GetMuscle Get muscle by id
func GetMuscle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := service.GetMuscle(id)

	if result.ID == 0 {
		c.JSON(404, "Muscle not found")
	} else {
		c.JSON(200, result)
	}
}

// GetMuscles Get all non deleted Muscles and by using a filter
func GetMuscles(c *gin.Context) {
	filter := model.Muscle{}
	reqErr := c.BindJSON(&filter)

	if reqErr == nil {
		result := service.GetMuscles(filter)
		c.JSON(200, result)
	}
}

// UpdateMuscle Update specific Muscle using id param in URL
func UpdateMuscle(c *gin.Context) {

	updatedMuscle := model.Muscle{}
	c.BindJSON(&updatedMuscle)

	err := service.UpdateMuscle(updatedMuscle)
	if err != nil {
		fmt.Println(err.Error())
		error := service.GetGormErrorCode(err.Error())

		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// DeleteMuscle Delete Muscle by id, logical delete
func DeleteMuscle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	service.DeleteMuscle(id)

	c.String(200, c.Param("id"))

}
