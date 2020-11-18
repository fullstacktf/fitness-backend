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
)

func init() {
	godotenv.Load()
	godotenv.Load("../.env")
	DatabaseHost = os.Getenv("DATABASE_HOST")
	DatabaseUser = os.Getenv("DATABASE_USER")
	DatabaseName = os.Getenv("DATABASE_NAME")
	DatabasePassword = os.Getenv("DATABASE_PASSWORD")
	DatabasePort, _ = strconv.Atoi(os.Getenv("DATABASE_PORT"))
}
