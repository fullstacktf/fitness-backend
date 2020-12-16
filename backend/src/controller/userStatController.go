package controller

import (
	"fmt"
	"strconv"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-gonic/gin"
)

var userStatModel model.UserStat

// GetUserStat Gets the stats of the specified user
func GetUserStat(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userId"))
	result := service.GetUserStat(userID)

	if result.UserID == 0 {
		c.JSON(404, "User stats not found")
	} else {
		c.JSON(200, result)
	}
}

// UpdateUserStats Updates stats of the specified user
func UpdateUserStats(c *gin.Context) {
	updatedUserStats := model.UserStat{}
	c.BindJSON(&updatedUserStats)

	err := service.UpdateUserStats(updatedUserStats)
	if err != nil {
		fmt.Println(err.Error())
		error := service.GetGormErrorCode(err.Error())

		c.JSON(500, error)
	} else {
		c.String(200, "ok")
	}
}
