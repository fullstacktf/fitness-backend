package service

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	"github.com/fullstacktf/fitness-backend/constants"
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

// GenerateRandomPassword Generates a random password with the indicated charset and length
func GenerateRandomPassword(length int, charset string) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// SendMail Sends an email
func SendMail(to []string, content string) {

	// Sender data.
	from := "supp.youlift.xyz@gmail.com"
	password := "YbrxEqh0Uhee9pQL27sN"

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(content)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		return
	}
}

// SendRegistrationMail Sends a registration/welcome email with the auto-generated password
func SendRegistrationMail(to string, pass string) {
	dest := make([]string, 1)
	dest[0] = to
	message := fmt.Sprintf(constants.RegistrationMessage, to, pass)

	SendMail(dest, message)
}

// SendResetPasswordMail Sends an email with a new auto-generated password
func SendResetPasswordMail(to string, pass string) {
	dest := make([]string, 1)
	dest[0] = to
	message := fmt.Sprintf(constants.ResetPasswordMessage, pass)

	SendMail(dest, message)
}
