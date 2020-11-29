package model

import "time"

// UserStat model
type UserStat struct {
	UserID       uint64    `gorm:"column:user_id;type:bigint(20) unsigned"`
	LastWeight   float32   `gorm:"column:last_weight;type:float unsigned"`
	Height       float32   `gorm:"column:height;type:float unsigned"`
	InjuredSince time.Time `gorm:"column:injured_since;type:date"`
	InjuredUntil time.Time `gorm:"column:injured_until;type:date"`
	User         User      `gorm:"foreignKey:UserID;"`
}

// TableName Function to change the name of a table.
func (u *UserStat) TableName() string {
	return "user_stats"
}

// UpdateUserStat Updates the user stats
func (u *UserStat) UpdateUserStat() string {
	return "UpdateUserStat"
}
