package model

import (
	"time"

	"gorm.io/gorm"
)

// History model
type History struct {
	gorm.Model
	Measure                    float32                 `gorm:"column:measure;type:float unsigned"`
	MeasureUnit                string                  `gorm:"column:measure_unit;type:enum('s', 'kg', 'ms')"`
	MeasureType                string                  `gorm:"column:measure_type;type:enum('weight', 'exerciseMetric')"`
	Date                       time.Time               `gorm:"column:date;type:date"`
	RoutineSpecificExercisesID uint64                  `gorm:"column:routine_specific_exercises_id;type:bigint(20) unsigned"`
	RoutineSpecificExercise    RoutineSpecificExercise `gorm:"foreignKey:RoutineSpecificExercisesID;"`
}

// TableName Function to change the name of a table. In this case of History model
func (h *History) TableName() string {
	return "history"
}
