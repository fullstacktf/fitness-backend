package service

import (
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateBaseRoutine Creates a base routine with the given data
func CreateBaseRoutine(data model.BaseRoutine) error {

	output := storage.DB.Create(&data)

	return output.Error
}

// GetBaseRoutine Get base routine by id
func GetBaseRoutine(id int) *model.BaseRoutine {
	baseRoutine := model.BaseRoutine{}
	storage.DB.Find(&baseRoutine, id)

	baseRoutine.GetBaseRoutineAssociations()

	return &baseRoutine
}

// GetBaseRoutines Get all non deleted base routines and by using a filter
func GetBaseRoutines(filter model.BaseRoutine) *[]model.BaseRoutine {
	baseRoutines := []model.BaseRoutine{}
	storage.DB.Where(&filter).Find(&baseRoutines)

	for i := 0; i < len(baseRoutines); i++ {
		baseRoutines[i].GetBaseRoutineAssociations()
	}

	return &baseRoutines
}

// UpdateBaseRoutine Update specific base routine using id param in URL
func UpdateBaseRoutine(updatedBaseRoutine model.BaseRoutine) error {

	storage.DB.Model(&updatedBaseRoutine).Association("BaseExercises").Replace(updatedBaseRoutine.BaseExercises)

	output := storage.DB.Save(&updatedBaseRoutine)

	return output.Error
}

// DeleteBaseRoutine Delete BaseRoutine by id, logical delete
func DeleteBaseRoutine(id int) {
	deletedBaseRoutine := GetBaseRoutine(id)

	storage.DB.Delete(&deletedBaseRoutine)

}
