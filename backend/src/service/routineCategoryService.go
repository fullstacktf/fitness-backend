package service

import (
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateRoutineCategory creates a category with the given data
func CreateRoutineCategory(data model.RoutineCategory) error {

	output := storage.DB.Create(&data)

	return output.Error
}

// GetRoutineCategory Get category by id
func GetRoutineCategory(id int) *model.RoutineCategory {
	category := model.RoutineCategory{}
	storage.DB.Find(&category, id)

	return &category
}

// GetRoutineCategories Get all non deleted categories and by using a filter
func GetRoutineCategories() *[]model.RoutineCategory {
	categories := []model.RoutineCategory{}
	storage.DB.Find(&categories)

	return &categories
}

// UpdateRoutineCategory Update specific category using id param in URL
func UpdateRoutineCategory(updatedCategory model.RoutineCategory) error {

	output := storage.DB.Save(&updatedCategory)

	return output.Error
}

// DeleteRoutineCategory Delete category by id, logical delete
func DeleteRoutineCategory(id int) {
	deletedCategory := GetRoutineCategory(id)

	storage.DB.Delete(&deletedCategory)

}
