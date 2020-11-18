package main

import (
	"log"
	"os"
	"time"

	"github.com/fullstacktf/fitness-backend/config"
	"github.com/fullstacktf/fitness-backend/model"
	"github.com/fullstacktf/fitness-backend/routes"
	s "github.com/fullstacktf/fitness-backend/storage"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var err error

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Silent,
			Colorful:      true,
		},
	)

	s.DB, err = gorm.Open(mysql.Open(config.DbURL(config.BuildDBConfig())), &gorm.Config{Logger: newLogger})

	if err != nil {
		log.Fatal("error al conectar a la base de datos:", err)
	}

	err := s.DB.AutoMigrate(
		&model.Welcome{},
		&model.Permission{},
		&model.Role{},
		&model.User{},
		&model.UserPass{},
		&model.UserStat{},
		&model.UserWeightHistory{},
		&model.RoutineCategory{},
		&model.ExerciseCategory{},
		&model.BaseExercise{},
		&model.Muscle{},
		&model.AssignedRoutine{},
		&model.BaseRoutine{},
		&model.RoutineSpecificExercise{},
		&model.History{},
	)

	if err != nil {
		log.Fatal("error al crear tablas de la base de datos:", err)
	}

	model.PopulatePermissions()
	model.PopulateRoles()

	model.AssociatePermissions()

	routes.SetupRouter().Run()
}
