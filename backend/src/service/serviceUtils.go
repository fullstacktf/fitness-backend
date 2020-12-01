package service

import (
	"strconv"
	"strings"

	"github.com/fullstacktf/fitness-backend/model"
)

func GetGormErrorCode(errorStr string) model.ErrorVO {
	gormError := model.ErrorVO{}

	errorStr = strings.Split(errorStr, ":")[0]
	errorStr = strings.Split(errorStr, " ")[1]

	var code, _ = strconv.Atoi(errorStr)
	var message = ""

	switch code {
	case 1062:
		message = "Duplicate entry"
	default:
		message = "Ssytem error"
	}

	gormError.Code = code
	gormError.Message = message

	return gormError

}
