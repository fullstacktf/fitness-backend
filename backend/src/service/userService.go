package service

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateUser Create a new user
func CreateUser(data model.User) error {

	fmt.Println(data)

	result := storage.DB.Create(&data)

	return result.Error
}

// GetUser Get user by id
func GetUser(id int) *model.User {
	user := model.User{}
	storage.DB.Find(&user, id)

	return &user
}

// GetUser Get user by DNI
func GetUserByDni(dni string) *model.User {
	user := model.User{}

	storage.DB.Where("dni=?", dni).Find(&user)

	return &user
}

// GetUser Get user by Email
func GetUserByEmail(email string) *model.User {
	user := model.User{}

	storage.DB.Where("email=?", email).Find(&user)

	return &user
}

// GetUsers Get all users
func GetUsers() *[]model.User {
	users := []model.User{}
	storage.DB.Find(&users)

	return &users
}

// UpdateUser Update specific user
func UpdateUser(updatedUser model.User) error {

	return nil
}

// DeleteUser Delete user by id
func DeleteUser() string {
	return "DeleteUser"
}
