package service

import (
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/storage"
)

// CreateHistory Create a new history point
func CreateHistory(data model.History) error {
	output := storage.DB.Create(&data)

	return output.Error
}

// GetHistoryPoints Gets a paginated list of history point of the specified user
func GetHistoryPoints(data model.Pagedrequest, specExerciseID int) *[]model.History {

	historyPoints := []model.History{}

	storage.DB.Where("routine_specific_exercises_id = ?", specExerciseID).Order("date desc").Limit(data.Size).Offset(data.Size * data.Offset).Find(&historyPoints)

	return &historyPoints

}
