package routes

import (
	"net/http"
	"strings"
	"time"

	"github.com/fullstacktf/fitness-backend/constants"
	"github.com/fullstacktf/fitness-backend/controller"
	"github.com/fullstacktf/fitness-backend/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	userkey = "user"
)

// SetupRouter Function to initialize and configure gin-gonic
func SetupRouter() *gin.Engine {

	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{Domain: constants.DomainProduction})
	r.Use(sessions.Sessions("mysession", store))
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://youlift.xyz", "http://youlift.xyz"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/login", login)
	r.GET("/logout", logout)

	session := r.Group("/session").Use(AuthRequired)
	{
		session.GET("/", getSessionID)
	}

	r.POST("/register", controller.RegisterUser)
	v1 := r.Group("/v1")
	{

		v1.GET("/", controller.GetWelcome)

		user := v1.Group("/user").Use(AuthRequired)
		// Users
		user.Use(checkOwnActionOrPerm("id", constants.UPDATE_USERS)).PUT("/:id", controller.UpdateUser)

		user.Use(checkOwnActionOrPerm("id", constants.READ_USERS)).GET("/:id", controller.GetUser)

		user.Use(checkPermission(constants.READ_USERS)).Use(checkRole("Admin")).POST("/filter", controller.GetUsers)

		user.Use(checkPermission(constants.CREATE_USERS)).POST("/", controller.CreateUser)

		user.Use(checkPermission(constants.DELETE_USERS)).DELETE("/:id", controller.DeleteUser)

		// Paswords. Only accesible to admins
		password := v1.Group("/resetPassword").Use(AuthRequired)
		password.Use(checkRole("Admin")).PUT("/:userId", controller.ResetUserPass)

		// User stats
		userStat := v1.Group("/userStat").Use(AuthRequired)
		userStat.Use(checkOwnActionOrPerm("userId", constants.UPDATE_USER_STATS)).GET("/:userId", controller.GetUserStat)

		userStat.PUT("/:userId", controller.UpdateUserStats)

		// User weight history
		userWeight := v1.Group("/userWeight").Use(AuthRequired)
		userWeight.Use(checkOwnActionOrPerm("userId", constants.UPDATE_USER_STATS)).POST("/:userId", controller.CreateUserWeightHistoryPoint)

		userWeight.POST("/:userId/points", controller.GetWeightHistoryPoints)

		// Assigned routines. Used to asign custom routines to users by a trainer
		assignedRoutine := v1.Group("/assignedRoutine").Use(AuthRequired)
		assignedRoutine.Use(checkOwnActionOrPerm("userId", constants.READ_ROUTINES)).GET("/:id/byUser/:userId", controller.GetAssignedRoutinesByUserID)

		assignedRoutine.GET("/:id", controller.GetAssignedRoutine)

		assignedRoutine.POST("/filter", controller.GetAssignedRoutines)

		assignedRoutine.Use(checkPermission(constants.CREATE_ROUTINES)).POST("/", controller.AssignRoutineToUser)

		assignedRoutine.Use(checkPermission(constants.UPDATE_ROUTINES)).PUT("/", controller.UpdateAssignedRoutine)

		assignedRoutine.Use(checkPermission(constants.DELETE_ROUTINES)).DELETE("/:id", controller.DeleteAssignedRoutine)

		// History. Update available to the trainer only
		history := v1.Group("/history").Use(AuthRequired)
		history.Use(checkOwnActionOrPerm("userId", constants.UPDATE_USER_STATS)).POST("/:userId", controller.CreateHistory)

		history.GET("/:userId/:specExerciseId", controller.GetHistoryPoints)

		// Routine specific exercises. Used to store specific exercises defined within a custom routine.
		routineSpecificExercises := v1.Group("/routineSpecificExercise").Use(AuthRequired)

		routineSpecificExercises.POST("/filter", controller.GetSpecificExercises)

		routineSpecificExercises.Use(checkPermission(constants.READ_ROUTINES_EXERCISES)).GET("/:id", controller.GetSpecificExercise)

		routineSpecificExercises.Use(checkPermission(constants.CREATE_ROUTINES_EXERCISES)).POST("/", controller.CreateSpecificExercise)

		routineSpecificExercises.Use(checkPermission(constants.UPDATE_ROUTINES_EXERCISES)).PUT("/", controller.UpdateSpecificExercise)

		routineSpecificExercises.Use(checkPermission(constants.DELETE_ROUTINES_EXERCISES)).DELETE("/:id", controller.DeleteSpecificExercise)

		//Permissions
		permission := v1.Group("/permission").Use(AuthRequired)
		permission.Use(checkPermission(constants.READ_PERMISSIONS)).GET("/:id", controller.GetPermission)

		permission.Use(checkPermission(constants.READ_PERMISSIONS)).POST("/filter", controller.GetPermissions)

		permission.Use(checkPermission(constants.CREATE_PERMISSIONS)).POST("/", controller.CreatePermission)

		permission.Use(checkPermission(constants.UPDATE_PERMISSIONS)).PUT("/", controller.UpdatePermission)

		permission.Use(checkPermission(constants.DELETE_PERMISSIONS)).DELETE("/:id", controller.DeletePermission)

		//Roles
		role := v1.Group("/role").Use(AuthRequired)
		role.Use(checkPermission(constants.READ_ROLES)).GET("/:id", controller.GetRole)

		role.Use(checkPermission(constants.READ_ROLES)).POST("/filter", controller.GetRoles)

		role.Use(checkPermission(constants.CREATE_ROLES)).POST("/", controller.CreateRole)

		role.Use(checkPermission(constants.UPDATE_ROLES)).PUT("/", controller.UpdateRole)

		role.Use(checkPermission(constants.DELETE_ROLES)).DELETE("/:id", controller.DeleteRole)

		//BaseExercise
		baseExercise := v1.Group("/baseExercise").Use(AuthRequired)
		baseExercise.GET("/:id", controller.GetBaseExercise)

		baseExercise.POST("/filter", controller.GetBaseExercises)

		baseExercise.Use(checkPermission(constants.CREATE_BASE_EXERCISES)).POST("/", controller.CreateBaseExercise)

		baseExercise.Use(checkPermission(constants.UPDATE_BASE_EXERCISES)).PUT("/", controller.UpdateBaseExercise)

		baseExercise.Use(checkPermission(constants.DELETE_BASE_EXERCISES)).DELETE("/:id", controller.DeleteBaseExercise)

		//BaseRoutine
		baseRoutine := v1.Group("/baseRoutine").Use(AuthRequired)
		baseRoutine.GET("/:id", controller.GetBaseRoutine)

		baseRoutine.POST("/filter", controller.GetBaseRoutines)

		baseRoutine.Use(checkPermission(constants.CREATE_BASE_ROUTINES)).POST("/", controller.CreateBaseRoutine)

		baseRoutine.Use(checkPermission(constants.UPDATE_BASE_ROUTINES)).PUT("/", controller.UpdateBaseRoutine)

		baseRoutine.Use(checkPermission(constants.DELETE_BASE_ROUTINES)).DELETE("/:id", controller.DeleteBaseRoutine)

		//RoutineCategory
		routineCategory := v1.Group("/routineCategory").Use(AuthRequired)
		routineCategory.GET("/:id", controller.GetRoutineCategory)

		routineCategory.POST("/filter", controller.GetRoutineCategories)

		routineCategory.Use(checkPermission(constants.CREATE_ROUTINE_CATEGORIES)).POST("/", controller.CreateRoutineCategory)

		routineCategory.Use(checkPermission(constants.UPDATE_ROUTINE_CATEGORIES)).PUT("/", controller.UpdateRoutineCategory)

		routineCategory.Use(checkPermission(constants.DELETE_ROUTINE_CATEGORIES)).DELETE("/:id", controller.DeleteRoutineCategory)

		//ExerciseCategory
		exerciseCategory := v1.Group("/exerciseCategory").Use(AuthRequired)
		exerciseCategory.GET("/:id", controller.GetExerciseCategory)

		exerciseCategory.POST("/filter", controller.GetExerciseCategories)

		exerciseCategory.Use(checkPermission(constants.CREATE_EXERCISE_CATEGORY)).POST("/", controller.CreateExerciseCategory)

		exerciseCategory.Use(checkPermission(constants.UPDATE_EXERCISE_CATEGORY)).PUT("/", controller.UpdateExerciseCategory)

		exerciseCategory.Use(checkPermission(constants.DELETE_EXERCISE_CATEGORY)).DELETE("/:id", controller.DeleteExerciseCategory)

		//Muscles
		muscle := v1.Group("/muscle").Use(AuthRequired)
		muscle.GET("/:id", controller.GetMuscle)

		muscle.POST("/filter", controller.GetMuscles)

		muscle.Use(checkPermission(constants.CREATE_MUSCLES)).POST("/", controller.CreateMuscle)

		muscle.Use(checkPermission(constants.UPDATE_MUSCLES)).PUT("/", controller.UpdateMuscle)

		muscle.Use(checkPermission(constants.DELETE_MUSCLES)).DELETE("/:id", controller.DeleteMuscle)
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

func getSessionID(c *gin.Context) {
	session := sessions.Default(c)
	c.JSON(http.StatusOK, session.Get(userkey))
}
