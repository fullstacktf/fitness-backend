package model

import (
	"log"

	"github.com/fullstacktf/fitness-backend/storage"
	"gorm.io/gorm"
)

// Permission model
type Permission struct {
	gorm.Model
	Description string `gorm:"column:description;type:varchar(30);unique"`
}

// TableName Function to change the name of a table.
func (p *Permission) TableName() string {
	return "permissions"
}

// CreatePermission Creates a permission
func (p *Permission) CreatePermission() string {
	return "Create a Permission"
}

// GetPermission Function to return a specific permission. Not for the API controller
func GetPermission(id int) *Permission {
	permission := Permission{}
	result := storage.DB.Find(&permission, id)

	if result.Error != nil {
		log.Fatal(result.Error)
	}

	return &permission
}

// GetPermission Gets a permission by ID
func (p *Permission) GetPermission() string {
	return "GetPermission"
}

//GetPermissions Gets all permissions
func (p *Permission) GetPermissions() string {
	return "Get all permissions"
}

//UpdatePermission Updates the description of the permission by ID
func (p *Permission) UpdatePermission() string {
	return "Update a permission"
}

//DeletePermission Deletes a permission by ID
func (p *Permission) DeletePermission() string {
	return "Delete a permission"
}

// PopulatePermissions Function to populate the Permission model
func PopulatePermissions() {
	createRolesPermission := Permission{
		Description: "CREATE_ROLES",
	}

	readRolesPermission := Permission{
		Description: "READ_ROLES",
	}

	updateRolesPermission := Permission{
		Description: "UPDATE_ROLES",
	}

	deleteRolesPermission := Permission{
		Description: "DELETE_ROLES",
	}

	// permissions table

	createPermissionsPermission := Permission{
		Description: "CREATE_PERMISSIONS",
	}

	readPermissionsPermission := Permission{
		Description: "READ_PERMISSIONS",
	}

	updatePermissionsPermission := Permission{
		Description: "UPDATE_PERMISSIONS",
	}

	deletePermissionsPermission := Permission{
		Description: "DELETE_PERMISSIONS",
	}

	createUsersPermission := Permission{
		Description: "CREATE_USERS",
	}

	readUsersPermission := Permission{
		Description: "READ_USERS",
	}

	updateUsersPermission := Permission{
		Description: "UPDATE_USERS",
	}

	deleteUsersPermission := Permission{
		Description: "DELETE_USERS",
	}

	updatePasswordPermission := Permission{
		Description: "UPDATE_PASSWORD",
	}

	createRoutinesPermission := Permission{
		Description: "CREATE_ROUTINES",
	}

	readRoutinesPermission := Permission{
		Description: "READ_ROUTINES",
	}

	updateRoutinesPermission := Permission{
		Description: "UPDATE_ROUTINES",
	}

	deleteRoutinesPermission := Permission{
		Description: "DELETE_ROUTINES",
	}

	createRoutineCategoriesPermission := Permission{
		Description: "CREATE_ROUTINE_CATEGORIES",
	}

	readRoutineCategoriesPermission := Permission{
		Description: "READ_ROUTINE_CATEGORIES",
	}

	updateRoutineCategoriesPermission := Permission{
		Description: "UPDATE_ROUTINE_CATEGORIES",
	}

	deleteRoutineCategoriesPermission := Permission{
		Description: "DELETE_ROUTINE_CATEGORIES",
	}

	createRoutinesUserPermission := Permission{
		Description: "CREATE_ROUTINES_USER",
	}

	readRoutinesUserPermission := Permission{
		Description: "READ_ROUTINES_USER",
	}

	updateRoutinesUserPermission := Permission{
		Description: "UPDATE_ROUTINES_USER",
	}

	deleteRoutinesUserPermission := Permission{
		Description: "DELETE_ROUTINES_USER",
	}

	createExercisesPermission := Permission{
		Description: "CREATE_EXERCISES",
	}

	readExercisesPermission := Permission{
		Description: "READ_EXERCISES",
	}

	updateExercisesPermission := Permission{
		Description: "UPDATE_EXERCISES",
	}

	deleteExercisesPermission := Permission{
		Description: "DELETE_EXERCISES",
	}

	createRoutinesExercisesPermission := Permission{
		Description: "CREATE_ROUTINES_EXERCISES",
	}

	readRoutinesExercisesPermission := Permission{
		Description: "READ_ROUTINES_EXERCISES",
	}

	updateRoutinesExercisesPermission := Permission{
		Description: "UPDATE_ROUTINES_EXERCISES",
	}

	deleteRoutinesExercisesPermission := Permission{
		Description: "DELETE_ROUTINES_EXERCISES",
	}

	createExerciseCategoryPermission := Permission{
		Description: "CREATE_EXERCISE_CATEGORY",
	}

	readExerciseCategoryPermission := Permission{
		Description: "READ_EXERCISE_CATEGORY",
	}

	updateExerciseCategoryPermission := Permission{
		Description: "UPDATE_EXERCISE_CATEGORY",
	}

	deleteExerciseCategoryPermission := Permission{
		Description: "DELETE_EXERCISE_CATEGORY",
	}

	createMusclesPermission := Permission{
		Description: "CREATE_MUSCLES",
	}

	readMusclesPermission := Permission{
		Description: "READ_MUSCLES",
	}

	updateMusclesPermission := Permission{
		Description: "UPDATE_MUSCLES",
	}

	deleteMusclesPermission := Permission{
		Description: "DELETE_MUSCLES",
	}

	createBaseRoutinesPermission := Permission{
		Description: "CREATE_BASE_ROUTINES",
	}

	readBaseRoutinesPermission := Permission{
		Description: "READ_BASE_ROUTINES",
	}

	updateBaseRoutinesPermission := Permission{
		Description: "UPDATE_BASE_ROUTINES",
	}

	deleteBaseRoutinesPermission := Permission{
		Description: "DELETE_BASE_ROUTINES",
	}

	permissions := []Permission{
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
	}

	storage.DB.Create(&permissions)
}
