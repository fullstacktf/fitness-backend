package service

import (
	"github.com/fullstacktf/fitness-backend/constants"
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateUser Create a new user with specified data in body
func CreateUser(data model.User) error {

	output := storage.DB.Create(&data)

	userStats := model.UserStat{}
	userStats.UserID = uint64(data.ID)
	userStats.Height = 0
	userStats.LastWeight = 0

	storage.DB.Create(&userStats)

	userPass := model.UserPass{}
	userPass.Password = GenerateRandomPassword(10, constants.PasswordCharset)
	userPass.UserID = uint64(data.ID)

	storage.DB.Create(&userPass)

	SendRegistrationMail(data.Email, userPass.Password)

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
func UpdateUser(updatedUser model.User) error {

	output := storage.DB.Save(&updatedUser)

	return output.Error
}

// DeleteUser Delete user by id, logical delete
func DeleteUser(id int) {
	deletedUser := GetUser(id)

	storage.DB.Delete(&deletedUser)

}
