package model

import (
	"time"

	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	DNI       string    `gorm:"column:dni;type:char(9)"`
	Name      string    `gorm:"column:name;type:varchar(50)"`
	Surname   string    `gorm:"column:surname;type:varchar(100)"`
	Email     string    `gorm:"column:email;type:varchar(50);unique;"`
	Phone     string    `gorm:"column:phone;type:varchar(20)"`
	BirthDate time.Time `gorm:"column:birth_date;type:date"`
	Address   string    `gorm:"column:address;type:varchar(100)"`
	UserRole  uint64    `gorm:"column:user_role;type:bigint(20) unsigned;"`
	Role      Role      `gorm:"foreignKey:UserRole;"`
}

// TableName Function to change the name of a table.
func (u *User) TableName() string {
	return "users"
}
