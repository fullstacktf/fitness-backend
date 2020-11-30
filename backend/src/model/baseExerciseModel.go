package model

import "gorm.io/gorm"

// BaseExercise model
type BaseExercise struct {
	gorm.Model
	Name             string           `gorm:"column:name;type:varchar(25)"`
	Description      string           `gorm:"column:description;type:varchar(60)"`
	VideoURL         string           `gorm:"column:video_url;type:varchar(50)"`
	CategoryID       uint64           `gorm:"column:category_id;type:bigint(20) unsigned"`
	Series           uint             `gorm:"column:series;type:tinyint unsigned"`
	Repetitions      uint             `gorm:"column:repetitions;type:tinyint unsigned"`
	ExerciseCategory ExerciseCategory `gorm:"foreignKey:CategoryID;"`
	Muscles          []*Muscle        `gorm:"many2many:exercises_muscles;"`
}

// TableName Function to change the name of a table.
func (be *BaseExercise) TableName() string {
	return "base_exercises"
}
