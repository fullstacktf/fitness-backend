package service

import (
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CheckLogin Checks if the given email has the given password and returns it's user ID
func CheckLogin(username string, password string) (bool, uint) {

	user := model.User{}
	userPass := model.UserPass{}

	storage.DB.Where("email = ? ", username).Find(&user)

	if user.ID == 0 {
		return false, 0
	}

	storage.DB.Where("user_id = ? ", user.ID).Find(&userPass)

	if password == userPass.Password {
		return true, user.ID
	}

	return false, 0

}

// CheckPermission Checks if a user has the specified permission
func CheckPermission(permission string, userID uint) bool {
	user := model.User{}
	role := model.Role{}
	permissions := []model.Permission{}

	storage.DB.Find(&user, userID)
	storage.DB.Model(&user).Association("Role").Find(&role)
	storage.DB.Model(&role).Association("Permissions").Find(&permissions)

	for i := 0; i < len(permissions); i++ {
		if permission == permissions[i].Description {
			return true
		}
	}

	return false

}

// CheckRole Checks if a user has the specified role
func CheckRole(role string, userID uint) bool {
	user := model.User{}
	userRole := model.Role{}

	storage.DB.Find(&user, userID)
	storage.DB.Model(&user).Association("Role").Find(&userRole)

	if userRole.Description == role {
		return true

	}

	return false

}
