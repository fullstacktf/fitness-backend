package controller

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/gin-gonic/gin"
)

var historyModel model.History

// CreateHistory Create a new history point
func CreateHistory(c *gin.Context) {
	prueba := historyModel.CreateHistory()
	fmt.Println(prueba)
}

// UpdateHistory Updates an existing history point
func UpdateHistory(c *gin.Context) {
	prueba := historyModel.CreateHistory()
	fmt.Println(prueba)
}
