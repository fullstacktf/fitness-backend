package service

import (
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateAssignedRoutine Assigns a routine with it's exercises to the user
func CreateAssignedRoutine(data model.AssignedRoutine) error {
	output := storage.DB.Create(&data)

	return output.Error
}

// GetAssignedRoutine Get GetAssignedRoutine by id
func GetAssignedRoutine(id int) *model.AssignedRoutine {
	assignedRoutine := model.AssignedRoutine{}
	storage.DB.Find(&assignedRoutine, id)

	return &assignedRoutine
}

// GetAssignedRoutines Get assigned routines with filtering
func GetAssignedRoutines(filter model.AssignedRoutine) *[]model.AssignedRoutine {
	assignedRoutines := []model.AssignedRoutine{}
	storage.DB.Where(&filter).Find(&assignedRoutines)

	return &assignedRoutines
}

// UpdateAssignedRoutine Update specific assigned routine using id param in URL
func UpdateAssignedRoutine(updatedAssignedRoutine model.AssignedRoutine) error {

	output := storage.DB.Save(&updatedAssignedRoutine)

	return output.Error
}

// DeleteAssignedRoutine Delete an assigned routine by id, logical delete
func DeleteAssignedRoutine(id int) {
	deletedAssignedRoutine := GetAssignedRoutine(id)

	storage.DB.Delete(&deletedAssignedRoutine)

}
