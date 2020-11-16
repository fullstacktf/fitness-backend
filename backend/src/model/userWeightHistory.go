package model

import "time"

// UserWeightHistory model
type UserWeightHistory struct {
	ID     uint8     `gorm:"column:id;type:mediumint unsigned;autoIncrement;primaryKey"`
	Weight float32   `gorm:"column:weight;type:float unsigned"`
	Date   time.Time `gorm:"column:date;type:date"`
	UserID uint8     `gorm:"column:user_id;type:mediumint unsigned"`
	User   User      `gorm:"foreignKey:UserID;references:ID"`
}

// TableName for UserPass model
func (uwh *UserWeightHistory) TableName() string {
	return "user_weight_history"
}
