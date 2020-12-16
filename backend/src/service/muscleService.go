package service

import (
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateMuscle Create a new muscle with specified data in body
func CreateMuscle(data model.Muscle) error {

	output := storage.DB.Create(&data)

	return output.Error
}

// GetMuscle Get muscle by id
func GetMuscle(id int) *model.Muscle {
	muscle := model.Muscle{}
	storage.DB.Find(&muscle, id)

	return &muscle
}

// GetMuscles Get all non deleted muscles and by using a filter
func GetMuscles() *[]model.Muscle {
	muscles := []model.Muscle{}
	storage.DB.Find(&muscles)

	return &muscles
}

// UpdateMuscle Update specific muscle using id param in URL
func UpdateMuscle(updatedMuscle model.Muscle) error {

	output := storage.DB.Save(&updatedMuscle)

	return output.Error
}

// DeleteMuscle Delete muscle by id, logical delete
func DeleteMuscle(id int) {
	deletedMuscle := GetMuscle(id)

	storage.DB.Delete(&deletedMuscle)

}
