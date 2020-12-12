package controller

import (
	"strconv"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-gonic/gin"
)

var userPassModel model.UserPass

// ResetUserPass Resets the users pass
func ResetUserPass(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("userId"))

	user := service.GetUser(userID)

	if user.ID == 0 {
		c.JSON(404, "User not found")
		return
	}

	service.ResetUserPass(userID)
}
