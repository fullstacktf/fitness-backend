package controller

import (
	"strconv"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-gonic/gin"
)

var userWeightHistoryModel model.UserWeightHistory

// CreateUserWeightHistoryPoint Creates a new weight history data point
func CreateUserWeightHistoryPoint(c *gin.Context) {
	newWeightPoint := model.UserWeightHistory{}
	userID, _ := strconv.Atoi(c.Param("userId"))
	bodyErr := c.BindJSON(&newWeightPoint)

	if uint64(userID) != newWeightPoint.UserID {
		c.String(400, "Bad request")
		return
	}

	err := service.CreateUserWeightHistoryPoint(newWeightPoint)

	if bodyErr == nil {
		if err != nil {
			error := service.GetGormErrorCode(err.Error())

			c.JSON(500, error)
		} else {
			c.String(200, "ok")
		}
	}
}

// GetWeightHistoryPoints Gets a paged weight history points
func GetWeightHistoryPoints(c *gin.Context) {
	pagedrequest := model.Pagedrequest{}
	userID, _ := strconv.Atoi(c.Param("userId"))

	bodyErr := c.BindJSON(&pagedrequest)

	if bodyErr == nil {
		result := service.GetWeightHistoryPoints(pagedrequest, userID)
		c.JSON(200, result)
	}
}
