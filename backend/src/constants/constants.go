package constants

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// DatabaseHost host name
	DatabaseHost string
	// DatabaseUser user name
	DatabaseUser string
	// DatabaseName database name
	DatabaseName string
	// DatabasePassword database password
	DatabasePassword string
	// DatabasePort port number
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

	fmt.Println("VARIABLE:", DatabaseHost)
}
