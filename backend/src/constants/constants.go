package constants

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DatabaseHost Database host name
	DatabaseHost string
	// DatabaseUser Database user name
	DatabaseUser string
	// DatabaseName Database name
	DatabaseName string
	// DatabasePassword Database password
	DatabasePassword string
	// DatabasePort Database port number
	DatabasePort int

	// RegistrationMessage Registration/Welcome message with parameters
	RegistrationMessage string
	// ResetPasswordMessage Reset password message with parameters
	ResetPasswordMessage string
	// PasswordCharset Charset for the password generation
	PasswordCharset string

	// EmailSender Email account for automated emails
	EmailSender string
	// EmailPassword Automated emails account's password
	EmailPassword string
)

func init() {
	godotenv.Load()
	godotenv.Load("../.env")
	DatabaseHost = os.Getenv("DATABASE_HOST")
	DatabaseUser = os.Getenv("DATABASE_USER")
	DatabaseName = os.Getenv("DATABASE_NAME")
	DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	DatabasePort, _ = strconv.Atoi(os.Getenv("DATABASE_PORT"))

	RegistrationMessage = "Bienvenido a YouLift! %s,\n\nEstás a nada de disfrutar de una vida más activa y saludable!\n\nAccede a nuestra página https://www.youlift.xyz/ e inicia sesión con la contraseña indicada a continuación.\n\nContraseña: %s"
	ResetPasswordMessage = "Tu nueva contraseña en youlift.xyz\n\n%s"
	PasswordCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	EmailSender = os.Getenv("EMAIL_SENDER")
	EmailPassword = os.Getenv("EMAIL_PASSWORD")
}
