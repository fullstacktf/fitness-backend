package service

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// AssignRoutineToUser Creates an assigned routine for the specified user, importing all exercises from base exercises
func AssignRoutineToUser(data model.AssignedRoutine) error {

	baseRoutine := GetBaseRoutine(int(data.BaseRoutineID))

	newAssignedRoutine := model.AssignedRoutine{}
	newAssignedRoutine.UserID = uint64(data.UserID)
	newAssignedRoutine.Name = baseRoutine.Name
	newAssignedRoutine.Description = baseRoutine.Description
	newAssignedRoutine.BaseRoutineID = data.BaseRoutineID

	outputCreate := storage.DB.Create(&newAssignedRoutine)

	if outputCreate.Error != nil {
		fmt.Println(outputCreate.Error)
		return outputCreate.Error
	}

	baseExercises := baseRoutine.BaseExercises

	return CreateSpecificExercisesFromBase(baseExercises, uint64(newAssignedRoutine.ID))

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

// PopulateAssignedRoutine populates assigned routine element with basic data
func PopulateAssignedRoutine() {
	exampleUser1Routine := model.AssignedRoutine{
		UserID:        1,
		BaseRoutineID: 2,
	}

	exampleUser2Routine := model.AssignedRoutine{
		UserID:        2,
		BaseRoutineID: 3,
	}

	exampleUser3Routine := model.AssignedRoutine{
		UserID:        3,
		BaseRoutineID: 2,
	}

	assignedRoutines := []model.AssignedRoutine{
		exampleUser1Routine,
		exampleUser2Routine,
		exampleUser3Routine,
	}

	for _, routine := range assignedRoutines {
		AssignRoutineToUser(routine)
	}
}
