package controller

import (
	"strconv"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-gonic/gin"
)

var historyModel model.History

// CreateHistory Create a new history point
func CreateHistory(c *gin.Context) {
	newHistoryPoint := model.History{}
	userID, _ := strconv.Atoi(c.Param("userId"))
	bodyErr := c.BindJSON(&newHistoryPoint)

	specificExercise := service.GetSpecificExercise(int(newHistoryPoint.RoutineSpecificExercisesID))
	assignedRoutine := service.GetAssignedRoutine(int(specificExercise.AssignedRoutinesID))

	if uint64(userID) != assignedRoutine.UserID {
		c.String(400, "Bad request")
		return
	}

	err := service.CreateHistory(newHistoryPoint)

	if bodyErr == nil {
		if err != nil {
			error := service.GetGormErrorCode(err.Error())

			c.JSON(500, error)
		} else {
			c.String(200, "ok")
		}
	}
}

// GetHistoryPoints Gets a paginated list of history point of the specified user
func GetHistoryPoints(c *gin.Context) {
	pagedrequest := model.Pagedrequest{}
	specExerciseID, _ := strconv.Atoi(c.Param("specExerciseId"))

	bodyErr := c.BindJSON(&pagedrequest)

	if bodyErr == nil {
		result := service.GetHistoryPoints(pagedrequest, specExerciseID)
		c.JSON(200, result)
	}
}
