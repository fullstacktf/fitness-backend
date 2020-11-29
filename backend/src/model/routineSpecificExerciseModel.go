package model

import "gorm.io/gorm"

// RoutineSpecificExercise model
type RoutineSpecificExercise struct {
	gorm.Model
	BaseExercisesID    uint64          `gorm:"column:base_exercises_id;type:bigint(20) unsigned"`
	AssignedRoutinesID uint64          `gorm:"column:assigned_routines_id;type:bigint(20) unsigned"`
	Series             uint8           `gorm:"column:series;type:tinyint unsigned"`
	Repetitions        uint8           `gorm:"column:repetitions;type:tinyint unsigned"`
	BaseExercise       BaseExercise    `gorm:"foreignKey:BaseExercisesID;"`
	AssignedRoutine    AssignedRoutine `gorm:"foreignKey:AssignedRoutinesID;"`
}

// TableName for RoutineSpecificExercise model
func (r *RoutineSpecificExercise) TableName() string {
	return "routine_specific_exercises"
}

// CreateRoutineSpecificExercise Create an instance of a routine specific exercise
func (r *RoutineSpecificExercise) CreateRoutineSpecificExercise() string {
	return "CreateRoutineSpecificExercise"
}

// GetRoutineSpecificExercise Get a routine specific exercise
func (r *RoutineSpecificExercise) GetRoutineSpecificExercise() string {
	return "GetRoutineSpecificExercise"
}

// UpdateRoutineSpecificExercise Update a routine specific exercise
func (r *RoutineSpecificExercise) UpdateRoutineSpecificExercise() string {
	return "UpdateRoutineSpecificExercise"
}

// DeleteRoutineSpecificExercise Deletes a routine specific exercise
func (r *RoutineSpecificExercise) DeleteRoutineSpecificExercise() string {
	return "DeleteRoutineSpecificExercise"
}
