package routes

import (
	"github.com/fullstacktf/fitness-backend/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter Function to initialize and configure gin-gonic
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/", controller.GetWelcome)

		// Users

		v1.POST("/users/", controller.CreateUser)

		v1.GET("/user/:id", controller.GetUser)

		v1.GET("/users/", controller.GetUsers)

		v1.PUT("/users/", controller.UpdateUser)

		v1.DELETE("/users/:id", controller.DeleteUser)

		// Paswords. Only accesible to admins

		v1.PUT("/password/:userId", controller.UpdateUserPass)

		// User stats

		v1.PUT("/userStat/:userId", controller.UpdateUserStat)

		// User weight history

		v1.POST("/userWeight/:userId", controller.CreateUserWeightHistory)

		v1.PUT("/userWeight/:userId", controller.UpdateUserWeightHistory)

		// Assigned routines. Used to asign custom routines to users by a trainer

		v1.POST("/assignedRoutine/", controller.CreateAssignedRoutine)

		v1.GET("/assignedRoutine/:id", controller.GetAssignedRoutine)

		v1.GET("/assignedRoutines/", controller.GetAssignedRoutines)

		v1.PUT("/assignedRoutine/", controller.UpdateAssignedRoutine)

		v1.DELETE("/assignedRoutine/:id", controller.DeleteAssignedRoutine)

		// History. Update available to the trainer only

		v1.POST("/history/", controller.CreateHistory)

		v1.PUT("/history/:id", controller.UpdateHistory)

		// Routine specific exercises. Used to store specific exercises defined within a custom routine.

		v1.POST("/routineSpecificExercises/", controller.CreateRoutineSpecificExercise)

		v1.GET("/routineSpecificExercises/:id", controller.GetRoutineSpecificExercise)

		v1.PUT("/routineSpecificExercises/:id", controller.UpdateRoutineSpecificExercise)

		v1.DELETE("/routineSpecificExercises/:id", controller.DeleteRoutineSpecificExercise)

		//Permissions

		v1.POST("/permission/", controller.CreatePermission)

		v1.GET("/permission/:id", controller.GetPermission)

		v1.GET("/permissions/", controller.GetPermissions)

		v1.PUT("/permission/", controller.UpdatePermission)

		v1.DELETE("/permission/:id", controller.DeletePermission)

		//Roles

		v1.POST("/role/", controller.CreateRole)

		v1.GET("/role/:id", controller.GetRole)

		v1.GET("/roles/", controller.GetRoles)

		v1.PUT("role/", controller.UpdateRole)

		v1.DELETE("/role/:id", controller.DeleteRole)

		//BaseExercise

		v1.POST("/baseExercise/", controller.CreateBaseExercise)

		v1.GET("/baseExercise/:id", controller.GetBaseExercise)

		v1.GET("/baseExercises/", controller.GetBaseExercises)

		v1.PUT("/baseExercise/", controller.UpdateBaseExercise)

		v1.DELETE("/baseExercise/:id", controller.DeleteBaseExercise)

		//BaseRoutine

		v1.POST("/baseRoutine/", controller.CreateBaseRoutine)

		v1.GET("/baseRoutine/:id", controller.GetBaseRoutine)

		v1.GET("/baseRoutines/", controller.GetBaseRoutines)

		v1.PUT("baseRoutine/", controller.UpdateBaseRoutine)

		v1.DELETE("/baseRoutine/:id", controller.DeleteBaseRoutine)

		//RoutineCategory

		v1.POST("/routineCategory/", controller.CreateRoutineCategory)

		v1.GET("/routineCategory/:id", controller.GetRoutineCategory)

		v1.GET("/routineCategories/", controller.GetRoutineCategories)

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

		v1.GET("/muscles/", controller.GetMuscles)

		v1.PUT("muscle/", controller.UpdateMuscle)

		v1.DELETE("/muscle/:id", controller.DeleteMuscle)
	}

	return r
}
