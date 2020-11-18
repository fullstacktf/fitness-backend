package model

import (
	"log"

	"github.com/fullstacktf/fitness-backend/storage"
)

// Permission model
type Permission struct {
	ID          uint8  `gorm:"column:id;type:mediumint unsigned;primaryKey"`
	Description string `gorm:"column:description;type:varchar(30);unique"`
}

//TableName function
func (p *Permission) TableName() string {
	return "permissions"
}

// CreatePermission Creates a permission
func (p *Permission) CreatePermission() string {
	return "Create a Permission"
}

// GetPermission Function to return a specific permission. Not for the API controller
func GetPermission(id uint8) *Permission {
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
		ID:          1,
		Description: "CREATE_ROLES",
	}

	readRolesPermission := Permission{
		ID:          2,
		Description: "READ_ROLES",
	}

	updateRolesPermission := Permission{
		ID:          3,
		Description: "UPDATE_ROLES",
	}

	deleteRolesPermission := Permission{
		ID:          4,
		Description: "DELETE_ROLES",
	}

	// permissions table

	createPermissionsPermission := Permission{
		ID:          5,
		Description: "CREATE_PERMISSIONS",
	}

	readPermissionsPermission := Permission{
		ID:          6,
		Description: "READ_PERMISSIONS",
	}

	updatePermissionsPermission := Permission{
		ID:          7,
		Description: "UPDATE_PERMISSIONS",
	}

	deletePermissionsPermission := Permission{
		ID:          8,
		Description: "DELETE_PERMISSIONS",
	}

	createUsersPermission := Permission{
		ID:          9,
		Description: "CREATE_USERS",
	}

	readUsersPermission := Permission{
		ID:          10,
		Description: "READ_USERS",
	}

	updateUsersPermission := Permission{
		ID:          11,
		Description: "UPDATE_USERS",
	}

	deleteUsersPermission := Permission{
		ID:          12,
		Description: "DELETE_USERS",
	}

	updatePasswordPermission := Permission{
		ID:          13,
		Description: "UPDATE_PASSWORD",
	}

	createRoutinesPermission := Permission{
		ID:          14,
		Description: "CREATE_ROUTINES",
	}

	readRoutinesPermission := Permission{
		ID:          15,
		Description: "READ_ROUTINES",
	}

	updateRoutinesPermission := Permission{
		ID:          16,
		Description: "UPDATE_ROUTINES",
	}

	deleteRoutinesPermission := Permission{
		ID:          17,
		Description: "DELETE_ROUTINES",
	}

	createRoutineCategoriesPermission := Permission{
		ID:          18,
		Description: "CREATE_ROUTINE_CATEGORIES",
	}

	readRoutineCategoriesPermission := Permission{
		ID:          19,
		Description: "READ_ROUTINE_CATEGORIES",
	}

	updateRoutineCategoriesPermission := Permission{
		ID:          20,
		Description: "UPDATE_ROUTINE_CATEGORIES",
	}

	deleteRoutineCategoriesPermission := Permission{
		ID:          21,
		Description: "DELETE_ROUTINE_CATEGORIES",
	}

	createRoutinesUserPermission := Permission{
		ID:          22,
		Description: "CREATE_ROUTINES_USER",
	}

	readRoutinesUserPermission := Permission{
		ID:          23,
		Description: "READ_ROUTINES_USER",
	}

	updateRoutinesUserPermission := Permission{
		ID:          24,
		Description: "UPDATE_ROUTINES_USER",
	}

	deleteRoutinesUserPermission := Permission{
		ID:          25,
		Description: "DELETE_ROUTINES_USER",
	}

	createExercisesPermission := Permission{
		ID:          26,
		Description: "CREATE_EXERCISES",
	}

	readExercisesPermission := Permission{
		ID:          27,
		Description: "READ_EXERCISES",
	}

	updateExercisesPermission := Permission{
		ID:          28,
		Description: "UPDATE_EXERCISES",
	}

	deleteExercisesPermission := Permission{
		ID:          29,
		Description: "DELETE_EXERCISES",
	}

	createRoutinesExercisesPermission := Permission{
		ID:          30,
		Description: "CREATE_ROUTINES_EXERCISES",
	}

	readRoutinesExercisesPermission := Permission{
		ID:          31,
		Description: "READ_ROUTINES_EXERCISES",
	}

	updateRoutinesExercisesPermission := Permission{
		ID:          32,
		Description: "UPDATE_ROUTINES_EXERCISES",
	}

	deleteRoutinesExercisesPermission := Permission{
		ID:          33,
		Description: "DELETE_ROUTINES_EXERCISES",
	}

	createExerciseCategoryPermission := Permission{
		ID:          34,
		Description: "CREATE_EXERCISE_CATEGORY",
	}

	readExerciseCategoryPermission := Permission{
		ID:          35,
		Description: "READ_EXERCISE_CATEGORY",
	}

	updateExerciseCategoryPermission := Permission{
		ID:          36,
		Description: "UPDATE_EXERCISE_CATEGORY",
	}

	deleteExerciseCategoryPermission := Permission{
		ID:          37,
		Description: "DELETE_EXERCISE_CATEGORY",
	}

	createMusclesPermission := Permission{
		ID:          38,
		Description: "CREATE_MUSCLES",
	}

	readMusclesPermission := Permission{
		ID:          39,
		Description: "READ_MUSCLES",
	}

	updateMusclesPermission := Permission{
		ID:          40,
		Description: "UPDATE_MUSCLES",
	}

	deleteMusclesPermission := Permission{
		ID:          41,
		Description: "DELETE_MUSCLES",
	}

	createBaseRoutinesPermission := Permission{
		ID:          42,
		Description: "CREATE_BASE_ROUTINES",
	}

	readBaseRoutinesPermission := Permission{
		ID:          43,
		Description: "READ_BASE_ROUTINES",
	}

	updateBaseRoutinesPermission := Permission{
		ID:          44,
		Description: "UPDATE_BASE_ROUTINES",
	}

	deleteBaseRoutinesPermission := Permission{
		ID:          45,
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
