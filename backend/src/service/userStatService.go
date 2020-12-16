package service

import (
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// GetUserStat Gets the stats of the specified user
func GetUserStat(userID int) *model.UserStat {
	userStat := model.UserStat{}
	storage.DB.Where("user_id = ?", userID).Find(&userStat)

	return &userStat
}

// UpdateUserStats Updates stats of the specified user
func UpdateUserStats(updatedUserStat model.UserStat) error {

	output := storage.DB.Where("user_id = ? ", updatedUserStat.UserID).Save(&updatedUserStat)

	return output.Error
}
