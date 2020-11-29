package model

import "gorm.io/gorm"

// RoutineCategory model
type RoutineCategory struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(25)"`
}

// TableName Function to change the name of a table. In this case of RoutineCategory model
func (rc *RoutineCategory) TableName() string {
	return "routine_category"
}

// CreateRoutineCategory Creates a routine category
func (rc *RoutineCategory) CreateRoutineCategory() string {
	return "Create a Routinecategory"
}

// GetRoutineCategory Gets a routine category
func (rc *RoutineCategory) GetRoutineCategory() string {
	return "Get a Routinecategory"
}

// GetRoutineCategories Gets all routine categories
func (rc *RoutineCategory) GetRoutineCategories() string {
	return "Get all Routinecategory"
}

// UpdateRoutineCategory Updates a routine category
func (rc *RoutineCategory) UpdateRoutineCategory() string {
	return "Update a Routinecategory"
}

// DeleteRoutineCategory Deletes a routine category
func (rc *RoutineCategory) DeleteRoutineCategory() string {
	return "Delete a Routinecategory"
}
