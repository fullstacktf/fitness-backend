package routes

import (
	"fmt"
	"net/http"

	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func checkPermission(permission string) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		session := sessions.Default(c)
		userID := session.Get(userkey).(uint)

		fmt.Println("CHECKING PERMISSIONS")

		if service.CheckPermission(permission, userID) {
			return
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "unauthorized"})

	}

	return gin.HandlerFunc(fn)
}
