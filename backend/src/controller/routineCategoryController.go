package controller

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/gin-gonic/gin"
)

var routineCategoryModel model.RoutineCategory

// CreateRoutineCategory Creates a routine category
func CreateRoutineCategory(c *gin.Context) {
	prueba := routineCategoryModel.CreateRoutineCategory()
	fmt.Println(prueba)
}

// GetRoutineCategory Gets a routine category
func GetRoutineCategory(c *gin.Context) {
	prueba := routineCategoryModel.GetRoutineCategory()
	fmt.Println(prueba)
}

// GetRoutineCategories Gets all routine categories
func GetRoutineCategories(c *gin.Context) {
	prueba := routineCategoryModel.GetRoutineCategories()
	fmt.Println(prueba)
}

// UpdateRoutineCategory Updates a routine category
func UpdateRoutineCategory(c *gin.Context) {
	prueba := routineCategoryModel.UpdateRoutineCategory()
	fmt.Println(prueba)
}

// DeleteRoutineCategory Deletes a routine category
func DeleteRoutineCategory(c *gin.Context) {
	prueba := routineCategoryModel.DeleteRoutineCategory()
	fmt.Println(prueba)
}
