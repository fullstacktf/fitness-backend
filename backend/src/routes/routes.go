package routes

import (
	"github.com/fullstacktf/fitness-backend/controller"

	"github.com/gin-gonic/gin"
)

// SetupRouter function
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/", controller.GetWelcome)

		// Users

		v1.POST("/users/", controller.CreateUser)

		v1.GET("/user/:id", controller.GetUser)

		v1.GET("/users/", controller.GetUsers)

		v1.PUT("/users/:id", controller.UpdateUser)

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

		v1.PUT("/assignedRoutine/:id", controller.UpdateAssignedRoutine)

		v1.DELETE("/assignedRoutine/:id", controller.DeleteAssignedRoutine)

		// History. Update available to the trainer only

		v1.POST("/history/", controller.CreateHistory)

		v1.PUT("/history/:id", controller.UpdateHistory)

		// Routine specific exercises. Used to store specific exercises defined within a custom routine.

		v1.POST("/routineSpecificExercises/", controller.CreateRoutineSpecificExercise)

		v1.GET("/routineSpecificExercises/:id", controller.GetRoutineSpecificExercise)

		v1.PUT("/routineSpecificExercises/:id", controller.UpdateRoutineSpecificExercise)

		v1.DELETE("/routineSpecificExercises/:id", controller.DeleteRoutineSpecificExercise)

	}
	return r
}
