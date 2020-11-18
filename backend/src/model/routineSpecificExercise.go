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
