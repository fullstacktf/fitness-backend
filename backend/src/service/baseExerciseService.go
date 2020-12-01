package service

import (
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateBaseExercise Creates a BaseExercise with the given data
func CreateBaseExercise(data model.BaseExercise) error {

	output := storage.DB.Create(&data)

	return output.Error
}

// GetBaseExercise Get BaseExercise by id
func GetBaseExercise(id int) *model.BaseExercise {
	baseEx := model.BaseExercise{}
	storage.DB.Find(&baseEx, id)

	return &baseEx
}

// GetBaseExercises Get all non deleted BaseExercises and by using a filter
func GetBaseExercises(filter model.BaseExercise) *[]model.BaseExercise {
	baseExs := []model.BaseExercise{}
	storage.DB.Where(&filter).Find(&baseExs)

	return &baseExs
}

// UpdateBaseExercise Update specific BaseExercise using id param in URL
func UpdateBaseExercise(updatedBaseExercise model.BaseExercise) error {

	output := storage.DB.Save(&updatedBaseExercise)

	return output.Error
}

// DeleteBaseExercise Delete BaseExercise by id, logical delete
func DeleteBaseExercise(id int) {
	deletedBaseExercise := GetBaseExercise(id)

	storage.DB.Delete(&deletedBaseExercise)

}
