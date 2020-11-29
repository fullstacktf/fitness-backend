package model

import (
	"time"

	"gorm.io/gorm"
)

// UserWeightHistory model
type UserWeightHistory struct {
	gorm.Model
	Weight float32   `gorm:"column:weight;type:float unsigned"`
	Date   time.Time `gorm:"column:date;type:date"`
	UserID uint64    `gorm:"column:user_id;type:bigint(20) unsigned"`
	User   User      `gorm:"foreignKey:UserID;"`
}

// TableName Function to change the name of a table. In this case of UserWeightHistory model
func (uwh *UserWeightHistory) TableName() string {
	return "user_weight_history"
}

// CreateUserWeightHistory Creates a new weight history data point
func (uwh *UserWeightHistory) CreateUserWeightHistory() string {
	return "CreateUserWeightHistory"
}

// UpdateUserWeightHistory Updates an existing weight history data point
func (uwh *UserWeightHistory) UpdateUserWeightHistory() string {
	return "UpdateUserWeightHistory"
}
