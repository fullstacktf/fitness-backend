package controller

import (
	"fmt"
	"strconv"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-gonic/gin"
)

var routineCategoryModel model.RoutineCategory

// CreateRoutineCategory creates a category with the given data
func CreateRoutineCategory(c *gin.Context) {
	newCategory := model.RoutineCategory{}
	bodyErr := c.BindJSON(&newCategory)

	err := service.CreateRoutineCategory(newCategory)

	if bodyErr == nil {
		if err != nil {
			error := service.GetGormErrorCode(err.Error())

			c.JSON(500, error)
		} else {
			c.String(200, "ok")
		}
	}
}

// GetRoutineCategory Get category by id
func GetRoutineCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := service.GetRoutineCategory(id)

	if result.ID == 0 {
		c.JSON(404, "Routine category not found")
	} else {
		c.JSON(200, result)
	}
}

// GetRoutineCategories Get all non deleted categories and by using a filter
func GetRoutineCategories(c *gin.Context) {
	filter := model.RoutineCategory{}
	reqErr := c.BindJSON(&filter)

	if reqErr == nil {
		result := service.GetRoutineCategories(filter)
		c.JSON(200, result)
	}
}

// UpdateRoutineCategory Update specific category using id param in URL
func UpdateRoutineCategory(c *gin.Context) {

	updatedCategory := model.RoutineCategory{}
	c.BindJSON(&updatedCategory)

	err := service.UpdateRoutineCategory(updatedCategory)
	if err != nil {
		fmt.Println(err.Error())
		error := service.GetGormErrorCode(err.Error())

		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// DeleteRoutineCategory Delete category by id, logical delete
func DeleteRoutineCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	service.DeleteRoutineCategory(id)

	c.String(200, c.Param("id"))

}
