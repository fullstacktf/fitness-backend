package service

import (
	"github.com/fullstacktf/fitness-backend/constants"
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// ResetUserPass Generates a new password for the specified user and sends an email
func ResetUserPass(userID int) {

	user := model.User{}

	storage.DB.Find(&user, userID)

	updatedPassword := model.UserPass{}

	storage.DB.Where("user_id = ?", userID).Find(&updatedPassword)

	updatedPassword.Password = GenerateRandomPassword(10, constants.PasswordCharset)

	storage.DB.Where("user_id = ?", userID).Save(updatedPassword)

	SendResetPasswordMail(user.Email, updatedPassword.Password)

}
