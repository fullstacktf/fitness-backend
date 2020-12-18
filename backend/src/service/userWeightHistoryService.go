package service

import (
	"time"

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

// PopulateWeightHistory populates weight history element with basic data
func PopulateWeightHistory() {

	historyPoints := make([]model.UserWeightHistory, 3)

	historyPointA := model.UserWeightHistory{
		Weight: 61,
		UserID: 1,
		Date:   time.Now().UTC(),
	}

	historyPointB := model.UserWeightHistory{
		Weight: 63,
		UserID: 1,
		Date:   time.Now().UTC(),
	}

	historyPointC := model.UserWeightHistory{
		Weight: 62,
		UserID: 1,
		Date:   time.Now().UTC(),
	}

	historyPoints[0] = historyPointA
	historyPoints[1] = historyPointB
	historyPoints[2] = historyPointC
}
