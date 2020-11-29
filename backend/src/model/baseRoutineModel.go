package model

import "gorm.io/gorm"

// BaseRoutine model
type BaseRoutine struct {
	gorm.Model
	CategoryID      uint64          `gorm:"column:category_id;type:bigint(20) unsigned"`
	Name            string          `gorm:"column:name;type:varchar(15)"`
	Description     string          `gorm:"column:description;type:varchar(30)"`
	RoutineCategory RoutineCategory `gorm:"foreignKey:CategoryID;"`
	BaseExercises   []*BaseExercise `gorm:"many2many:base_routines_base_exercises"`
}

// TableName Function to change the name of a table.
func (br *BaseRoutine) TableName() string {
	return "base_routines"
}

// CreateBaseRoutine Creates a base routine
func (br *BaseRoutine) CreateBaseRoutine() string {
	return "Create a Base Routine"
}

// GetBaseRoutine Gets a base routine
func (br *BaseRoutine) GetBaseRoutine() string {
	return "Get a BaseRoutine"
}

// GetBaseRoutines Gets all base routine
func (br *BaseRoutine) GetBaseRoutines() string {
	return "Get all BaseExercises"
}

// UpdateBaseRoutine Updates a base routine
func (br *BaseRoutine) UpdateBaseRoutine() string {
	return "Update a Base Routine"
}

// DeleteBaseRoutine Deletes a base routine
func (br *BaseRoutine) DeleteBaseRoutine() string {
	return "Delete a Base Routine"
}
