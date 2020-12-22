package service

import (
	"time"

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

	userRole := model.Role{}
	storage.DB.Model(&user).Association("Role").Find(&userRole)

	user.Role = userRole

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

// PopulateUser Function to populate the user model with example data
func PopulateUser() {
	user := model.User{
		DNI:       "01234567A",
		Name:      "usuarioEjemplo",
		Email:     "usuarioEjemplo@youlift.xyz",
		Phone:     "",
		BirthDate: time.Date(1990, 01, 05, 0, 0, 0, 0, time.UTC),
		Address:   "calle ejemplo",
		Biography: "Example biography",
		UserRole:  3,
	}

	admin := model.User{
		DNI:       "01234567B",
		Name:      "adminEjemplo",
		Email:     "adminEjemplo@youlift.xyz",
		Phone:     "",
		BirthDate: time.Date(1990, 01, 05, 0, 0, 0, 0, time.UTC),
		Address:   "calle ejemplo",
		Biography: "Example biography",
		UserRole:  1,
	}

	trainer := model.User{
		DNI:       "01234567C",
		Name:      "entrenadorEjemplo",
		Email:     "entrenadorEjemplo@youlift.xyz",
		Phone:     "",
		BirthDate: time.Date(1990, 01, 05, 0, 0, 0, 0, time.UTC),
		Address:   "calle ejemplo",
		Biography: "Example biography",
		UserRole:  1,
	}

	CreateUser(user)
	CreateUser(admin)
	CreateUser(trainer)
}
