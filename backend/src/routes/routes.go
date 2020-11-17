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

		//Permissions

		v1.POST("/permission/", controller.CreatePermission)

		v1.GET("/permission/:id", controller.GetPermission)

		v1.GET("/permission/", controller.GetPermissions)

		v1.PUT("/permission/:id", controller.UpdatePermission)

		v1.DELETE("/permission/:id", controller.DeletePermission)

		//Roles

		v1.POST("/role/", controller.CreateRole)

		v1.GET("/role/:id", controller.GetRole)

		v1.GET("/role/", controller.GetRoles)

		v1.PUT("role/:id", controller.UpdateRole)

		v1.DELETE("/role/:id", controller.DeleteRole)

		//BaseExercise

		v1.POST("/baseExercise/", controller.CreateBaseExercise)

		v1.GET("/baseExercise/:id", controller.GetBaseExercise)

		v1.GET("/baseExercise/", controller.GetBaseExercises)

		v1.PUT("baseExercise/:id", controller.UpdateBaseExercise)

		v1.DELETE("/baseExercise/:id", controller.DeleteBaseExercise)

		//BaseRoutine

		v1.POST("/baseRoutine/", controller.CreateBaseRoutine)

		v1.GET("/baseRoutine/:id", controller.GetBaseRoutine)

		v1.GET("/baseRoutine/", controller.GetBaseRoutines)

		v1.PUT("baseRoutine/:id", controller.UpdateBaseRoutine)

		v1.DELETE("/baseRoutine/:id", controller.DeleteBaseRoutine)

		//RoutineCategory

		v1.POST("/routineCategory/", controller.CreateRoutineCategory)

		v1.GET("/routineCategory/:id", controller.GetRoutineCategory)

		v1.GET("/routineCategory/", controller.GetRoutineCategories)

		v1.PUT("routineCategory/:id", controller.UpdateRoutineCategory)

		v1.DELETE("/routineCategory/:id", controller.DeleteRoutineCategory)

		//Muscles

		v1.POST("/muscle/", controller.CreateMuscle)

		v1.GET("/muscle/:id", controller.GetMuscle)

		v1.GET("/muscle/", controller.GetMuscles)

		v1.PUT("muscle/:id", controller.UpdateMuscle)

		v1.DELETE("/muscle/:id", controller.DeleteMuscle)
	}

	return r
}
