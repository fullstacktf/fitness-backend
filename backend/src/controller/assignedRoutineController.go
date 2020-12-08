package controller

import (
	"fmt"
	"strconv"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-gonic/gin"
)

var assignedRoutineModel model.AssignedRoutine

// AssignRoutineToUser Creates an assigned routine for the specified user, importing all exercises from base exercises
func AssignRoutineToUser(c *gin.Context) {
	assignedRoutineModel := model.AssignedRoutine{}
	c.BindJSON(&assignedRoutineModel)

	err := service.AssignRoutineToUser(assignedRoutineModel)

	if err != nil {
		error := service.GetGormErrorCode(err.Error())
		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// GetAssignedRoutine Get AssignedRoutine by id
func GetAssignedRoutine(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	result := service.GetAssignedRoutine(id)

	if result.ID == 0 {
		c.JSON(404, "Assigned routine not found")
	} else {
		c.JSON(200, result)
	}
}

// GetAssignedRoutines Get assigned routines with filtering
func GetAssignedRoutines(c *gin.Context) {
	filter := model.AssignedRoutine{}
	reqErr := c.BindJSON(&filter)

	if reqErr == nil {
		result := service.GetAssignedRoutines(filter)
		c.JSON(200, result)
	}
}

// UpdateAssignedRoutine Update specific assigned routine
func UpdateAssignedRoutine(c *gin.Context) {
	updatedAssignedRoutine := model.AssignedRoutine{}
	c.BindJSON(&updatedAssignedRoutine)

	err := service.UpdateAssignedRoutine(updatedAssignedRoutine)
	if err != nil {
		fmt.Println(err.Error())
		error := service.GetGormErrorCode(err.Error())

		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}

// DeleteAssignedRoutine Deletes an assigned routine by id, logical delete
func DeleteAssignedRoutine(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	service.DeleteAssignedRoutine(id)

	c.String(200, c.Param("id"))
}
