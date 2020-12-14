package routes

import (
	"net/http"
	"strconv"

	"github.com/fullstacktf/fitness-backend/constants"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func checkPermission(permission constants.PermissionDef) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		session := sessions.Default(c)
		userID := session.Get(userkey).(uint)

		if service.CheckPermission(string(permission), userID) {
			return
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "unauthorized"})

	}

	return gin.HandlerFunc(fn)
}

func checkRole(role string) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		session := sessions.Default(c)
		userID := session.Get(userkey).(uint)

		if service.CheckRole(role, userID) {
			return
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "unauthorized"})

	}

	return gin.HandlerFunc(fn)
}

func checkOwnAction(paramName string) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		session := sessions.Default(c)
		userID := session.Get(userkey).(uint)

		requestedID, _ := strconv.Atoi(c.Param(paramName))

		if uint(requestedID) == userID {
			return
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "unauthorized"})

	}

	return gin.HandlerFunc(fn)
}

func checkOwnActionOrPerm(paramName string, permission constants.PermissionDef) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		session := sessions.Default(c)
		userID := session.Get(userkey).(uint)

		requestedID, _ := strconv.Atoi(c.Param(paramName))

		if service.CheckPermission(string(permission), userID) || uint(requestedID) == userID {
			return
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "unauthorized"})

	}

	return gin.HandlerFunc(fn)
}
