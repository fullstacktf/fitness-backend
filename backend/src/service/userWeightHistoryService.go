package service

import (
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateUserWeightHistoryPoint Creates a new weight history data point
func CreateUserWeightHistoryPoint(data model.UserWeightHistory) error {

	output := storage.DB.Create(&data)

	return output.Error
}

// GetWeightHistoryPoints Gets a paged weight history points
func GetWeightHistoryPoints(data model.Pagedrequest, userID int) *[]model.UserWeightHistory {

	historyPoints := []model.UserWeightHistory{}

	storage.DB.Where("user_id = ?", userID).Order("date desc").Limit(data.Size).Offset(data.Size * data.Offset).Find(&historyPoints)

	return &historyPoints

}
