package config

import (
	"fmt"

	"github.com/fullstacktf/fitness-backend/constants"
)

// DBConfig Database configuration struct
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

// BuildDBConfig Function for init the DBConfig struct
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     constants.DatabaseHost,
		Port:     constants.DatabasePort,
		User:     constants.DatabaseUser,
		DBName:   constants.DatabaseName,
		Password: constants.DatabasePassword,
	}
	fmt.Println(dbConfig)
	return &dbConfig
}

// DbURL function
func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
