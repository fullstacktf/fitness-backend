package model

import (
	"github.com/fullstacktf/fitness-backend/storage"
	"gorm.io/gorm"
)

// BaseExercise model
type BaseExercise struct {
	gorm.Model
	Name               string           `gorm:"column:name;type:varchar(100)"`
	Description        string           `gorm:"column:description;type:varchar(600)"`
	VideoURL           string           `gorm:"column:video_url;type:varchar(100)"`
	CategoryID         uint64           `gorm:"column:category_id;type:bigint(20) unsigned"`
	DefaultSeries      uint8            `gorm:"column:default_series;type:tinyint unsigned"`
	DefaultRepetitions uint8            `gorm:"column:default_repetitions;type:tinyint unsigned"`
	ExerciseCategory   ExerciseCategory `gorm:"foreignKey:CategoryID;"`
	Muscles            []*Muscle        `gorm:"many2many:exercises_muscles;"`
}

// TableName Function to change the name of a table.
func (be *BaseExercise) TableName() string {
	return "base_exercises"
}

// GetBaseExerciseAssociations Adds all the object's associations
func (be *BaseExercise) GetBaseExerciseAssociations() {
	category := ExerciseCategory{}
	storage.DB.Model(&be).Association("ExerciseCategory").Find(&category)
	be.ExerciseCategory = category

	muscles := []*Muscle{}
	storage.DB.Model(&be).Association("Muscles").Find(&muscles)
	be.Muscles = muscles
}
