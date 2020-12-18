package service

import (
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateExerciseCategory creates a exercise category with given data
func CreateExerciseCategory(data model.ExerciseCategory) error {

	output := storage.DB.Create(&data)

	return output.Error
}

// GetExerciseCategory Get category by id
func GetExerciseCategory(id int) *model.ExerciseCategory {
	category := model.ExerciseCategory{}
	storage.DB.Find(&category, id)

	return &category
}

// GetExerciseCategories Get all non deleted categories and by using a filter
func GetExerciseCategories() *[]model.ExerciseCategory {
	categories := []model.ExerciseCategory{}
	storage.DB.Find(&categories)

	return &categories
}

// UpdateExerciseCategory Update specific category using id param in URL
func UpdateExerciseCategory(updatedCategory model.ExerciseCategory) error {

	output := storage.DB.Save(&updatedCategory)

	return output.Error
}

// DeleteExerciseCategory Delete category by id, logical delete
func DeleteExerciseCategory(id int) {
	deletedCategory := GetExerciseCategory(id)

	storage.DB.Delete(&deletedCategory)

}

// PopulateExerciseCategory populates exercise category element with basic data
func PopulateExerciseCategory() {

	fuerza := model.ExerciseCategory{
		Name: "Fuerza",
	}

	rehab := model.ExerciseCategory{
		Name: "Rehabilitación",
	}

	peso := model.ExerciseCategory{
		Name: "Pérdida de peso",
	}

	mantenimiento := model.ExerciseCategory{
		Name: "Mantenimiento",
	}

	categories := []model.ExerciseCategory{
		fuerza,
		rehab,
		peso,
		mantenimiento,
	}

	storage.DB.Create(&categories)

}
