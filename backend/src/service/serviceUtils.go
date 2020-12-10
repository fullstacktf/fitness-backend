package service

import (
	"strconv"
	"strings"

	"github.com/fullstacktf/fitness-backend/model"
)

// GetGormErrorCode Gets an error code from mysql and format as JSON
func GetGormErrorCode(errorStr string) model.ErrorVO {
	gormError := model.ErrorVO{}

	errArray := strings.Split(errorStr, ":")

	var message = ""
	for i := 1; i < len(errArray); i++ {
		message += errArray[i]
	}

	errorStr = errArray[0]
	errorStr = strings.Split(errorStr, " ")[1]

	var code, _ = strconv.Atoi(errorStr)

	gormError.Code = code
	gormError.Message = message

	return gormError

}
