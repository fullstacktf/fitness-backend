package service

import (
	"errors"

	"github.com/fullstacktf/fitness-backend/constants"
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreatePermission Creates a permission
func CreatePermission(data model.Permission) error {
	result := storage.DB.Create(&data)
	return result.Error
}

// GetPermission Function to return a specific permission. Not for the API controller
func GetPermission(id uint8) *model.Permission {
	permission := model.Permission{}
	storage.DB.Find(&permission, id)
	return &permission
}

//GetPermissions Gets all permissions
func GetPermissions() *[]model.Permission {
	permissions := []model.Permission{}
	storage.DB.Find(&permissions)
	return &permissions
}

//UpdatePermission Updates the description of the permission by ID
func UpdatePermission(updatedPermission model.Permission) error {
	output := storage.DB.Save(&updatedPermission)
	return output.Error
}

//DeletePermission Deletes a permission by ID
func DeletePermission(id uint8) error {
	permission := GetPermission(id)

	if permission.ID == 0 {
		return errors.New("Not Found")
	} else {
		result := storage.DB.Delete(&permission)
		return result.Error
	}
}

// PopulatePermissions Function to populate the Permission model
func PopulatePermissions() {
	createRolesPermission := model.Permission{
		Description: string(constants.CREATE_ROLES),
	}

	readRolesPermission := model.Permission{
		Description: string(constants.READ_ROLES),
	}

	updateRolesPermission := model.Permission{
		Description: string(constants.UPDATE_ROLES),
	}

	deleteRolesPermission := model.Permission{
		Description: string(constants.DELETE_ROLES),
	}

	// permissions table

	createPermissionsPermission := model.Permission{
		Description: string(constants.CREATE_PERMISSIONS),
	}

	readPermissionsPermission := model.Permission{
		Description: string(constants.READ_PERMISSIONS),
	}

	updatePermissionsPermission := model.Permission{
		Description: string(constants.UPDATE_PERMISSIONS),
	}

	deletePermissionsPermission := model.Permission{
		Description: string(constants.DELETE_PERMISSIONS),
	}

	createUsersPermission := model.Permission{
		Description: string(constants.CREATE_USERS),
	}

	readUsersPermission := model.Permission{
		Description: string(constants.READ_USERS),
	}

	updateUsersPermission := model.Permission{
		Description: string(constants.UPDATE_USERS),
	}

	deleteUsersPermission := model.Permission{
		Description: string(constants.DELETE_USERS),
	}

	updatePasswordPermission := model.Permission{
		Description: string(constants.UPDATE_PASSWORD),
	}

	createRoutinesPermission := model.Permission{
		Description: string(constants.CREATE_ROUTINES),
	}

	readRoutinesPermission := model.Permission{
		Description: string(constants.READ_ROUTINES),
	}

	updateRoutinesPermission := model.Permission{
		Description: string(constants.UPDATE_ROUTINES),
	}

	deleteRoutinesPermission := model.Permission{
		Description: string(constants.DELETE_ROUTINES),
	}

	createRoutineCategoriesPermission := model.Permission{
		Description: string(constants.CREATE_ROUTINE_CATEGORIES),
	}

	readRoutineCategoriesPermission := model.Permission{
		Description: string(constants.READ_ROUTINE_CATEGORIES),
	}

	updateRoutineCategoriesPermission := model.Permission{
		Description: string(constants.UPDATE_ROUTINE_CATEGORIES),
	}

	deleteRoutineCategoriesPermission := model.Permission{
		Description: string(constants.DELETE_ROUTINE_CATEGORIES),
	}

	createRoutinesUserPermission := model.Permission{
		Description: string(constants.CREATE_ROUTINES_USER),
	}

	readRoutinesUserPermission := model.Permission{
		Description: string(constants.READ_ROUTINES_USER),
	}

	updateRoutinesUserPermission := model.Permission{
		Description: string(constants.UPDATE_ROUTINES_USER),
	}

	deleteRoutinesUserPermission := model.Permission{
		Description: string(constants.DELETE_ROUTINES_USER),
	}

	createExercisesPermission := model.Permission{
		Description: string(constants.CREATE_EXERCISES),
	}

	readExercisesPermission := model.Permission{
		Description: string(constants.READ_EXERCISES),
	}

	updateExercisesPermission := model.Permission{
		Description: string(constants.UPDATE_EXERCISES),
	}

	deleteExercisesPermission := model.Permission{
		Description: string(constants.DELETE_EXERCISES),
	}

	createRoutinesExercisesPermission := model.Permission{
		Description: string(constants.CREATE_ROUTINES_EXERCISES),
	}

	readRoutinesExercisesPermission := model.Permission{
		Description: string(constants.READ_ROUTINES_EXERCISES),
	}

	updateRoutinesExercisesPermission := model.Permission{
		Description: string(constants.UPDATE_ROUTINES_EXERCISES),
	}

	deleteRoutinesExercisesPermission := model.Permission{
		Description: string(constants.DELETE_ROUTINES_EXERCISES),
	}

	createExerciseCategoryPermission := model.Permission{
		Description: string(constants.CREATE_EXERCISE_CATEGORY),
	}

	readExerciseCategoryPermission := model.Permission{
		Description: string(constants.READ_EXERCISE_CATEGORY),
	}

	updateExerciseCategoryPermission := model.Permission{
		Description: string(constants.UPDATE_EXERCISE_CATEGORY),
	}

	deleteExerciseCategoryPermission := model.Permission{
		Description: string(constants.DELETE_EXERCISE_CATEGORY),
	}

	createMusclesPermission := model.Permission{
		Description: string(constants.CREATE_MUSCLES),
	}

	readMusclesPermission := model.Permission{
		Description: string(constants.READ_MUSCLES),
	}

	updateMusclesPermission := model.Permission{
		Description: string(constants.UPDATE_MUSCLES),
	}

	deleteMusclesPermission := model.Permission{
		Description: string(constants.DELETE_MUSCLES),
	}

	createBaseRoutinesPermission := model.Permission{
		Description: string(constants.CREATE_BASE_ROUTINES),
	}

	readBaseRoutinesPermission := model.Permission{
		Description: string(constants.READ_BASE_ROUTINES),
	}

	updateBaseRoutinesPermission := model.Permission{
		Description: string(constants.UPDATE_BASE_ROUTINES),
	}

	deleteBaseRoutinesPermission := model.Permission{
		Description: string(constants.DELETE_BASE_ROUTINES),
	}

	updateUserStatsPermission := model.Permission{
		Description: string(constants.UPDATE_USER_STATS),
	}

	createBaseExercisePermission := model.Permission{
		Description: string(constants.CREATE_BASE_EXERCISES),
	}

	readBaseExercisePermission := model.Permission{
		Description: string(constants.READ_BASE_EXERCISES),
	}

	updateBaseExercisePermission := model.Permission{
		Description: string(constants.UPDATE_BASE_EXERCISES),
	}

	deleteBaseExercisePermission := model.Permission{
		Description: string(constants.DELETE_BASE_EXERCISES),
	}

	permissions := []model.Permission{
		createRolesPermission,
		readRolesPermission,
		updateRolesPermission,
		deleteRolesPermission,
		createPermissionsPermission,
		readPermissionsPermission,
		updatePermissionsPermission,
		deletePermissionsPermission,
		createUsersPermission,
		readUsersPermission,
		updateUsersPermission,
		deleteUsersPermission,
		updatePasswordPermission,
		createRoutinesPermission,
		readRoutinesPermission,
		updateRoutinesPermission,
		deleteRoutinesPermission,
		createRoutineCategoriesPermission,
		readRoutineCategoriesPermission,
		updateRoutineCategoriesPermission,
		deleteRoutineCategoriesPermission,
		createRoutinesUserPermission,
		readRoutinesUserPermission,
		updateRoutinesUserPermission,
		deleteRoutinesUserPermission,
		createExercisesPermission,
		readExercisesPermission,
		updateExercisesPermission,
		deleteExercisesPermission,
		createRoutinesExercisesPermission,
		readRoutinesExercisesPermission,
		updateRoutinesExercisesPermission,
		deleteRoutinesExercisesPermission,
		createExerciseCategoryPermission,
		readExerciseCategoryPermission,
		updateExerciseCategoryPermission,
		deleteExerciseCategoryPermission,
		createMusclesPermission,
		readMusclesPermission,
		updateMusclesPermission,
		deleteMusclesPermission,
		createBaseRoutinesPermission,
		readBaseRoutinesPermission,
		updateBaseRoutinesPermission,
		deleteBaseRoutinesPermission,
		updateUserStatsPermission,
		createBaseExercisePermission,
		readBaseExercisePermission,
		updateBaseExercisePermission,
		deleteBaseExercisePermission,
	}

	storage.DB.Create(&permissions)
}
