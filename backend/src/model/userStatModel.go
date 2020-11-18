package model

import "time"

// UserStat model
type UserStat struct {
	UserID       uint8     `gorm:"column:user_id;type:mediumint unsigned"`
	LastWeight   float32   `gorm:"column:last_weight;type:float unsigned"`
	Height       float32   `gorm:"column:height;type:float unsigned"`
	InjuredSince time.Time `gorm:"column:injured_since;type:date"`
	InjuredUntil time.Time `gorm:"column:injured_until;type:date"`
	User         User      `gorm:"foreignKey:UserID;references:ID"`
}

// TableName for User stats model
func (u *UserStat) TableName() string {
	return "user_stats"
}

// UpdateUserStat Updates the user stats
func (u *UserStat) UpdateUserStat() string {
	return "UpdateUserStat"
}
