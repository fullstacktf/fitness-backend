package controller

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/gin-gonic/gin"
)

var userWeightHistoryModel model.UserWeightHistory

// CreateUserWeightHistory Creates a new weight history data point
func CreateUserWeightHistory(c *gin.Context) {
	prueba := userWeightHistoryModel.CreateUserWeightHistory()
	fmt.Println(prueba)
}

// UpdateUserWeightHistory Updates an existing weight history data point
func UpdateUserWeightHistory(c *gin.Context) {
	prueba := userWeightHistoryModel.UpdateUserWeightHistory()
	fmt.Println(prueba)
}
