package controller

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/gin-gonic/gin"
)

var userPassModel model.UserPass

// UpdateUserPass updates the users pass
func UpdateUserPass(c *gin.Context) {
	prueba := userPassModel.UpdateUserPass()
	fmt.Println(prueba)
}