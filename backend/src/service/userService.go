package service

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateUser Create a new user with specified data in body
func CreateUser(data model.User) error {

	fmt.Println(data)

	output := storage.DB.Create(&data)

	return output.Error
}

// GetUser Get user by id
func GetUser(id int) *model.User {
	user := model.User{}
	storage.DB.Find(&user, id)

	return &user
}

// GetUsers Get all non deleted users and by using a filter
func GetUsers(filter model.User) *[]model.User {
	users := []model.User{}
	storage.DB.Where(&filter).Find(&users)

	return &users
}

// UpdateUser Update specific user using id param in URL
func UpdateUser(updatedUser model.User, id int) error {
	updatedUser.ID = uint(id)

	output := storage.DB.Save(&updatedUser)

	return output.Error
}

// DeleteUser Delete user by id, logical delete
func DeleteUser(id int) {
	deletedUser := GetUser(id)

	storage.DB.Delete(&deletedUser)

}
