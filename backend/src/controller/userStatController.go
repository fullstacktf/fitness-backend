package controller

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/gin-gonic/gin"
)

var userStatModel model.UserStat

// UpdateUserStat updates the user stats
func UpdateUserStat(c *gin.Context) {
	prueba := userStatModel.UpdateUserStat()
	fmt.Println(prueba)
}
