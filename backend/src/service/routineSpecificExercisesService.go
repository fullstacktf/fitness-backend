package service

import (
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateSpecificExercisesFromBase Creates specific exercises for the indicated assigned routine from the indicated base exercises
func CreateSpecificExercisesFromBase(baseExercises []*model.BaseExercise, assignedRoutineID uint64) error {

	for i := 0; i < len(baseExercises); i++ {
		newExercise := model.RoutineSpecificExercise{}
		newExercise.BaseExercisesID = uint64(baseExercises[i].ID)
		newExercise.AssignedRoutinesID = assignedRoutineID
		CreateSpecificExerciseFromBase(newExercise)
	}

	return nil
}

// CreateSpecificExerciseFromBase Creates only one specific exercise for the indicated assigned routine from the indicated base exercise
func CreateSpecificExerciseFromBase(data model.RoutineSpecificExercise) error {

	baseExercise := GetBaseExercise(int(data.BaseExercisesID))

	newExercise := model.RoutineSpecificExercise{}
	newExercise.BaseExercisesID = uint64(baseExercise.ID)
	newExercise.Repetitions = baseExercise.DefaultRepetitions
	newExercise.Series = baseExercise.DefaultSeries
	newExercise.AssignedRoutinesID = data.AssignedRoutinesID
	output := storage.DB.Create(&newExercise)

	return output.Error
}

// GetSpecificExercise Returns a routine specific exercise
func GetSpecificExercise(id int) *model.RoutineSpecificExercise {
	exercise := model.RoutineSpecificExercise{}
	storage.DB.Find(&exercise, id)

	return &exercise
}

// GetSpecificExercises Get all non deleted routine specific exercises and by using a filter
func GetSpecificExercises(filter model.RoutineSpecificExercise) *[]model.RoutineSpecificExercise {
	exercises := []model.RoutineSpecificExercise{}
	storage.DB.Where(&filter).Find(&exercises)

	for i := 0; i < len(exercises); i++ {
		exercises[i].BaseExercise = *(GetBaseExercise(int(exercises[i].BaseExercisesID)))
	}

	return &exercises
}

// UpdateSpecificExercise Update specific category using id param in URL
func UpdateSpecificExercise(updatedExercise model.RoutineSpecificExercise) error {

	output := storage.DB.Save(&updatedExercise)

	return output.Error
}

// DeleteSpecificExercise Delete routine specific exercise by id, logical delete
func DeleteSpecificExercise(id int) {
	deletedExercise := GetSpecificExercise(id)

	storage.DB.Delete(&deletedExercise)

}
