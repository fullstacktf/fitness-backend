package model

// RoutineSpecificExercise model
type RoutineSpecificExercise struct {
	ID                 uint8           `gorm:"column:id;type:mediumint unsigned;autoIncrement;primaryKey"`
	BaseExercisesID    uint8           `gorm:"column:base_exercises_id;type:mediumint unsigned"`
	AssignedRoutinesID uint8           `gorm:"column:assigned_routines_id;type:mediumint unsigned"`
	Series             uint8           `gorm:"column:series;type:tinyint unsigned"`
	Repetitions        uint8           `gorm:"column:repetitions;type:tinyint unsigned"`
	BaseExercise       BaseExercise    `gorm:"foreignKey:BaseExercisesID;references:ID"`
	AssignedRoutine    AssignedRoutine `gorm:"foreignKey:AssignedRoutinesID;references:ID"`
}

// TableName for RoutineSpecificExercise model
func (r *RoutineSpecificExercise) TableName() string {
	return "routine_specific_exercises"
}

// CreateRoutineSpecificExercise Create an instance of a specific exercise
func (r *RoutineSpecificExercise) CreateRoutineSpecificExercise() string {
	return "CreateRoutineSpecificExercise"
}

// GetRoutineSpecificExercise Get a specific exercise
func (r *RoutineSpecificExercise) GetRoutineSpecificExercise() string {
	return "GetRoutineSpecificExercise"
}

// UpdateRoutineSpecificExercise Update a specific exercise
func (r *RoutineSpecificExercise) UpdateRoutineSpecificExercise() string {
	return "UpdateRoutineSpecificExercise"
}

// DeleteRoutineSpecificExercise Delete a specific exercise
func (r *RoutineSpecificExercise) DeleteRoutineSpecificExercise() string {
	return "DeleteRoutineSpecificExercise"
}
