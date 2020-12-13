package routes

import (
	"net/http"
	"strings"

	"github.com/fullstacktf/fitness-backend/controller"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

const (
	userkey = "user"
)

// SetupRouter Function to initialize and configure gin-gonic
func SetupRouter() *gin.Engine {

	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/login", login)
	r.GET("/logout", logout)
	v1 := r.Group("/v1")
	{

		v1.GET("/", controller.GetWelcome)

		user := v1.Group("/user").Use(AuthRequired)
		// Users
		user.POST("/", controller.CreateUser).Use(checkPermission("CREATE_USERS"))

		user.GET("/:id", controller.GetUser).Use(checkPermission("READ_USERS"))

		user.GET("/", controller.GetUsers).Use(checkPermission("READ_USERS"))

		user.PUT("/", controller.UpdateUser).Use(checkPermission("UPDATE_USERS"))

		user.DELETE("/:id", controller.DeleteUser).Use(checkPermission("DELETE_USERS"))

		// Paswords. Only accesible to admins

		v1.PUT("/resetPassword/:userId", controller.ResetUserPass)

		// User stats

		v1.GET("/userStat/:userId", controller.GetUserStat)

		v1.PUT("/userStat/:userId", controller.UpdateUserStats)

		// User weight history

		v1.POST("/userWeight/:userId", controller.CreateUserWeightHistoryPoint)

		v1.GET("/userWeight/:userId", controller.GetWeightHistoryPoints)

		// Assigned routines. Used to asign custom routines to users by a trainer

		v1.POST("/assignedRoutine/", controller.AssignRoutineToUser)

		v1.GET("/assignedRoutine/:id", controller.GetAssignedRoutine)

		v1.GET("/assignedRoutine/", controller.GetAssignedRoutines)

		v1.PUT("/assignedRoutine/", controller.UpdateAssignedRoutine)

		v1.DELETE("/assignedRoutine/:id", controller.DeleteAssignedRoutine)

		// History. Update available to the trainer only

		v1.POST("/history/", controller.CreateHistory)

		v1.GET("/history/:specExerciseId", controller.GetHistoryPoints)

		// Routine specific exercises. Used to store specific exercises defined within a custom routine.

		v1.POST("/routineSpecificExercises/", controller.CreateSpecificExercise)

		v1.GET("/routineSpecificExercises/:id", controller.GetSpecificExercise)

		v1.GET("/routineSpecificExercises/", controller.GetSpecificExercises)

		v1.PUT("/routineSpecificExercises/", controller.UpdateSpecificExercise)

		v1.DELETE("/routineSpecificExercises/:id", controller.DeleteSpecificExercise)

		//Permissions

		v1.POST("/permission/", controller.CreatePermission)

		v1.GET("/permission/:id", controller.GetPermission)

		v1.GET("/permission/", controller.GetPermissions)

		v1.PUT("/permission/", controller.UpdatePermission)

		v1.DELETE("/permission/:id", controller.DeletePermission)

		//Roles

		v1.POST("/role/", controller.CreateRole)

		v1.GET("/role/:id", controller.GetRole)

		v1.GET("/role/", controller.GetRoles)

		v1.PUT("role/", controller.UpdateRole)

		v1.DELETE("/role/:id", controller.DeleteRole)

		//BaseExercise

		v1.POST("/baseExercise/", controller.CreateBaseExercise)

		v1.GET("/baseExercise/:id", controller.GetBaseExercise)

		v1.GET("/baseExercise/", controller.GetBaseExercises)

		v1.PUT("/baseExercise/", controller.UpdateBaseExercise)

		v1.DELETE("/baseExercise/:id", controller.DeleteBaseExercise)

		//BaseRoutine

		v1.POST("/baseRoutine/", controller.CreateBaseRoutine)

		v1.GET("/baseRoutine/:id", controller.GetBaseRoutine)

		v1.GET("/baseRoutine/", controller.GetBaseRoutines)

		v1.PUT("baseRoutine/", controller.UpdateBaseRoutine)

		v1.DELETE("/baseRoutine/:id", controller.DeleteBaseRoutine)

		//RoutineCategory

		v1.POST("/routineCategory/", controller.CreateRoutineCategory)

		v1.GET("/routineCategory/:id", controller.GetRoutineCategory)

		v1.GET("/routineCategory/", controller.GetRoutineCategories)

		v1.PUT("routineCategory/", controller.UpdateRoutineCategory)

		v1.DELETE("/routineCategory/:id", controller.DeleteRoutineCategory)

		//ExerciseCategory

		v1.POST("/exerciseCategory/", controller.CreateExerciseCategory)

		v1.GET("/exerciseCategory/:id", controller.GetExerciseCategory)

		v1.GET("/exerciseCategory/", controller.GetExerciseCategories)

		v1.PUT("exerciseCategory/", controller.UpdateExerciseCategory)

		v1.DELETE("/exerciseCategory/:id", controller.DeleteExerciseCategory)

		//Muscles

		v1.POST("/muscle/", controller.CreateMuscle)

		v1.GET("/muscle/:id", controller.GetMuscle)

		v1.GET("/muscle/", controller.GetMuscles)

		v1.PUT("muscle/", controller.UpdateMuscle)

		v1.DELETE("/muscle/:id", controller.DeleteMuscle)
	}

	return r
}

// AuthRequired is a simple middleware to check the session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

// Login is a handler that parses a form and checks for specific data
func login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")

	if strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parameters can't be empty"})
		return
	}

	login, userID := service.CheckLogin(username, password)
	if !login {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	session.Set(userkey, userID)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
}

// Logout logouts
func logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)

	if user == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	session.Delete(userkey)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
