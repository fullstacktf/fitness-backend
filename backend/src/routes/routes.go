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
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/login", login)
	r.GET("/logout", logout)

	session := r.Group("/session")
	{
		session.GET("/", getSessionID).Use(AuthRequired)
	}

	r.POST("/register", controller.RegisterUser)
	v1 := r.Group("/v1")
	{

		v1.GET("/", controller.GetWelcome)

		user := v1.Group("/user").Use(AuthRequired)
		// Users
		user.GET("/:id", controller.GetUser).Use(checkPermission(constants.READ_USERS))

		user.GET("/", controller.GetUsers).Use(checkPermission(constants.READ_USERS))

		user.POST("/", controller.CreateUser).Use(checkPermission(constants.CREATE_USERS))

		user.PUT("/", controller.UpdateUser).Use(checkPermission(constants.CREATE_USERS))

		user.DELETE("/:id", controller.DeleteUser).Use(checkPermission(constants.DELETE_USERS))

		// Paswords. Only accesible to admins
		password := v1.Group("/resetPassword").Use(AuthRequired)
		password.PUT("/:userId", controller.ResetUserPass).Use(checkRole("Admin"))

		// User stats
		userStat := v1.Group("/userStat").Use(AuthRequired)
		userStat.GET("/:userId", controller.GetUserStat).Use(checkOwnActionOrPerm("userId", constants.UPDATE_USER_STATS))

		userStat.PUT("/:userId", controller.UpdateUserStats)

		// User weight history
		userWeight := v1.Group("/userWeight").Use(AuthRequired)
		userWeight.POST("/:userId", controller.CreateUserWeightHistoryPoint).Use(checkOwnActionOrPerm("userId", constants.UPDATE_USER_STATS))

		userWeight.GET("/:userId", controller.GetWeightHistoryPoints)

		// Assigned routines. Used to asign custom routines to users by a trainer
		assignedRoutine := v1.Group("/assignedRoutine").Use(AuthRequired)
		assignedRoutine.GET("/:id/byUser/:userId", controller.GetAssignedRoutineByUserID).Use(checkOwnActionOrPerm("userId", constants.READ_ROUTINES))

		assignedRoutine.GET("/:id", controller.GetAssignedRoutine).Use(checkPermission(constants.READ_ROUTINES))

		assignedRoutine.GET("/", controller.GetAssignedRoutines).Use(checkPermission(constants.READ_ROUTINES))

		assignedRoutine.POST("/", controller.AssignRoutineToUser).Use(checkPermission(constants.CREATE_ROUTINES))

		assignedRoutine.PUT("/", controller.UpdateAssignedRoutine).Use(checkPermission(constants.UPDATE_ROUTINES))

		assignedRoutine.DELETE("/:id", controller.DeleteAssignedRoutine).Use(checkPermission(constants.DELETE_ROUTINES))

		// History. Update available to the trainer only
		history := v1.Group("/history").Use(AuthRequired)
		history.POST("/:userId", controller.CreateHistory).Use(checkOwnActionOrPerm("userId", constants.UPDATE_USER_STATS))

		history.GET("/:userId/:specExerciseId", controller.GetHistoryPoints)

		// Routine specific exercises. Used to store specific exercises defined within a custom routine.
		routineSpecificExercises := v1.Group("/routineSpecificExercises").Use(AuthRequired)
		routineSpecificExercises.GET("/:id", controller.GetSpecificExercise)

		routineSpecificExercises.GET("/", controller.GetSpecificExercises)

		routineSpecificExercises.POST("/", controller.CreateSpecificExercise)

		routineSpecificExercises.PUT("/", controller.UpdateSpecificExercise)

		routineSpecificExercises.DELETE("/:id", controller.DeleteSpecificExercise)

		//Permissions
		permission := v1.Group("/permission").Use(AuthRequired)
		permission.GET("/:id", controller.GetPermission).Use(checkPermission(constants.READ_PERMISSIONS))

		permission.GET("/", controller.GetPermissions).Use(checkPermission(constants.READ_PERMISSIONS))

		permission.POST("/", controller.CreatePermission).Use(checkPermission(constants.CREATE_PERMISSIONS))

		permission.PUT("/", controller.UpdatePermission).Use(checkPermission(constants.UPDATE_PERMISSIONS))

		permission.DELETE("/:id", controller.DeletePermission).Use(checkPermission(constants.DELETE_PERMISSIONS))

		//Roles
		role := v1.Group("/role").Use(AuthRequired)
		role.GET("/:id", controller.GetRole).Use(checkPermission(constants.READ_ROLES))

		role.GET("/", controller.GetRoles).Use(checkPermission(constants.READ_ROLES))

		role.POST("/", controller.CreateRole).Use(checkPermission(constants.CREATE_ROLES))

		role.PUT("/", controller.UpdateRole).Use(checkPermission(constants.UPDATE_ROLES))

		role.DELETE("/:id", controller.DeleteRole).Use(checkPermission(constants.DELETE_ROLES))

		//BaseExercise
		baseExercise := v1.Group("/baseExercise").Use(AuthRequired)
		baseExercise.GET("/:id", controller.GetBaseExercise).Use(checkPermission(constants.READ_BASE_EXERCISES))

		baseExercise.GET("/", controller.GetBaseExercises).Use(checkPermission(constants.READ_BASE_EXERCISES))

		baseExercise.POST("/", controller.CreateBaseExercise).Use(checkPermission(constants.CREATE_BASE_EXERCISES))

		baseExercise.PUT("/", controller.UpdateBaseExercise).Use(checkPermission(constants.UPDATE_BASE_EXERCISES))

		baseExercise.DELETE("/:id", controller.DeleteBaseExercise).Use(checkPermission(constants.DELETE_BASE_EXERCISES))

		//BaseRoutine
		baseRoutine := v1.Group("/baseRoutine").Use(AuthRequired)
		baseRoutine.GET("/:id", controller.GetBaseRoutine).Use(checkPermission(constants.READ_BASE_ROUTINES))

		baseRoutine.GET("/", controller.GetBaseRoutines).Use(checkPermission(constants.READ_BASE_ROUTINES))

		baseRoutine.POST("/", controller.CreateBaseRoutine).Use(checkPermission(constants.CREATE_BASE_ROUTINES))

		baseRoutine.PUT("/", controller.UpdateBaseRoutine).Use(checkPermission(constants.UPDATE_BASE_ROUTINES))

		baseRoutine.DELETE("/:id", controller.DeleteBaseRoutine).Use(checkPermission(constants.DELETE_BASE_ROUTINES))

		//RoutineCategory
		routineCategory := v1.Group("/routineCategory").Use(AuthRequired)
		routineCategory.GET("/:id", controller.GetRoutineCategory).Use(checkPermission(constants.READ_ROUTINE_CATEGORIES))

		routineCategory.GET("/", controller.GetRoutineCategories).Use(checkPermission(constants.READ_ROUTINE_CATEGORIES))

		routineCategory.POST("/", controller.CreateRoutineCategory).Use(checkPermission(constants.CREATE_ROUTINE_CATEGORIES))

		routineCategory.PUT("/", controller.UpdateRoutineCategory).Use(checkPermission(constants.UPDATE_ROUTINE_CATEGORIES))

		routineCategory.DELETE("/:id", controller.DeleteRoutineCategory).Use(checkPermission(constants.DELETE_ROUTINE_CATEGORIES))

		//ExerciseCategory
		exerciseCategory := v1.Group("/exerciseCategory").Use(AuthRequired)
		exerciseCategory.GET("/:id", controller.GetExerciseCategory).Use(checkPermission(constants.READ_EXERCISE_CATEGORY))

		exerciseCategory.GET("/", controller.GetExerciseCategories).Use(checkPermission(constants.READ_EXERCISE_CATEGORY))

		exerciseCategory.POST("/", controller.CreateExerciseCategory).Use(checkPermission(constants.CREATE_EXERCISE_CATEGORY))

		exerciseCategory.PUT("/", controller.UpdateExerciseCategory).Use(checkPermission(constants.UPDATE_EXERCISE_CATEGORY))

		exerciseCategory.DELETE("/:id", controller.DeleteExerciseCategory).Use(checkPermission(constants.DELETE_EXERCISE_CATEGORY))

		//Muscles
		muscle := v1.Group("/muscle").Use(AuthRequired)
		muscle.GET("/:id", controller.GetMuscle).Use(checkPermission(constants.READ_MUSCLES))

		muscle.GET("/", controller.GetMuscles).Use(checkPermission(constants.READ_MUSCLES))

		muscle.POST("/", controller.CreateMuscle).Use(checkPermission(constants.CREATE_MUSCLES))

		muscle.PUT("/", controller.UpdateMuscle).Use(checkPermission(constants.UPDATE_MUSCLES))

		muscle.DELETE("/:id", controller.DeleteMuscle).Use(checkPermission(constants.DELETE_MUSCLES))
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
