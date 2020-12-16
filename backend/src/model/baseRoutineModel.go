package model

import (
	"github.com/fullstacktf/fitness-backend/storage"
	"gorm.io/gorm"
)

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

func (br *BaseRoutine) GetBaseRoutineAssociations() {
	baseExercises := []*BaseExercise{}
	storage.DB.Model(&br).Association("BaseExercises").Find(&baseExercises)
	br.BaseExercises = baseExercises

	routineCategory := RoutineCategory{}
	storage.DB.Model(&br).Association("RoutineCategory").Find(&routineCategory)
	br.RoutineCategory = routineCategory
}
