package model

import "gorm.io/gorm"

// ExerciseCategory model
type ExerciseCategory struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(25)"`
}

// TableName Function to change the name of a table. In this case of ExerciseCategory model
func (ec *ExerciseCategory) TableName() string {
	return "exercise_category"
}
