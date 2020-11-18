package controller

import (
	"fmt"
	"net/http"

	"github.com/fullstacktf/fitness-backend/model"

	"github.com/gin-gonic/gin"
)

// GetWelcome Function to return the user welcome message
func GetWelcome(c *gin.Context) {
	var welcome model.Welcome
	welcome.Get()
	c.String(http.StatusOK, fmt.Sprintf("%s at %s", welcome.Message, welcome.Date))
}
